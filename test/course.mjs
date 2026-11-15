// Browser tests of the course page and lesson exercises.
// For each exercise with checks: the starting code fails, the reference solution passes,
// and plausible wrong answers are each caught with a message.
// Needs a server on :8790 (python3 -m http.server 8790) and Playwright with Chromium.
import { chromium } from 'playwright';
import { readFileSync } from 'node:fs';
const BASE = process.env.URL || 'http://localhost:8790/';
const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 1400, height: 860 } });
const errors = []; page.on('pageerror', (e) => errors.push(String(e)));
let failed = 0;
const test = async (name, fn) => { try { await fn(); console.log('ok  ', name); } catch (e) { failed++; console.log('FAIL', name, '\n     ', e.message.split('\n').slice(0, 6).join('\n      ')); } };
const file = (p) => readFileSync('course/' + p, 'utf8');

// put code in an exercise (through the saved state), reload, open it
async function load(lesson, eid, code) {
  await page.evaluate(([id, src]) => { const s = JSON.parse(localStorage.getItem('cswithgo:v2') || '{}'); s.code ||= {}; if (src == null) delete s.code[id]; else s.code[id] = src; s.last = id; localStorage.setItem('cswithgo:v2', JSON.stringify(s)); }, [`${lesson}-${eid}`, code]);
  await page.goto(BASE + '#' + lesson); await page.reload();
  await page.waitForSelector(`.ex.on[data-id="${eid}"]`);
}
async function press(button, timeout = 120000) {
  await page.click('#' + button);
  await page.waitForFunction(() => /exit status|checks pass|failed|refused|stopped/.test(document.querySelector('#runMeta').textContent + document.querySelector('#buildStatus').textContent), null, { timeout });
  return { out: await page.textContent('#out'), meta: await page.textContent('#runMeta'), status: await page.textContent('#buildStatus'), problems: await page.textContent('#problems') };
}
const expectPass = (r) => { if (!/^all \d+ checks pass$/.test(r.meta)) throw new Error('expected all checks to pass:\n' + r.meta + '\n' + r.out); };
const expectFail = (r, re) => { if (!/failed/.test(r.meta) || !re.test(r.out)) throw new Error(`expected a failure matching ${re}:\n${r.meta}\n${r.out}`); };

