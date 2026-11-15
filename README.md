# cswithgo

Computer science with Go, in the browser: lessons, exercises, and the real Go compiler running in the page.

Right now it holds the Go playground: write a program, run it, see its output. The lessons come next, starting with data structures and algorithms.

Everything runs in the page. There is no backend and no build step.

- **The Go compiler and linker** are the official ones (Go 1.24.7), compiled to WebAssembly.
- **The editor** is CodeMirror 6 made to look and work like VS Code: completion and hover documentation from the real standard library, live compiler errors, gofmt.
- **Programs** run in a worker, one per run, so Stop ends them. They can print to stdout and stderr; there is no file system, network or stdin.

## Run it

Any static file server works:

```
python3 -m http.server 8790
```

then open http://localhost:8790/. For GitHub Pages: Settings → Pages → deploy from the `main` branch, folder `/`.

The first build downloads the toolchain (18 MB) and keeps it in the browser's cache storage.

## Packages a program can import

bufio, bytes, cmp, container/heap, container/list, container/ring, context, crypto/sha256, encoding/binary, encoding/hex, encoding/json, errors, fmt, hash/crc32, io, iter, log, maps, math, math/bits, math/rand, os, runtime, slices, sort, strconv, strings, sync, sync/atomic, testing (for `testing.Benchmark`), time, unicode, unicode/utf8, unsafe.

To add one: put it in `PKGS` in `tools/build-toolchain.sh` and in `pkgs` in `tools/genapi/main.go`, then rebuild both (below).

## Layout

| Path | What |
|---|---|
| `index.html`, `js/playground.js`, `css/editor.css` | the playground page |
| `js/toolchain.js`, `js/toolchain-worker.js` | the Go compiler and linker, in a worker, with an in-memory file system |
| `js/runner.js`, `js/run-worker.js` | runs a compiled program in its own worker and streams its output |
| `js/editor.js`, `js/vendor/`, `js/goapi.json`, `js/godoc/` | the editor; `goapi.json` and `godoc/` are the standard library's API and documentation |
| `toolchain/` | compiler, linker, gofmt and standard library, built by `tools/build-toolchain.sh` |
| `tools/` | build scripts: toolchain, API index (`genapi`), vendored editor |
| `test/e2e.mjs` | browser tests (Chromium through Playwright) |

## Rebuilding

```
bash tools/build-toolchain.sh        # compiler, linker, gofmt, standard library (needs Go 1.24.7)
go run tools/genapi/main.go          # the editor's API index and documentation
bash tools/vendor/build-codemirror.sh  # the editor library (needs npm)
```

## Tests

```
python3 -m http.server 8790 &
node test/e2e.mjs
```

Go is distributed under a BSD-style licence (see https://go.dev/LICENSE); `js/wasm_exec.js` and the files in `toolchain/` are built from it. CodeMirror is MIT-licensed (`js/vendor/CODEMIRROR-LICENSE`).
