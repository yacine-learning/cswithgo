// Runs one compiled Go program (js/wasm). What it prints is sent to the page as it happens.
// There is no file system and no network: os.Stdout and os.Stderr are the only outputs.
const dec = new TextDecoder();
const out = (fd, buf) => postMessage({ t: 'out', fd, text: dec.decode(buf) });
const nosys = (cb) => { const e = new Error('not available in the browser'); e.code = 'ENOSYS'; cb(e); };
globalThis.fs = {
  constants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1, O_DIRECTORY: -1 },
  writeSync(fd, buf) { out(fd, buf); return buf.length; },
  write(fd, buf, off, len, pos, cb) {
    if (fd !== 1 && fd !== 2) return nosys(cb);
    out(fd, buf.subarray(off, off + len)); cb(null, len);
  },
  read(fd, buf, off, len, pos, cb) { if (fd === 0) return cb(null, 0); nosys(cb); },   // stdin is empty
  fstat(fd, cb) { if (fd <= 2) return cb(null, { dev: 0, ino: 0, mode: 0o020666, nlink: 1, uid: 0, gid: 0, rdev: 0, size: 0, blksize: 4096, blocks: 0, atimeMs: 0, mtimeMs: 0, ctimeMs: 0, isDirectory: () => false }); nosys(cb); },
};
for (const n of ['open', 'close', 'stat', 'lstat', 'readdir', 'mkdir', 'unlink', 'rmdir', 'rename', 'ftruncate', 'truncate', 'fsync', 'chmod', 'fchmod', 'chown', 'fchown', 'lchown', 'utimes', 'readlink', 'link', 'symlink'])
  globalThis.fs[n] = (...a) => nosys(a[a.length - 1]);
globalThis.process = { getuid: () => -1, getgid: () => -1, geteuid: () => -1, getegid: () => -1, getgroups: () => [], pid: 1, ppid: 0, umask: () => 0o22, cwd: () => '/', chdir: () => {} };
importScripts('wasm_exec.js');

onmessage = (e) => {
  const go = new Go(); go.argv = ['main']; go.env = {};
  go.exit = (code) => postMessage({ t: 'exit', code });
  WebAssembly.instantiate(e.data.module, go.importObject).then((inst) => go.run(inst))
    .catch((err) => postMessage({ t: 'crash', msg: String(err && err.message || err) }));
};