await page.goto(BASE);
await test('the home page is the map: 14 + 17 tiles, none lit', async () => {
  await page.evaluate(() => localStorage.removeItem('cswithgo:v2')); await page.goto(BASE); await page.reload();
  await page.waitForSelector('.tile');
  const n = (await page.$$('.tiles')).map(() => 0).length, tiles = await page.$$('.tile');
  if (n !== 2 || tiles.length !== 31) throw new Error(`${n} sections, ${tiles.length} tiles`);
  if ((await page.$$('.tile.done')).length) throw new Error('a tile is lit on a fresh start');
  if (!(await page.textContent('.progress')).includes('0 of 31')) throw new Error(await page.textContent('.progress'));
});
await test('a tile opens its page; Map goes back', async () => {
  await page.click('.tile[href="#D5"]'); await page.waitForSelector('.lesson h1 .no');
  const h = await page.textContent('.lesson h1'); if (!h.includes('D5') || !h.includes('Hash table')) throw new Error(h);
  for (const t of ['History', 'The problem it fixes', 'Where it is used']) if (!(await page.textContent('.lesson')).includes(t)) throw new Error('missing ' + t);
  if ((await page.$$('.ex')).length !== 2) throw new Error('expected 2 exercise cards');
  await page.click('.crumb a'); await page.waitForSelector('.tile');
});
await test('a page without exercises: no editor, Mark as read lights its tile', async () => {
  await page.goto(BASE + '#D9'); await page.waitForSelector('.mark button');
  if (await page.isVisible('.bench')) throw new Error('the editor is shown on a reading page');
  await page.click('.mark button'); await page.waitForSelector('.mark button:not(.primary)');
  await page.goto(BASE); await page.waitForSelector('.tile');
  if (!(await page.$('.tile.done[href="#D9"]'))) throw new Error('D9 tile not lit');
  if (!(await page.textContent('.progress')).includes('1 of 31')) throw new Error(await page.textContent('.progress'));
});
await test('A1-e1 the starting code compiles and fails Check', async () => {
  await load('A1', 'e1', null);
  expectFail(await press('check'), /duplicate next to each other: HasDuplicate\(\[4 4\]\) returned false, expected true/);
});
await test('A1-e1 the solution is locked until Check ran, then opens', async () => {
  await page.evaluate(() => { localStorage.removeItem('cswithgo:v2'); }); await load('A1', 'e1', null);
  if (!(await page.isDisabled('.ex[data-id="e1"] [data-a="sol"]'))) throw new Error('solution button should be disabled');
  await press('check');
  await page.click('.ex[data-id="e1"] [data-a="sol"]');
  await page.waitForFunction(() => document.querySelector('.ex[data-id="e1"] .sol').textContent.includes('seen[x] = true'), null, { timeout: 5000 });
});
await test('A1-e1 hints appear one at a time', async () => {
  await page.click('.ex[data-id="e1"] [data-a="hint"]');
  await page.click('.ex[data-id="e1"] [data-a="hint"]');
  const n = (await page.$$('.ex[data-id="e1"] .hint')).length; if (n !== 2) throw new Error(n + ' hints shown');
});
await test('A1-e1 the solution passes, Run shows slow against fast', async () => {
  await load('A1', 'e1', file('A1/e1.solution.go'));
  expectPass(await press('check'));
  const r = await press('run'); if (!/2000000\s+\d+ms/.test(r.out)) throw new Error(r.out);
});
const e2 = (body) => `package main\n\nfunc HasDuplicate(xs []int) bool {\n${body}\n}\n`;
await test('A1-e1 wrong: marks before checking (always true)', async () => {
  await load('A1', 'e1', e2('\tseen := map[int]bool{}\n\tfor _, x := range xs {\n\t\tseen[x] = true\n\t\tif seen[x] {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false'));
  expectFail(await press('check'), /one value: HasDuplicate\(\[7\]\) returned true, expected false/);
});
await test('A1-e1 wrong: the nested loop is caught as too slow', async () => {
  await load('A1', 'e1', e2('\tfor i := range xs {\n\t\tfor j := i + 1; j < len(xs); j++ {\n\t\t\tif xs[i] == xs[j] {\n\t\t\t\treturn true\n\t\t\t}\n\t\t}\n\t}\n\treturn false'));
  expectFail(await press('check'), /stopped after 8 s: the code under test takes far too long|took \d/);
});
await test('A1-e1 wrong: map[int]int counting from zero misses zero', async () => {
  await load('A1', 'e1', e2('\tseen := map[int]int{}\n\tfor i, x := range xs {\n\t\tif seen[x] != 0 {\n\t\t\treturn true\n\t\t}\n\t\tseen[x] = i\n\t}\n\treturn false'));
  expectFail(await press('check'), /duplicate next to each other|duplicate far apart|zero and negative/);
});
await test('A1-e1 a compiler error is shown on the student file', async () => {
  await load('A1', 'e1', e2('\tx := 1\n\treturn false'));
  const r = await press('check'); if (!r.status.includes('refused') || !/main\.go:4:\d+: declared and not used: x/.test(r.problems)) throw new Error(JSON.stringify(r));
});
await test('A1-e1 renaming the function names the check file in the error', async () => {
  await load('A1', 'e1', 'package main\n\nfunc HasDup(xs []int) bool { return false }\n');
  const r = await press('check'); if (!/check\.go \(the exercise's checks\).*undefined: HasDuplicate/.test(r.problems)) throw new Error(r.problems);
});
await test('A1-e2 the starting code fails Check with "no answer yet"', async () => {
  await load('A1', 'e2', null);
  expectFail(await press('check'), /A: no answer yet/);
});
await test('A1-e2 Run shows the doubling ratios of the four functions', async () => {
  const r = await press('run'); if (!/C .*×4\.00/.test(r.out) || !/D .*×1\.00/.test(r.out)) throw new Error(r.out);
});
await test('A1-e2 the solution passes', async () => {
  await load('A1', 'e2', file('A1/e2.solution.go'));
  expectPass(await press('check'));
});
await test('A1-e2 wrong: B taken for quadratic (nested loop trap)', async () => {
  await load('A1', 'e2', file('A1/e2.solution.go').replace('"B": Linear', '"B": Quadratic'));
  expectFail(await press('check'), /B: you said O\(n²\)/);
});
await test('A1-e2 wrong: D taken for linear (bounded loop trap)', async () => {
  await load('A1', 'e2', file('A1/e2.solution.go').replace('"D": Constant', '"D": Linear'));
  expectFail(await press('check'), /D: you said O\(n\)/);
});
await test('progress and code survive a reload', async () => {
  await page.reload(); await page.waitForSelector('.ex[data-id="e2"]');
  if (!(await page.textContent('.ex[data-id="e1"] .st')).includes('checked')) throw new Error('e1 status lost');
  if (!(await page.textContent('.cm-content')).includes('"D": Linear')) throw new Error('e2 code lost');
});
await test('passing every exercise of a page lights its tile', async () => {
  await load('A1', 'e1', file('A1/e1.solution.go')); expectPass(await press('check'));
  await load('A1', 'e2', file('A1/e2.solution.go')); expectPass(await press('check'));
  await page.goto(BASE); await page.waitForSelector('.tile');
  if (!(await page.$('.tile.done[href="#A1"]'))) throw new Error('A1 tile not lit');
});
if (process.env.SHOT) { await load('A1', 'e1', null); await press('check'); await page.screenshot({ path: process.env.SHOT }); }
if (errors.length) { failed++; console.log('FAIL page errors:', errors); }
await browser.close();
console.log(failed ? `${failed} failed` : 'all passed');
process.exit(failed ? 1 : 0);
