// The course: a home page with every data structure and algorithm as a tile (lit when done), and one page
// per tile: the text on the left; on the right the editor, Run, Check and the output.
import { Editor } from './editor.js';
import { toolchain } from './toolchain.js';
import { run, stop } from './runner.js';

const $ = (s) => document.querySelector(s);
const esc = (s) => String(s).replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]);
const KEY = 'cswithgo:v2';
const S = (() => { try { return JSON.parse(localStorage.getItem(KEY)) || {}; } catch { return {}; } })();
for (const k of ['code', 'checked', 'passed', 'hints', 'done']) S[k] ||= {};
const save = () => { try { localStorage.setItem(KEY, JSON.stringify(S)); } catch {} };
const toast = (msg) => { const t = document.createElement('div'); t.className = 'toast'; t.textContent = msg; $('#toasts').append(t); setTimeout(() => t.remove(), 2600); };
const status = (msg, cls = '') => { const el = $('#buildStatus'); el.textContent = msg; el.className = 'build-status ' + cls; };
const files = new Map();
const text = (path) => files.get(path) ?? fetch('course/' + path).then((r) => (r.ok ? r.text() : null)).then((t) => (files.set(path, t), t));

let index, lesson, ex = null;          // ex: the exercise in the editor
const exId = (e) => `${lesson.n}-${e.id}`;

// ── navigation ──
const pages = () => index.sections.flatMap((sec) => sec.pages.map((p) => ({ ...p, section: sec.title })));
const lessons = new Map();
const getLesson = (n) => lessons.get(n) || fetch(`course/${n}/lesson.json`).then((r) => r.json()).then((l) => (lessons.set(n, l), l));
// a page is done when every exercise passes its checks; a page without exercises when marked as read
const isDone = (n, l) => (l.exercises.length ? l.exercises.every((e) => S.passed[`${n}-${e.id}`]) : !!S.done[n]);

async function openHome() {
  lesson = null; ex = null; document.body.classList.add('home'); $('#nav').value = ''; S.lesson = ''; save();
  const ls = await Promise.all(pages().map((p) => getLesson(p.n)));
  const all = pages(); let done = 0;
  const tiles = (sec) => sec.pages.map((p) => {
    const l = ls[all.findIndex((x) => x.n === p.n)], d = isDone(p.n, l); if (d) done++;
    const k = l.exercises.length, passed = l.exercises.filter((e) => S.passed[`${p.n}-${e.id}`]).length;
    const st = d ? 'done' : k ? `${passed} of ${k} exercise${k > 1 ? 's' : ''}` : 'reading';
    return `<a class="tile${d ? ' done' : ''}" href="#${p.n}"><span class="code">${esc(p.n)}</span><b>${esc(p.title)}</b><small>${esc(p.tag)}</small><span class="st">${st}</span></a>`;
  }).join('');
  const body = index.sections.map((sec) => `<h2>${esc(sec.title)}</h2><div class="tiles">${tiles(sec)}</div>`).join('');
  $('#lesson').innerHTML = `<h1>Data structures and algorithms, in Go</h1><p class="intro">Every structure and algorithm of the course. A tile lights up when you have passed all its exercises (or marked a page without exercises as read). Click one to open it.</p><p class="progress"><b>${done}</b> of ${all.length} done</p>` + body;
  $('#lesson').scrollTop = 0; document.title = 'cswithgo';
  if (location.hash) history.replaceState(null, '', location.pathname);
}

async function openLesson(n) {
  const list = pages(), meta = list.find((l) => l.n === n);
  if (!meta) return openHome();
  lesson = { ...(await getLesson(meta.n)), n: meta.n };
  const html = await text(`${meta.n}/lesson.html`);
  $('#nav').value = meta.n; S.lesson = meta.n; save();
  document.body.classList.remove('home'); document.body.classList.toggle('read', !lesson.exercises.length);
  const i = list.indexOf(meta), prev = list[i - 1], next = list[i + 1];
  const link = (l, dir) => l ? `<a class="${dir}" href="#${l.n}"><small>${dir === 'prev' ? '← Previous' : 'Next →'}</small>${esc(l.n)} ${esc(l.title)}</a>` : '<span></span>';
  const mark = lesson.exercises.length ? '' : `<div class="mark"><button class="btn${S.done[meta.n] ? '' : ' primary'}" data-a="done">${S.done[meta.n] ? 'Read ✓ (click to undo)' : 'Mark as read'}</button></div>`;
  $('#lesson').innerHTML = `<div class="crumb"><a href="#">Map</a> › ${esc(meta.section)}</div><h1><span class="no">${esc(meta.n)}</span>${esc(meta.title)}</h1>` + (html || '')
    + mark + `<nav class="lnav">${link(prev, 'prev')}${link(next, 'next')}</nav>`;
  document.title = `${meta.n} ${meta.title} · cswithgo`;
  for (const e of lesson.exercises) { const slot = $(`[data-ex="${e.id}"]`); if (slot) { slot.className = 'ex'; slot.dataset.id = e.id; drawExercise(e); } }
  $('#lesson').scrollTop = 0;
  const last = lesson.exercises.find((e) => exId(e) === S.last) || lesson.exercises[0];
  ex = null; if (last) await openExercise(last.id);
  if (location.hash.slice(1) !== meta.n) history.replaceState(null, '', '#' + meta.n);
}

// ── an exercise card in the lesson ──
function drawExercise(e) {
  const card = $(`[data-ex="${e.id}"]`); if (!card) return;
  const id = exId(e), shown = S.hints[id] || 0, hints = e.hints || [];
  const hasCheck = e.kind !== 'Read';
  const st = S.passed[id] ? '<span class="st ok">checks pass</span>' : S.checked[id] ? '<span class="st">checked, not all passing</span>' : '';
  card.classList.toggle('on', ex?.id === e.id);
  card.innerHTML = `<div class="h"><span class="id">${esc(id)}</span><span class="kind" data-k="${esc(e.kind)}">${esc(e.kind)}</span><span class="t">${esc(e.title)}</span>${st}</div>
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
  if (b.dataset.a === 'done') { S.done[lesson.n] = !S.done[lesson.n]; save(); return openLesson(lesson.n); }
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
      const was = isDone(lesson.n, lesson);
      S.checked[id] = true; S.passed[id] = ok; save(); drawExercise(e);
      if (!was && isDone(lesson.n, lesson)) toast(`${lesson.n} ${lesson.title}: every exercise passes. Its tile is lit on the map.`);
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
$('#nav').innerHTML = '<option value="">Map of the course</option>' + index.sections.map((sec) => `<optgroup label="${esc(sec.title)}">${sec.pages.map((p) => `<option value="${p.n}">${p.n} ${esc(p.title)}</option>`).join('')}</optgroup>`).join('');
$('#nav').addEventListener('change', () => { const n = $('#nav').value; n ? openLesson(n) : openHome(); });
const route = () => { const n = location.hash.slice(1); if (!n) return openHome(); if (n !== lesson?.n) return openLesson(n); };
window.addEventListener('hashchange', route);
await route();
toolchain.load().then(() => { editor.toolchainReady(); edStatus(); });
