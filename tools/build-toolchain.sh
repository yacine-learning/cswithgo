#!/usr/bin/env bash
# Builds everything the browser needs to compile and link Go programs:
#   toolchain/compile.wasm.gz, toolchain/link.wasm.gz  the Go compiler and linker, as WebAssembly
#   toolchain/gofmt.wasm.gz                            go/format, for the editor's Format command
#   toolchain/std.pack.gz                              the standard library packages, compiled for js/wasm
#   js/wasm_exec.js                                    Go's JavaScript support file
# Needs the Go version below installed. Run from the repo root: bash tools/build-toolchain.sh
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT=$PWD
GOVER=${GOVER:-go1.24.7}
if [ -d "/usr/local/$GOVER" ]; then export PATH="/usr/local/$GOVER/bin:$PATH"; fi
[ "$(go env GOVERSION)" = "$GOVER" ] || { echo "need $GOVER, found $(go env GOVERSION)"; exit 1; }
GOROOT=$(go env GOROOT)
# packages a program may import (their dependencies are added automatically)
PKGS="bufio bytes cmp container/heap container/list container/ring context crypto/sha256 encoding/binary encoding/hex encoding/json errors fmt hash/crc32 io iter log maps math math/bits math/rand os runtime slices sort strconv strings sync sync/atomic testing time unicode unicode/utf8 unsafe"
TMP=$(mktemp -d); trap 'rm -rf "$TMP"' EXIT
OUT=$ROOT/toolchain; mkdir -p "$OUT"

echo "== compiler and linker"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o "$TMP/compile.wasm" cmd/compile
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o "$TMP/link.wasm" cmd/link
gzip -9 -n -c "$TMP/compile.wasm" > "$OUT/compile.wasm.gz"
gzip -9 -n -c "$TMP/link.wasm" > "$OUT/link.wasm.gz"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o "$TMP/gofmt.wasm" tools/gofmt/main.go
gzip -9 -n -c "$TMP/gofmt.wasm" > "$OUT/gofmt.wasm.gz"

echo "== standard library"
export GOOS=js GOARCH=wasm GOCACHE="$TMP/cache" GOFLAGS=-trimpath
go list -export -deps -f '{{if .Export}}{{.ImportPath}} {{.Export}}{{end}}' $PKGS > "$TMP/list"
python3 - "$TMP" "$OUT" <<'PY'
import gzip, json, struct, sys
tmp, out = sys.argv[1:]
index, blob = {}, bytearray()
for line in open(tmp + '/list'):
    name, path = line.split()
    data = open(path, 'rb').read()
    index[name] = [len(blob), len(data)]
    blob += data
head = json.dumps(index).encode()
with gzip.GzipFile(out + '/std.pack.gz', 'wb', 9, mtime=0) as f:
    f.write(struct.pack('<I', len(head)) + head + blob)
print(len(index), 'packages,', len(blob) // 1024, 'KiB')
PY
cp "$GOROOT/lib/wasm/wasm_exec.js" "$ROOT/js/wasm_exec.js"
echo "$(go env GOVERSION)-$(cat "$OUT"/*.gz | sha256sum | cut -c1-12)" > "$OUT/VERSION"
ls -la "$OUT"
