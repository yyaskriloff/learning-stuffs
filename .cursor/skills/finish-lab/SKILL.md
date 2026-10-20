---
name: finish-lab
description: Verifies a completed DSA/Go lab passes its instructor-provided tests, reviews the solution with improvement tips, grades oral Complexity/Reasoning answers against the syllabus rubrics, checks off met boxes in SYLLABUS.md, then commits and pushes. Use when the user says they finished a lab and wants it checked, reviewed, or committed.
disable-model-invocation: true
---

# Finish a lab

## 1. Verify

- Find the lab the user names, otherwise the one with uncommitted changes.
- Confirm the lab's `_test.go` is the instructor suite (complete cases for required edge cases). If required edge cases are missing from the suite, **add them yourself** before grading; do not ask the learner to write them.
- If the learner documented an open contract and matching cases are still `t.Skip`, **complete those tests** to match the documented contract, then re-run.
- Run `go vet ./<lab path>/` and `go test ./<lab path>/ -v -count=1`. Also run `-race` if the lab uses goroutines.
- If anything fails, or a test is still skipped for an undocumented contract, stop. Report the failure and switch to the `unstick-lab` workflow. Do not commit.
- Do not treat "learner didn't write tests" as a failure. The learner's deliverable is the implementation (and oral Complexity/Reasoning answers when the lab requires them). Do not require or create a `notes.md`.

## 2. Review the solution

Keep it short, as a quick review:

- **Verdict:** pass or not, against the lab's pass criteria.
- **What works:** one or two specific strengths.
- **Improvements:** at most three, most important first. Each names the code and says why (correctness risk, complexity, idiomatic Go, readability). Describe the improvement; do not rewrite their solution unless asked. Do not critique the learner for not authoring the test file.
- **Go tip:** one idiom or standard-library point relevant to the code, if there is one.

When the lab is assessed by Complexity or Reasoning, grade those orally in chat (complexity rubric 0–4; reasoning rubric criteria). Do not ask for a written `notes.md`. A green test suite does not satisfy either on its own — quiz the learner if they have not already explained in this session.

If a pass criterion fails, record it in `mistakes.md` at the repo root (create it if missing) with its date, lab, category (algorithm, Go, debugging, or reasoning), cause, and the remediation from the lab README.

## 3. Check off boxes

In `SYLLABUS.md`, change `- [ ]` to `- [x]` only for this lab's boxes whose criteria are met. Leave Complexity or Reasoning boxes unchecked until the oral answer passes. Do not touch other labs.

## 4. Commit and push

- Stage only the lab directory, `SYLLABUS.md`, and `mistakes.md` if changed. Include any test completions you made for open contracts or missing edge cases. Never `git add -A` or `.`.
- If `SYLLABUS.md` has unrelated uncommitted edits, ask before staging it.
- Commit message: `Complete <lesson number> lab <name>`, e.g. `Complete 1.2 lab a-sum-count-even`. If boxes remain open, use `Progress on <lesson number> lab <name>`.
- `git push origin HEAD`. If the push fails, report the error and stop.

## 5. Next step

Say which boxes remain open, if any. If the lab is fully checked, suggest the `scaffold-next-lab` skill.
