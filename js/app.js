// The course page: a lesson on the left; on the right the editor, Run, Check and the output.
import { Editor } from './editor.js';
import { toolchain } from './toolchain.js';
import { run, stop } from './runner.js';

const $ = (s) => document.querySelector(s);
const esc = (s) => String(s).replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]);
const KEY = 'cswithgo:v1';
const S = (() => { try { return JSON.parse(localStorage.getItem(KEY)) || {}; } catch { return {}; } })();
for (const k of ['code', 'checked', 'passed', 'hints']) S[k] ||= {};
const save = () => { try { localStorage.setItem(KEY, JSON.stringify(S)); } catch {} };
const toast = (msg) => { const t = document.createElement('div'); t.className = 'toast'; t.textContent = msg; $('#toasts').append(t); setTimeout(() => t.remove(), 2600); };
const status = (msg, cls = '') => { const el = $('#buildStatus'); el.textContent = msg; el.className = 'build-status ' + cls; };
const files = new Map();
const text = (path) => files.get(path) ?? fetch('course/' + path).then((r) => (r.ok ? r.text() : null)).then((t) => (files.set(path, t), t));

let index, lesson, ex = null;          // ex: the exercise in the editor
const exId = (e) => `${lesson.n}-${e.id}`;

// ── navigation ──
function lessonsOf(idx) { const out = []; for (const m of idx.modules) for (const c of m.chapters) for (const l of c.lessons) out.push({ ...l, chapter: c.title, module: m.title }); return out; }
async function openLesson(n) {
  const list = lessonsOf(index), meta = list.find((l) => l.n === n) || list[0];
  lesson = await (await fetch(`course/${meta.n}/lesson.json`)).json();
  const html = await text(`${meta.n}/lesson.html`);
  $('#nav').value = meta.n; S.lesson = meta.n; save();
  $('#lesson').innerHTML = `<h1><span class="no">${esc(meta.n)}</span>${esc(lesson.title)}</h1>` + html;
  for (const e of lesson.exercises) { const slot = $(`[data-ex="${e.id}"]`); if (slot) { slot.className = 'ex'; slot.dataset.id = e.id; drawExercise(e); } }
  $('#lesson').scrollTop = 0;
  const last = lesson.exercises.find((e) => exId(e) === S.last) || lesson.exercises[0];
  if (last) await openExercise(last.id);
  if (location.hash.slice(1) !== meta.n) history.replaceState(null, '', '#' + meta.n);
}

// ── an exercise card in the lesson ──
function drawExercise(e) {
  const card = $(`[data-ex="${e.id}"]`); if (!card) return;
  const id = exId(e), shown = S.hints[id] || 0, hints = e.hints || [];
  const hasCheck = e.kind !== 'Read';
  const st = S.passed[id] ? '<span class="st ok">checks pass</span>' : S.checked[id] ? '<span class="st">checked, not all passing</span>' : '';
  card.classList.toggle('on', ex?.id === e.id);
  card.innerHTML = `<div class="h"><span class="id">${esc(id)}</span><span class="kind">${esc(e.kind)}</span><span class="t">${esc(e.title)}</span>${st}</div>
    <p>${e.task}</p>
    ${hints.slice(0, shown).map((h, i) => `<div class="hint"><b>Hint ${i + 1}.</b> ${h}</div>`).join('')}
    <div class="btns">
      <button class="btn primary" data-a="open">${ex?.id === e.id ? 'In the editor' : 'Open in the editor'}</button>
      ${shown < hints.length ? `<button class="btn" data-a="hint">Show a hint (${shown + 1} of ${hints.length})</button>` : ''}
      ${hasCheck ? `<button class="btn" data-a="sol" ${S.checked[id] ? '' : 'disabled title="Run Check once first"'}>Show the solution</button>` : ''}
    </div><div class="sol" hidden></div>`;
}
$('#lesson').addEventListener('click', async (ev) => {
  const b = ev.target.closest('button[data-a]'); if (!b) return;
  const card = b.closest('.ex'), e = lesson.exercises.find((x) => x.id === card.dataset.id), id = exId(e);
  if (b.dataset.a === 'open') openExercise(e.id);
  else if (b.dataset.a === 'hint') { S.hints[id] = (S.hints[id] || 0) + 1; save(); drawExercise(e); }
  else if (b.dataset.a === 'sol') {
    const box = card.querySelector('.sol');
    if (!box.hidden) { box.hidden = true; return; }
    box.innerHTML = `<pre>${esc(await text(`${lesson.n}/${e.id}.solution.go`))}</pre>`; box.hidden = false;
  }
});

