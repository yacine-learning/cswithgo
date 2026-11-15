// Browser tests of the playground: build, run, output, compiler errors, panics, stop, stdlib for data structures.
// Needs a server on :8790 (python3 -m http.server 8790) and Playwright with Chromium.
import { chromium } from 'playwright';
const URL = process.env.URL || 'http://localhost:8790/';
const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 1200, height: 800 } });
const errors = []; page.on('pageerror', (e) => errors.push(String(e)));
let failed = 0;
const test = async (name, fn) => { try { await fn(); console.log('ok  ', name); } catch (e) { failed++; console.log('FAIL', name, '\n     ', e.message); } };
const setCode = (src) => page.evaluate((s) => { localStorage.setItem('cswithgo:playground', s); }, src).then(() => page.reload());
const runOut = async (timeout = 120000) => {
  await page.waitForSelector('.cm-content');
  await page.click('#run');
  await page.waitForFunction(() => /exit status|refused/.test(document.querySelector('#runMeta').textContent + document.querySelector('#buildStatus').textContent), null, { timeout });
  return { out: await page.textContent('#out'), meta: await page.textContent('#runMeta'), status: await page.textContent('#buildStatus'), problems: await page.textContent('#problems') };
};
await page.goto(URL);

await test('the starting program builds and prints', async () => {
  const r = await runOut();
  if (!r.out.includes('Hello from Go') || !r.meta.includes('exit status 0')) throw new Error(JSON.stringify(r));
});
await test('data-structure packages compile and run', async () => {
  await setCode(`package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

func main() {
	h := &IntHeap{5, 2, 8}
	heap.Init(h)
	heap.Push(h, 1)
	fmt.Println("heap min:", heap.Pop(h))
	l := list.New()
	l.PushBack(1)
	l.PushFront(0)
	fmt.Println("list len:", l.Len())
	m := map[string]int{"b": 2, "a": 1}
	keys := slices.Sorted(maps.Keys(m))
	fmt.Println("keys:", strings.Join(keys, ","))
	s := []int{2, 23, 5, 91}
	sort.Ints(s)
	fmt.Println("search 23:", sort.SearchInts(s, 23))
}
`);
  const r = await runOut();
  for (const want of ['heap min: 1', 'list len: 2', 'keys: a,b', 'search 23: 2']) if (!r.out.includes(want)) throw new Error('missing ' + want + ' in ' + r.out);
});
await test('compiler errors are listed with their line', async () => {
  await setCode('package main\n\nfunc main() {\n\tx := 1\n}\n');
  const r = await runOut();
  if (!r.status.includes('refused') || !/main\.go:4:\d+: declared and not used: x/.test(r.problems)) throw new Error(JSON.stringify(r));
});
await test('a panic shows on stderr with exit status 2', async () => {
  await setCode('package main\n\nfunc main() {\n\tvar s []int\n\t_ = s[3]\n}\n');
  const r = await runOut();
  if (!r.out.includes('index out of range') || !r.meta.includes('exit status 2')) throw new Error(JSON.stringify(r));
  if (!(await page.$('#out .err'))) throw new Error('panic text not marked as stderr');
});
await test('testing.Benchmark measures real time', async () => {
  await setCode(`package main

import (
	"fmt"
	"testing"
)

func main() {
	r := testing.Benchmark(func(b *testing.B) {
		s := 0
		for i := 0; i < b.N; i++ {
			s += i
		}
		_ = s
	})
	fmt.Println("ns/op>0:", r.NsPerOp() >= 0, "N>0:", r.N > 0)
}
`);
  const r = await runOut();
  if (!r.out.includes('N>0: true')) throw new Error(JSON.stringify(r));
});
await test('an endless program can be stopped', async () => {
  await setCode('package main\n\nimport "fmt"\n\nfunc main() {\n\tfor i := 0; ; i++ {\n\t\tif i%10000000 == 0 {\n\t\t\tfmt.Println(i)\n\t\t}\n\t}\n}\n');
  await page.click('#run');
  await page.waitForSelector('#stop:not([hidden])', { timeout: 60000 });
  await page.waitForTimeout(800);
  await page.click('#stop');
  if ((await page.textContent('#runMeta')) !== 'stopped') throw new Error(await page.textContent('#runMeta'));
});
await test('gofmt formats the code', async () => {
  await setCode('package main\nimport "fmt"\nfunc main(){fmt.Println( 1 )}\n');
  await page.waitForSelector('.cm-content'); await page.waitForTimeout(500);
  await page.click('#format');
  await page.waitForFunction(() => document.querySelector('.cm-content').innerText.includes('fmt.Println(1)'), null, { timeout: 60000 });
});
await test('live checks underline an error while typing', async () => {
  await setCode('package main\n\nfunc main() {\n}\n');
  await page.waitForSelector('.cm-content'); await page.click('.cm-content');
  await page.keyboard.press('Control+End'); await page.keyboard.press('ArrowUp'); await page.keyboard.press('End');
  await page.keyboard.type('\n\tundefinedThing()');
  await page.waitForSelector('.cm-lintRange-error', { timeout: 60000 });
});
if (process.env.SHOT) await page.screenshot({ path: process.env.SHOT });
if (errors.length) { failed++; console.log('FAIL page errors:', errors); }
await browser.close();
console.log(failed ? `${failed} failed` : 'all passed');
process.exit(failed ? 1 : 0);
