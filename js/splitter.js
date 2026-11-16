// Drag handles that resize the panes: the course next to the code (a vertical bar), and the code
// above the output (a horizontal bar). Sizes are kept in localStorage; a double-click resets one.
const KEY = 'cswithgo:layout';
const load = () => { try { return JSON.parse(localStorage.getItem(KEY)) || {}; } catch { return {}; } };
const store = (k, v) => { try { const s = load(); if (v == null) delete s[k]; else s[k] = v; localStorage.setItem(KEY, JSON.stringify(s)); } catch {} };
const clamp = (x, lo, hi) => Math.max(lo, Math.min(hi, x));

// drag(handle, onMove(event), onEnd()): pointer capture keeps the drag going over the editor
function drag(handle, onMove, onEnd) {
  handle.addEventListener('pointerdown', (e) => {
    if (e.button !== 0) return;
    e.preventDefault(); handle.setPointerCapture(e.pointerId);
    document.body.classList.add('resizing', handle.classList.contains('split-v') ? 'resizing-v' : 'resizing-h');
    const move = (ev) => onMove(ev);
    const up = () => {
      handle.removeEventListener('pointermove', move); handle.removeEventListener('pointerup', up); handle.removeEventListener('pointercancel', up);
      document.body.classList.remove('resizing', 'resizing-v', 'resizing-h'); onEnd();
    };
    handle.addEventListener('pointermove', move); handle.addEventListener('pointerup', up); handle.addEventListener('pointercancel', up);
  });
}

// the course column (index.html only): width of .lesson, between 280 px and the window minus 360 px
export function splitColumns(app, lesson, bench) {
  const bar = document.createElement('div');
  bar.className = 'split-v'; bar.title = 'Drag to resize, double-click to reset';
  bar.setAttribute('role', 'separator'); bar.setAttribute('aria-orientation', 'vertical');
  app.insertBefore(bar, bench);
  const set = (w) => app.style.setProperty('--course-w', w == null ? '' : w + 'px');
  const fit = (w) => clamp(w, 280, Math.max(280, app.clientWidth - 360));
  const saved = load().course; if (saved) set(fit(saved));
  let w = saved;
  drag(bar, (e) => { w = fit(e.clientX - app.getBoundingClientRect().left); set(w); }, () => store('course', w));
  bar.addEventListener('dblclick', () => { w = null; set(null); store('course', null); });
  addEventListener('resize', () => { if (w) set(fit(w)); });
}

// the output panel (both pages): its height, between 60 px and what leaves the code a few lines
export function splitOutput(frame, out) {
  const bar = document.createElement('div');
  bar.className = 'split-h'; bar.title = 'Drag to resize, double-click to reset';
  bar.setAttribute('role', 'separator'); bar.setAttribute('aria-orientation', 'horizontal');
  out.parentNode.insertBefore(bar, out.previousElementSibling?.classList.contains('problems') ? out.previousElementSibling : out);
  const set = (h) => { out.style.height = h == null ? '' : h + 'px'; };
  const fit = (h) => clamp(h, 60, Math.max(60, frame.clientHeight - 220));
  const saved = load().output; if (saved) set(fit(saved));
  let h = saved;
  drag(bar, (e) => { h = fit(out.getBoundingClientRect().bottom - e.clientY); set(h); }, () => store('output', h));
  bar.addEventListener('dblclick', () => { h = null; set(null); store('output', null); });
  addEventListener('resize', () => { if (h) set(fit(h)); });
}
