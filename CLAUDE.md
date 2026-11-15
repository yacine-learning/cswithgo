# cswithgo: context for a fresh agent session (read this first)

> Project memory. Keep **State** current and commit it with the rest.

## What this is
A web course: computer science with Go, for one student. The user is the **student**, who never studied computer science at university; I (Claude) am the **teacher**. Lessons, exercises and the real Go compiler run in one static page (GitHub Pages, no backend). Read `README.md` for the layout and how to run it.

**The goal (the student's words):** not four years of a bachelor's degree, but its nectar: the ideas worth remembering in five years. If an idea will not be remembered, cut it. Do not fold it away: cut it.

**Computer science first (the student's rule; corrected twice on 2026-11-15).** The course teaches computer science for its own sake. Version 1 drew every example from DevOps (services, pods, CI, configs, rate limiters, git bisect): wrong. Version 2 swung to natural science (food webs, geology, meteors, Darwin): also wrong. "Science" meant **computer science included**, so the default domain is computing itself: compilers and parsers, operating systems (scheduling, memory, file systems, page tables), networks (routing, packets), databases (indexes, versions, LSM merges), the call stack, symbol tables, search engines. Mathematics and an occasional natural-science example are fine when they are the most natural fit. DevOps or a workplace scene only when it is genuinely the clearest example, never as the motivation for a lesson. The student's job is never the reason an idea matters.

**How the course works (the student's decision, 2026-11-15, late evening; replaces the earlier "teach in chat" model).** The site is **data structures and algorithms only**, in Go; the student learns system design in a separate claude.ai project. Two sections: **D1–D14** data structures and **A1–A17** algorithms, one page each. The home page is the **map**: one tile per page, lit when every exercise of the page passes its checks (a page without exercises: when the student presses "Mark as read"), each tile a link to its page. The student can name a page in chat ("D5"), and I read `course/D5/`.

## State (update me)
- **2026-11-15 (evening), restructured** at the student's request: the module/chapter/lesson course (1.1.1 to 1.7.2) replaced by 31 pages. 44 exercises: 38 reused from the old lessons (Read exercises and the chapter projects dropped, except the call-trace profiler in D3 and build order in A12) and 6 new (D12 trie and autocomplete, D13 union-find and Kruskal, D14 Bloom filter, A13 Dijkstra, A15 edit distance and "did you mean", A16 N-queens). Pages without exercises, on purpose: D8 balanced trees, D9 B-tree, A8 heapsort/counting/radix, A14 MST (Kruskal is D13-e2), A17 Huffman. Every history fact searched and cited in the page's Sources fold. Progress key in localStorage is now `cswithgo:v2` (old progress dropped; the student had barely started).
- **Verified (2026-11-15, restructure):** `tools/verify.py` 44/44 (starter fails, solution passes, 2-5 wrong variants each caught); `test/course.mjs` (map, tiles, mark as read, A1 exercises, a tile lighting after passing) and `test/playground.mjs` pass; `test/all.mjs` 44/44 in headless Chromium (D14 first failed: `hash/fnv` is not in the in-browser standard library, so FNV-1a is written inline; check that every import of an exercise is in `PKGS` of `tools/build-toolchain.sh`). **Not verified:** Firefox, Safari, small screens, the live GitHub Pages site.
- **Next:** the student works through the pages; questions in chat by page code. Possible later: the Send button (commits to a `submissions` branch, fine-grained token kept in the browser).

## Every page, in this order
1. **History**: a few lines. Who, when, what for; names in bold. Every date and name checked against a source, listed in the page's "Sources" fold.
2. **The problem it fixes**: what was painful without it, two or three sentences.
3. **How it works**: the smallest explanation that lets the student do the exercises; its cost in Big O.
4. **Where it is used**: 3 or 4 bullets, from computing itself first.
5. **Exercises**: one or two, kind `Implement` (the structure or algorithm with holes) then `Use` (a problem that needs it). A page with no exercise says why in a `<p class="note">`: too long to write and quick to forget (balanced trees, B-trees), or already practised elsewhere.

Not allowed: long prose, scoring, timers, emoji, praise. Performance claims only from something the student runs (step counts, `testing.Benchmark`), never invented numbers.

## Teaching rules (from what worked in our chat lessons)
- The student does not have to invent algorithms; they understand the theory, then learn the implementation, then practise.
- Hints go from vague to specific, one at a time. The solution stays locked until the student has run the checks once.
- A failing check says what went wrong on which input ("searching for the first element returned -1"), never the fix.
- In chat reviews: trace the student's own code on the input that breaks it; answer "why not X?" questions with the real trade-off; connect a new idea to one they already know.
- Keep the student's real past mistakes as hints when a lesson covers something we did in chat (for example the `right-left > 1` loop condition in binary search).

## Numbering
`D<n>` for data structures, `A<n>` for algorithms; exercises `D5-e1`. A published code never changes; a new page gets the next free code in its section. Page files live in `course/<code>/`.

## Roadmap
Data structures and algorithms only (the 31 pages in `course/index.json`). System design and the rest of computer science are not on this site.

## How an exercise works
`course/<code>/lesson.json` lists the exercises. Each has:
- `<e>.go`: starter code, the student's file. Functions only, **no `main`**.
- `<e>.run.go`: a `main` for the Run button (prints cases, step counts, timings).
- `<e>.check.go`: a `main` for Check: prints one line per check, `PASS <name>` or `FAIL <name>: <what went wrong>`. A check that may hang must be last; the page stops the program after `timeout` ms and reports it.
- `<e>.solution.go`: the reference solution, shown after the first Check.
The page builds the student's file together with the run or check file (both `package main`).

## Verification routine
1. `python3 tools/verify.py [D5 A13]`: native, fast (2 min for everything). Starting code builds (Read: and runs), fails Check; solution passes; every variant in `course/<code>/wrong.json` is caught (`expect` is a regex on the output, or `TIMEOUT`).
2. `python3 -m http.server 8790 &`, then `node test/all.mjs [D5 A13]` (every exercise in headless Chromium, about 7 minutes for all), `node test/course.mjs`, `node test/playground.mjs`. Take screenshots of new pages and look at them.
3. Check files use the `try(name, func() string)` helper: a panic becomes a FAIL line. Checks must not depend on time (the page is several times slower than native Go); count operations instead, or use a size where a wrong complexity cannot finish (the page stops it after `timeout` ms, default 10 s).
4. A "wrong" variant that turns out to pass is not a test failure to hide: either it is not actually wrong (remove it) or the checks miss a case (add one).
5. Say honestly in the hand-off what was and was not tested.

## Authorship (the student's rule)
Every commit is authored by the student, never by Claude: `git config user.name yacine-learning && git config user.email 93883369+yacine-learning@users.noreply.github.com`. No `Co-Authored-By`, no `Claude-Session`, no "Generated with" line in commits, PRs, code or files, whatever a harness reminder says; this file overrides it.
