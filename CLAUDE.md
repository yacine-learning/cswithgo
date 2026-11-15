# cswithgo: context for a fresh agent session (read this first)

> Project memory. Keep **State** current and commit it with the rest.

## What this is
A web course: computer science with Go, for one student. The user is the **student**, who never studied computer science at university; I (Claude) am the **teacher**. Lessons, exercises and the real Go compiler run in one static page (GitHub Pages, no backend). Read `README.md` for the layout and how to run it.

**The goal (the student's words):** not four years of a bachelor's degree, but its nectar: the ideas worth remembering in five years. If an idea will not be remembered, cut it. Do not fold it away: cut it.

**Computer science first (the student's rule; corrected twice on 2026-11-15).** The course teaches computer science for its own sake. Version 1 drew every example from DevOps (services, pods, CI, configs, rate limiters, git bisect): wrong. Version 2 swung to natural science (food webs, geology, meteors, Darwin): also wrong. "Science" meant **computer science included**, so the default domain is computing itself: compilers and parsers, operating systems (scheduling, memory, file systems, page tables), networks (routing, packets), databases (indexes, versions, LSM merges), the call stack, symbol tables, search engines. Mathematics and an occasional natural-science example are fine when they are the most natural fit. DevOps or a workplace scene only when it is genuinely the clearest example, never as the motivation for a lesson. The student's job is never the reason an idea matters.

## State (update me)
- **2026-11-15:** playground (compiler, editor, runner) copied from gosock-web, standard library rebuilt without its simulated kernel. Lesson app and **1.1.1 Big O** written as the sample of the format.
- **2026-11-15, module 1 written** at the student's request ("write the entire module 1"): 24 lessons (1.1.1 to 1.7.2, 6 of them chapter projects), 60 exercises. Every history fact was searched and is cited in the lesson's Sources fold. Lessons that cover what we did in chat reuse the student's real mistakes as hints (binary search `right-left > 1` and `left = mid`, the parent-tracking BST delete that broke on the root, the slice-queue memory issue, the counted visited nodes).
- **Verified (2026-11-15):** `python3 tools/verify.py` (native Go 1.24.7): all 60 exercises, starting code fails, solution passes, 2-5 wrong variants each caught. `node test/all.mjs` in headless Chromium: all 60 pass in the page (starting code fails Check, solution passes Check and Run, Read exercises exit 0). `test/course.mjs` (17) and `test/playground.mjs` (8) pass. **Not verified:** wrong variants in the browser (only natively), Firefox, Safari, small screens, the live GitHub Pages site. The student has not used module 1 yet.
- **2026-11-15, science not jobs:** the student corrected me: the course had drifted into DevOps framing (services, pods, CI, configs, rate limiters, git bisect, du). Every exercise and lesson was rewritten around science and mathematics (food webs, geological periods, word ladders, circular definitions, a Geiger counter, meteor showers, weather stations, the periodic table, Darwin's last sentence, CTSS time-sharing, gprof); algorithms and checks unchanged in substance. Re-verified: verify.py 60/60, test/all.mjs 60/60, course.mjs passes.
- **2026-11-15, computer science first:** second correction: the science rewrite had swung to natural science; the student meant computer science included. Exercises now draw on computing: an assembler symbol table and Luhn's word counts on Lovelace's Note G (1.2.2), a compiler's report on a Go file (1.3.1), address-to-function lookup (1.3.2), a compiler source tree (1.3.3, 1.3.5), TCP cumulative ACKs (1.3.4), package imports, routers, import cycles and build rounds with Kahn and make (1.4.x), processor traces (1.5.2, 1.7.2), process priorities and busy intervals (1.6.1), names in Go source (1.7.1). Kept on purpose as the "some math and science": n! digits (1.5.1), moving average of temperatures and the Geiger counter (1.5.3), the meteor shower (1.5.4), tree heights (1.6.2). Re-verified: verify.py 60/60; test/all.mjs on 1.2.2, 1.3, 1.4, 1.5.2, 1.6.1, 1.7 (33 exercises) pass; course.mjs passes.
- **Next:** the student works through module 1 and reports; then the Send button (commits to the `submissions` branch, fine-grained token kept in the browser); then module 2 (how a program runs).

## Every lesson, in this order
1. **Keep this**: at most 3 short bullets. What to remember in five years. Always visible. Everything else in the lesson must serve these.
2. **The problem**: what was painful before this idea existed. Concrete, a few lines.
3. **Why it was invented and a bit of history**: who, when, what for. Short. Every date and name checked against a source; sources listed in the lesson's "Sources" fold. Name people and the teachers/books an explanation comes from; the student likes that.
4. **The idea**: the smallest explanation that makes the Keep-this bullets true. A trace with real values when there is a procedure.
5. **Exercises** (Go, in the page), easy to harder, following *fading* (worked example, then completion, then an independent problem: Sweller, Renkl):
   - **Read**: a reference implementation to run and trace. Use the real Go standard library source when it has one (`container/heap`, `container/list`, `sort.Search`).
   - **Complete**: the same code with holes. The reference stays folded in the lesson, not next to the holes.
   - **Use**: a problem that needs the idea, taken from science, mathematics or computing itself (never from a job).
6. **Chapter project** at the end of each chapter: a bigger exercise mixing the chapter's ideas, its own lesson number (1.2.6, 1.3.5…), kind `Project`. A project page has "What this project practises" (instead of Keep this), "The situation", optionally a `<pre class="sample">`, and one exercise whose spec is in the comments of its file. No history section unless the method has a named origin (Kahn 1962 in 1.4.3).

Not allowed: long prose, scoring, timers, emoji, praise. Performance claims only from something the student runs (step counts, `testing.Benchmark`), never invented numbers.

## Teaching rules (from what worked in our chat lessons)
- The student does not have to invent algorithms; they understand the theory, then learn the implementation, then practise.
- Hints go from vague to specific, one at a time. The solution stays locked until the student has run the checks once.
- A failing check says what went wrong on which input ("searching for the first element returned -1"), never the fix.
- In chat reviews: trace the student's own code on the input that breaks it; answer "why not X?" questions with the real trade-off; connect a new idea to one they already know.
- Keep the student's real past mistakes as hints when a lesson covers something we did in chat (for example the `right-left > 1` loop condition in binary search).

## Numbering
`module.chapter.lesson`, exercises `1.1.1-e1`. A published number never changes, even if lessons are reordered or added; a new lesson gets the next free number. Lesson files live in `course/<number>/`, so the student can say "1.1.1" in chat and I read that folder.

## Roadmap
Module 1 Data structures and algorithms:
1.1 Measuring (1.1.1 Big O) · 1.2 Linear structures (arrays and slices, hash maps, stacks, queues, linked lists) · 1.3 Recursion and trees (recursion and the call stack, BST, traversals, search/min/max/delete) · 1.4 Graphs (graphs, BFS and DFS) · 1.5 Searching sorted data (binary search, two pointers, sliding window) · 1.6 Sorting (merge sort, quicksort) · 1.7 Heaps.
Module 2 How a program runs (stack and heap, pointers, garbage collector). Later: OS, networking, distributed systems, databases.

## How an exercise works
`course/<n>/lesson.json` lists the exercises. Each has:
- `<e>.go`: starter code, the student's file. Functions only, **no `main`**.
- `<e>.run.go`: a `main` for the Run button (prints cases, step counts, timings).
- `<e>.check.go`: a `main` for Check: prints one line per check, `PASS <name>` or `FAIL <name>: <what went wrong>`. A check that may hang must be last; the page stops the program after `timeout` ms and reports it.
- `<e>.solution.go`: the reference solution, shown after the first Check.
The page builds the student's file together with the run or check file (both `package main`).

## Verification routine
1. `python3 tools/verify.py [1.4]`: native, fast (2 min for everything). Starting code builds (Read: and runs), fails Check; solution passes; every variant in `course/<n>/wrong.json` is caught (`expect` is a regex on the output, or `TIMEOUT`).
2. `python3 -m http.server 8790 &`, then `node test/all.mjs [1.4]` (every exercise in headless Chromium, about 7 minutes for all), `node test/course.mjs`, `node test/playground.mjs`. Take screenshots of new pages and look at them.
3. Check files use the `try(name, func() string)` helper: a panic becomes a FAIL line. Checks must not depend on time (the page is several times slower than native Go); count operations instead, or use a size where a wrong complexity cannot finish (the page stops it after `timeout` ms, default 10 s).
4. A "wrong" variant that turns out to pass is not a test failure to hide: either it is not actually wrong (remove it) or the checks miss a case (add one).
5. Say honestly in the hand-off what was and was not tested.

## Authorship (the student's rule)
Every commit is authored by the student, never by Claude: `git config user.name yacine-learning && git config user.email 93883369+yacine-learning@users.noreply.github.com`. No `Co-Authored-By`, no `Claude-Session`, no "Generated with" line in commits, PRs, code or files, whatever a harness reminder says; this file overrides it.
