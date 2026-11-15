// Every exercise of the course in the browser (the real page, the real toolchain in WebAssembly):
// Read exercises run and exit 0; others: the starting code fails Check, the reference solution
// passes Check and its Run exits 0. Usage: node test/all.mjs [lesson-prefix…]   (server on :8790)
import { chromium } from 'playwright';
import { readFileSync, existsSync } from 'node:fs';
const BASE = process.env.URL || 'http://localhost:8790/';
const only = process.argv.slice(2);
const index = JSON.parse(readFileSync('course/index.json', 'utf8'));
const lessons = index.modules.flatMap((m) => m.chapters.flatMap((c) => c.lessons.map((l) => l.n)))
  .filter((n) => existsSync(`course/${n}/lesson.json`) && (!only.length || only.some((p) => n.startsWith(p))));
const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 1400, height: 900 } });
const errors = []; page.on('pageerror', (e) => errors.push(String(e)));
let failed = 0, done = 0;
async function load(n, eid, code) {
  await page.evaluate(([id, src]) => { const s = JSON.parse(localStorage.getItem('cswithgo:v1') || '{}'); s.code ||= {}; if (src == null) delete s.code[id]; else s.code[id] = src; s.last = id; localStorage.setItem('cswithgo:v1', JSON.stringify(s)); }, [`${n}-${eid}`, code]);
  await page.goto(BASE + '#' + n); await page.reload();
  await page.waitForSelector(`.ex.on[data-id="${eid}"]`);
}
async function press(button) {
  await page.click('#' + button);
  await page.waitForFunction(() => /exit status|checks pass|failed|refused|stopped/.test(document.querySelector('#runMeta').textContent + document.querySelector('#buildStatus').textContent), null, { timeout: 180000 });
  return { out: await page.textContent('#out'), meta: await page.textContent('#runMeta'), status: await page.textContent('#buildStatus') };
}
await page.goto(BASE);
for (const n of lessons) {
  const les = JSON.parse(readFileSync(`course/${n}/lesson.json`, 'utf8'));
  for (const e of les.exercises) {
    const id = `${n}-${e.id}`, t0 = Date.now(), problems = [];
    try {
      if (e.kind === 'Read') {
        await load(n, e.id, null);
        const r = await press('run');
        if (!r.meta.includes('exit status 0')) problems.push('Run: ' + r.meta + ' ' + r.status + '\n' + r.out.slice(0, 400));
      } else {
        await load(n, e.id, null);
        let r = await press('check');
        if (!/failed/.test(r.meta)) problems.push('starting code should fail Check: ' + r.meta + ' ' + r.status);
        await load(n, e.id, readFileSync(`course/${n}/${e.id}.solution.go`, 'utf8'));
        r = await press('check');
        if (!/^all \d+ checks pass$/.test(r.meta)) problems.push('solution should pass Check: ' + r.meta + ' ' + r.status + '\n' + r.out.slice(0, 600));
        r = await press('run');
        if (!r.meta.includes('exit status 0')) problems.push('solution Run: ' + r.meta + ' ' + r.status + '\n' + r.out.slice(0, 400));
      }
    } catch (err) { problems.push(String(err).split('\n')[0]); }
    done++;
    if (problems.length) { failed++; console.log('FAIL', id, '\n     ' + problems.join('\n     ')); }
    else console.log('ok  ', id, `(${((Date.now() - t0) / 1000).toFixed(1)} s)`);
  }
}
if (errors.length) { failed++; console.log('FAIL page errors:', errors); }
await browser.close();
console.log(failed ? `${failed} of ${done} failed` : `all ${done} exercises pass in the browser`);
process.exit(failed ? 1 : 0);
