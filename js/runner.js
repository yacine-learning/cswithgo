// Page side of a running Go program: one worker per run, so Stop (or a time limit) simply ends it.
let current = null;
// opts: onOutput(text, fd), onExit(code, ms, message), timeout (ms, optional)
export function run(module, { onOutput, onExit, timeout }) {
  stop();
  const w = new Worker(new URL('./run-worker.js', import.meta.url)), t0 = performance.now();
  let timer = null;
  const end = (code, msg) => { if (current !== w) return; current = null; clearTimeout(timer); w.terminate(); onExit(code, Math.round(performance.now() - t0), msg); };
  w.onmessage = (e) => { const d = e.data; if (d.t === 'out') onOutput(d.text, d.fd); else if (d.t === 'exit') end(d.code); else if (d.t === 'crash') end(2, d.msg); };
  if (timeout) timer = setTimeout(() => end(-1, `stopped after ${timeout / 1000} s`), timeout);
  current = w; w.postMessage({ module });
}
export function stop() { if (!current) return false; current.terminate(); current = null; return true; }
export const running = () => !!current;
