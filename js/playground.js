// A single-file Go playground: the editor, the real Go compiler, and the program's output.
import { Editor } from './editor.js';
import { toolchain } from './toolchain.js';
import { run, stop } from './runner.js';
import { splitOutput } from './splitter.js';

const $ = (s) => document.querySelector(s);
const esc = (s) => String(s).replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]);
const KEY = 'cswithgo:playground';
const START = `package main

import "fmt"

func main() {
	fmt.Println("Hello from Go, compiled in your browser")
}
`;
const toast = (msg) => { const t = document.createElement('div'); t.className = 'toast'; t.textContent = msg; $('#toasts').append(t); setTimeout(() => t.remove(), 2600); };
const status = (msg, cls = '') => { const el = $('#buildStatus'); el.textContent = msg; el.className = 'build-status ' + cls; };
const store = (() => { try { return localStorage; } catch { return null; } })();

function edStatus(st = editor?.status) {
  if (!st) return;
  const live = !toolchain.loaded ? '<span class="ed-dim">loading the Go compiler…</span>' : st.checking ? '<span class="ed-dim">checking…</span>' : '';
  $('#edStatus').innerHTML = `<button id="edProblems" title="Errors and warnings from the Go compiler (F8: next)"><span class="${st.errors ? 'ed-bad' : ''}">✕</span>${st.errors || 0}<span class="${st.warnings ? 'ed-warn' : ''}">⚠</span>${st.warnings || 0}</button>${live}<span class="sp"></span><span>Ln ${st.line}, Col ${st.col}</span><span>{ } Go</span><span>go1.24</span>`;
  $('#edCrumbs').innerHTML = 'main.go' + (st.func ? ` <span class="c-sep">›</span> <span class="c-fn">${esc(st.func)}</span>` : '');
}

let busy = false;
async function buildAndRun() {
  if (busy) return; busy = true; $('#run').disabled = true; $('#problems').hidden = true;
  status('Building…'); $('#runMeta').textContent = ''; $('#out').textContent = '';
  toolchain.progress((what) => status(what === 'preparing' ? 'Starting the Go compiler…' : `Loading the Go ${what} (18 MB, the first time only)…`));
  const r = await toolchain.build({ 'main.go': editor.value });
  busy = false; $('#run').disabled = false;
  if (!r.ok) {
    status('The compiler refused the program', 'bad');
    const probs = [];
    $('#problems').hidden = false;
    $('#problems').innerHTML = '<div class="p-h"><span>Problems</span> · go build</div>' + r.errors.split('\n').filter(Boolean).map((l) => {
      const m = /^main\.go:(\d+):(\d+)?:?\s*(.*)/.exec(l); if (m) probs.push({ line: +m[1], col: +(m[2] || 1), msg: m[3] });
      return `<a ${m ? `data-l="${m[1]}" data-c="${m[2] || 1}"` : ''}>${esc(l)}</a>`;
    }).join('');
    editor.setProblems(probs); return;
  }
  editor.setProblems([]);
  status(`Built in ${(r.ms / 1000).toFixed(1)} s`, 'ok'); $('#stop').hidden = false; $('#runMeta').textContent = 'running…';
  run(r.module, {
    onOutput: (text, fd) => { const s = document.createElement('span'); if (fd === 2) s.className = 'err'; s.textContent = text; $('#out').append(s); },
    onExit: (code, ms, msg) => {
      $('#stop').hidden = true;
      if (msg) { const s = document.createElement('span'); s.className = 'err'; s.textContent = '\n' + msg; $('#out').append(s); }
      $('#runMeta').textContent = `exit status ${code} · ${ms} ms`;
    },
  });
}

const api = await (await fetch('js/goapi.json')).json();
const editor = new Editor($('#editor'), {
  api, toolchain, toast, onBuild: buildAndRun, onStatus: edStatus,
  onChange: (v) => { clearTimeout(editor.t); editor.t = setTimeout(() => store?.setItem(KEY, v), 400); },
});
editor.value = store?.getItem(KEY) || START;
$('#run').addEventListener('click', buildAndRun);
$('#stop').addEventListener('click', () => { if (stop()) { $('#stop').hidden = true; $('#runMeta').textContent = 'stopped'; } });
$('#format').addEventListener('click', () => editor.format());
$('#edStatus').addEventListener('click', (e) => { if (e.target.closest('#edProblems')) editor.nextProblem(); });
$('#problems').addEventListener('click', (e) => { const a = e.target.closest('a[data-l]'); if (a) editor.goto(+a.dataset.l, +a.dataset.c); });
edStatus();
splitOutput($('.vsc'), $('.run-out'));
toolchain.load().then(() => { editor.toolchainReady(); edStatus(); });
