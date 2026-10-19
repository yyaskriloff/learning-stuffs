---
name: scaffold-next-lab
description: Finds the next incomplete lab in SYLLABUS.md, creates its directory with README, stubs, tests, and notes, verifies the expected failing state, then commits and pushes. Use when the user asks to set up the next lab, sublesson, or lesson.
disable-model-invocation: true
---

# Scaffold the next lab

The learner writes every solution. This skill builds the lab around the work, never the work itself.

## 1. Find the next lab

- Read `SYLLABUS.md` top to bottom. A lab starts at `**Lab \`<name>\`:**`. Its checkboxes are the `- [ ]` / `- [x]` lines under it, up to the next lab or the lesson's `**Required edge cases:**` / `**Assessed by:**` lines.
- The target is the **first lab that still has an unchecked box**. Never skip ahead past an open lab.
- Path: the lesson's `Directory:` value plus the lab name, e.g. `phase-1/1-2-loops/a-sum-count-even/`.
- If that directory already exists, stop and tell the user. Do not overwrite.
- If the target is in an `exit/` directory, stop: exit problems must be unseen. Offer to write the instructor's exit problems and tests only if the user explicitly asks.

## 2. Classify the lab

| Lab kind | Signal | What to generate |
| --- | --- | --- |
| Implementation | Named functions to implement | Stub `.go` + full black-box `_test.go` |
| Test-writing | The checkboxes ask the learner to write or demonstrate tests (e.g. 1.1, 1.6 `b-aliasing`, 1.7) | Stub or supplied code + `_test.go` skeleton listing required cases as `// TODO` |
| Written | Deliverable is `notes.md` only (e.g. 1.3, 1.4) | Prompts/snippets in `README.md` + empty `notes.md` sections |

Lab 1.7 `a-faulty`: write the deliberately faulty implementation with one subtle bug. Do not write the regression test and do not hint at the bug in comments.

## 3. Write the files

**`README.md`** — copied from the syllabus, plus the lab's test command:

```markdown
# <lesson number> <lesson title> — Lab `<name>`

## Task
<task text and checkboxes, unchecked>

## Contract
<contract text, or "None beyond the signatures.">

## Required edge cases
<list>

## Assessed by
<types>

## Pass criteria
<text>

## If it fails
<remediation text>

## Run
go test ./<lab path>/ -v
```

**`<name>.go`** — package named after the lab (`package sumcounteven`, letters only). Exact signatures from the syllabus, each body `panic("unimplemented")`. Doc comments state the contract only when the syllabus fixes it; when the learner must decide the contract, leave a `// Contract: decide and document` line instead.

**`<name>_test.go`** for implementation labs:
- Table-driven, `t.Run` per case, one named case per required edge case plus normal cases.
- Black-box: assert outputs and documented behavior only. Do not encode one particular algorithm.
- If the lab is assessed by Algorithm correctness, add a randomized test against a simple brute-force reference written in the test file. Seed with a fixed value and log it.
- Where the learner must choose a contract (e.g. empty-slice result, `math.MinInt / -1`), leave that case as a `t.Skip("decide the contract, then write this case")`.
- Add `bench_test.go` only when the lab asks for a benchmark, and only as a skeleton with sizes listed.

**`notes.md`** — one heading per Complexity or Reasoning item, empty under each heading.

Never write a working implementation, an answer key, or completed `notes.md` content.

## 4. Verify the expected state

- If no `go.mod` exists at the repo root and this is lab 1.1, leave module creation to the learner (it is a checkbox). Skip the test run and say so.
- Otherwise run `go vet ./<lab path>/` (must pass) and `go test ./<lab path>/` (must fail with `unimplemented`, or pass with skips/TODOs for test-writing labs). A compile error means the scaffold is broken: fix the scaffold, not the stub.

## 5. Commit and push

- Stage only the new lab directory: `git add <lab path>`. Never `git add -A` or `.`; other files may hold unrelated work.
- Commit message: `Scaffold <lesson number> lab <name>`, e.g. `Scaffold 1.2 lab a-sum-count-even`.
- `git push origin HEAD`. If the push fails, report the error and stop.

## 6. Report

Tell the user the lab path, what they need to write, the test command, and suggest starting with the `teach-concept` skill.
