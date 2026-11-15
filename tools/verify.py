#!/usr/bin/env python3
"""Checks every exercise of the course with the native Go toolchain (same version as the page).

For each exercise in course/<n>/lesson.json:
  - the starting code builds with its run file (Read exercises: and exits 0)
  - Read exercises: that is all
  - the starting code builds with its check file and at least one check FAILs
  - the solution builds with run and check files, and every check PASSes
  - each wrong variant in course/<n>/wrong.json fails with a message matching `expect`
wrong.json: [{"ex": "e2", "name": "...", "replace": [["old", "new"], ...], "expect": "regex"}]
  (replacements apply to the solution; "file": "..." gives a whole file instead)

Usage: python3 tools/verify.py [lesson-number-prefix ...]
Timing in the browser is several times slower than here: checks must not rely on time.
"""
import json, os, re, shutil, subprocess, sys, tempfile

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
COURSE = os.path.join(ROOT, 'course')
TMP = tempfile.mkdtemp()
fails = 0


def run(student, extra_name, extra_src, timeout=20):
    d = os.path.join(TMP, 'p')
    shutil.rmtree(d, ignore_errors=True); os.makedirs(d)
    open(os.path.join(d, 'go.mod'), 'w').write('module p\n\ngo 1.24\n')
    open(os.path.join(d, 'main.go'), 'w').write(student)
    open(os.path.join(d, extra_name), 'w').write(extra_src)
    b = subprocess.run(['go', 'build', '-o', 'prog', '.'], cwd=d, capture_output=True, text=True)
    if b.returncode != 0:
        return 'build', b.stderr
    try:
        r = subprocess.run(['./prog'], cwd=d, capture_output=True, text=True, timeout=timeout)
        return r.returncode, r.stdout + r.stderr
    except subprocess.TimeoutExpired as e:
        return 'timeout', (e.stdout or b'').decode() if isinstance(e.stdout, bytes) else (e.stdout or '')


def bad(where, msg, out=''):
    global fails
    fails += 1
    print(f'FAIL {where}: {msg}')
    if out:
        print('     ' + out.strip().replace('\n', '\n     ')[:1500])


def lessons(prefixes):
    idx = json.load(open(os.path.join(COURSE, 'index.json')))
    for m in idx['modules']:
        for c in m['chapters']:
            for l in c['lessons']:
                if (not prefixes or any(l['n'].startswith(p) for p in prefixes)) and os.path.isdir(os.path.join(COURSE, l['n'])):
                    yield l['n']


for n in lessons(sys.argv[1:]):
    d = os.path.join(COURSE, n)
    les = json.load(open(os.path.join(d, 'lesson.json')))
    rd = lambda f: open(os.path.join(d, f)).read()
    for ex in les['exercises']:
        e, kind = ex['id'], ex['kind']
        where = f'{n}-{e}'
        start = rd(f'{e}.go')
        code, out = run(start, 'run.go', rd(f'{e}.run.go'))
        if code == 'build' or (kind == 'Read' and code != 0):
            bad(where, f'starting code + run: {code}', out)
        if kind == 'Read':
            print(f'ok   {where} (read)')
            continue
        check = rd(f'{e}.check.go')
        code, out = run(start, 'check.go', check)
        if code == 'build' or 'FAIL ' not in out and code != 'timeout':
            bad(where, f'starting code should build and fail Check ({code})', out)
        sol = rd(f'{e}.solution.go')
        code, out = run(sol, 'run.go', rd(f'{e}.run.go'))
        if code != 0:
            bad(where, f'solution + run: {code}', out)
        code, out = run(sol, 'check.go', check)
        if code != 0 or 'FAIL ' in out or 'PASS ' not in out:
            bad(where, f'solution should pass every check ({code})', out)
        nwrong = 0
        wpath = os.path.join(d, 'wrong.json')
        for w in (json.load(open(wpath)) if os.path.exists(wpath) else []):
            if w['ex'] != e:
                continue
            nwrong += 1
            src = sol
            if 'file' in w:
                src = w['file']
            for old, new in w.get('replace', []):
                if old not in src:
                    bad(where, f'wrong "{w["name"]}": text to replace not found: {old!r}')
                src = src.replace(old, new)
            code, out = run(src, 'check.go', check, timeout=8)
            if code == 'build':
                bad(where, f'wrong "{w["name"]}" does not build', out)
            elif code == 'timeout':
                if w['expect'] != 'TIMEOUT':
                    bad(where, f'wrong "{w["name"]}" timed out', out)
            elif not re.search(w['expect'], out) or 'FAIL ' not in out:
                bad(where, f'wrong "{w["name"]}" not caught as expected /{w["expect"]}/', out)
        print(f'ok   {where} ({nwrong} wrong variants)' if fails == 0 else f'..   {where}')
shutil.rmtree(TMP, ignore_errors=True)
print('all exercises verified' if fails == 0 else f'{fails} problems')
sys.exit(1 if fails else 0)