// ── the editor ──
async function openExercise(eid) {
  const e = lesson.exercises.find((x) => x.id === eid), id = exId(e);
  if (ex && ex.id !== eid && !$('#run').disabled) stop();
  ex = e; S.last = id; save();
  editor.value = S.code[id] ?? await text(`${lesson.n}/${e.id}.go`);
  $('#fileName').textContent = `${id}.go`;
  $('#check').hidden = e.kind === 'Read';
  $('#problems').hidden = true; $('#out').textContent = ''; $('#runMeta').textContent = ''; status('');
  for (const x of lesson.exercises) drawExercise(x);
  edStatus();
}
function edStatus(st = editor?.status) {
  if (!st) return;
  const live = !toolchain.loaded ? '<span class="ed-dim">loading the Go compiler…</span>' : st.checking ? '<span class="ed-dim">checking…</span>' : '';
  $('#edStatus').innerHTML = `<button id="edProblems" title="Errors and warnings from the Go compiler (F8: next)"><span class="${st.errors ? 'ed-bad' : ''}">✕</span>${st.errors || 0}<span class="${st.warnings ? 'ed-warn' : ''}">⚠</span>${st.warnings || 0}</button>${live}<span class="sp"></span><span>Ln ${st.line}, Col ${st.col}</span><span>{ } Go</span><span>${esc(index?.go || '')}</span>`;
  $('#edCrumbs').innerHTML = ex ? `${esc(lesson.title)} <span class="c-sep">›</span> ${esc(ex.title)}${st.func ? ` <span class="c-sep">›</span> <span class="c-fn">${esc(st.func)}</span>` : ''}` : '';
}

