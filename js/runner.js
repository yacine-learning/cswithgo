// Page side of a running Go program: one worker per run, so Stop simply ends it.
let current = null;
export function run(module, { onOutput, onExit }) {
  stop();
  const w = new Worker(new URL('./run-worker.js', import.meta.url)), t0 = performance.now();
  const end = (code, msg) => { if (current !== w) return; current = null; w.terminate(); onExit(code, Math.round(performance.now() - t0), msg); };
  w.onmessage = (e) => { const d = e.data; if (d.t === 'out') onOutput(d.text, d.fd); else if (d.t === 'exit') end(d.code); else if (d.t === 'crash') end(2, d.msg); };
  current = w; w.postMessage({ module });
}
export function stop() { if (!current) return false; current.terminate(); current = null; return true; }
export const running = () => !!current;