// ── build, run, check ──
let busy = false;
function line(text, cls) { const s = document.createElement('span'); if (cls) s.className = cls; s.textContent = text; $('#out').append(s); }
async function buildWith(kind) {
  const extra = await text(`${lesson.n}/${ex.id}.${kind}.go`);
  status('Building…');
  toolchain.progress((what) => status(what === 'preparing' ? 'Starting the Go compiler…' : `Loading the Go ${what} (18 MB, the first time only)…`));
  const r = await toolchain.build({ 'main.go': editor.value, [`${kind}.go`]: extra });
  if (r.ok) { editor.setProblems([]); status(`Built in ${(r.ms / 1000).toFixed(1)} s`, 'ok'); return r.module; }
  status('The compiler refused the program', 'bad');
  const probs = []; $('#problems').hidden = false;
  $('#problems').innerHTML = '<div class="p-h"><span>Problems</span> · go build</div>' + r.errors.split('\n').filter(Boolean).map((l) => {
    const m = /^main\.go:(\d+):(\d+)?:?\s*(.*)/.exec(l); if (m) probs.push({ line: +m[1], col: +(m[2] || 1), msg: m[3] });
    return `<a ${m ? `data-l="${m[1]}" data-c="${m[2] || 1}"` : ''}>${esc(l.replace(/^(run|check)\.go:/, `$1.go (the exercise's ${kind === 'check' ? 'checks' : 'program'}):`))}</a>`;
  }).join('');
  editor.setProblems(probs); return null;
}
async function go(kind) {
  if (busy || !ex) return; if (kind === 'check' && ex.kind === 'Read') kind = 'run';
  busy = true; $('#run').disabled = $('#check').disabled = true; $('#problems').hidden = true; $('#out').textContent = ''; $('#runMeta').textContent = '';
  const module = await buildWith(kind);
  if (!module) { busy = false; $('#run').disabled = $('#check').disabled = false; return; }
  const id = exId(ex), e = ex; let pass = 0, fail = 0, buf = '';
  $('#stop').hidden = false; $('#runMeta').textContent = kind === 'check' ? 'checking…' : 'running…';
  const onOutput = kind === 'run' ? (t, fd) => line(t, fd === 2 ? 'fail' : '') : (t) => {
    buf += t; let i;
    while ((i = buf.indexOf('\n')) >= 0) {
      const l = buf.slice(0, i); buf = buf.slice(i + 1);
      if (l.startsWith('PASS ')) { pass++; line('✓ ' + l.slice(5) + '\n', 'pass'); } else if (l.startsWith('FAIL ')) { fail++; line('✗ ' + l.slice(5) + '\n', 'fail'); } else line(l + '\n', 'dim');
    }
  };
  run(module, { timeout: e.timeout || 10000, onOutput, onExit: (code, ms, msg) => {
    if (buf) line(buf, 'dim');
    busy = false; $('#run').disabled = $('#check').disabled = false; $('#stop').hidden = true;
    if (msg) line((kind === 'check' ? '✗ ' : '') + msg + (code === -1 ? ': the code under test takes far too long, or never ends' : '') + '\n', 'fail');
    if (kind === 'check') {
      const ok = code === 0 && fail === 0 && pass > 0;
      if (!ok && code !== 0) fail++;
      S.checked[id] = true; S.passed[id] = ok; save(); drawExercise(e);
      $('#runMeta').textContent = ok ? `all ${pass} checks pass` : `${pass} passed, ${fail} failed`;
    } else $('#runMeta').textContent = `exit status ${code} · ${ms} ms`;
  } });
}

// ── start ──
const api = await (await fetch('js/goapi.json')).json();
const editor = new Editor($('#editor'), { api, toolchain, toast, onBuild: () => go('run'), onStatus: edStatus,
  onChange: (v) => { if (!ex) return; S.code[exId(ex)] = v; clearTimeout(editor.t); editor.t = setTimeout(save, 400); } });
$('#run').addEventListener('click', () => go('run'));
$('#check').addEventListener('click', () => go('check'));
$('#stop').addEventListener('click', () => { if (stop()) { busy = false; $('#run').disabled = $('#check').disabled = false; $('#stop').hidden = true; $('#runMeta').textContent = 'stopped'; } });
$('#format').addEventListener('click', () => editor.format());
$('#reset').addEventListener('click', async () => { if (!ex || !confirm('Replace your code with the starting code of this exercise?')) return; delete S.code[exId(ex)]; save(); editor.value = await text(`${lesson.n}/${ex.id}.go`); });
$('#edStatus').addEventListener('click', (e) => { if (e.target.closest('#edProblems')) editor.nextProblem(); });
$('#problems').addEventListener('click', (e) => { const a = e.target.closest('a[data-l]'); if (a) editor.goto(+a.dataset.l, +a.dataset.c); });
index = await (await fetch('course/index.json')).json();
$('#nav').innerHTML = index.modules.map((m) => `<optgroup label="${esc(m.n + ' ' + m.title)}">${m.chapters.map((c) => c.lessons.map((l) => `<option value="${l.n}">${l.n} ${esc(l.title)}</option>`).join('')).join('')}</optgroup>`).join('');
$('#nav').addEventListener('change', () => openLesson($('#nav').value));
window.addEventListener('hashchange', () => { const n = location.hash.slice(1); if (n && n !== lesson?.n) openLesson(n); });
await openLesson(location.hash.slice(1) || S.lesson || '1.1.1');
toolchain.load().then(() => { editor.toolchainReady(); edStatus(); });
