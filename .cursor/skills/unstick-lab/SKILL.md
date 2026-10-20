---
name: unstick-lab
description: Reviews the learner's in-progress lab code and failing instructor-provided tests to find where they are stuck, then gives the smallest hint that moves them forward without fixing the solution. Use when the user is stuck on a lab, asks why a test fails, or asks a question about their current lab work.
---

# Unstick the current lab

## Instructor contract

- Never edit the learner's solution `.go` files.
- Do not ask the learner to rewrite or weaken the provided `_test.go` to make a wrong solution pass.
- You may fix a broken scaffold test only if it is clearly wrong relative to the syllabus contract (wrong expected value, missing required edge case). Prefer telling the learner the test was fixed rather than blaming their solution.
- Give the smallest hint that unblocks the next step.
- Mistakes are diagnostic: name the category of the miss.

## Workflow

1. **Find the lab.** Use the one the user names, otherwise the lab with the most recently modified files (`git status`), otherwise the first lab in `SYLLABUS.md` with an unchecked box.
2. **Read** the lab `README.md`, the learner's implementation, and the **provided** tests.
3. **Get evidence.** Run `go test ./<lab path>/ -v`. For a compile error, read it before the code. If tests pass but the user is still stuck, ask what they expected to happen.
4. **Diagnose** the miss as one of:
   - **Algorithm** — the idea is wrong or incomplete.
   - **Go** — the idea is right but the language behaves differently than expected (slice aliasing, integer overflow or truncation, nil vs empty, map iteration order, range variable copies).
   - **Debugging** — the code may be right but the learner is misreading the failing test output or the contract.
   - **Reasoning** — they can't explain why their code should work.
5. **Respond with:**
   - The category and the failing case, quoted from the test output.
   - One hint from the ladder below, starting at rung 1 unless they've already tried it:
     1. A question about the failing case ("what is `i` when the loop exits on `[5]`?").
     2. Ask them to trace the state by hand on the smallest failing input.
     3. Point to the region of code (function or line range) where the state goes wrong.
     4. Name the bug, without writing the fix.
6. **Answer direct questions** about Go or the concept straight, then tie them back to their code with a question.
7. If the same miss repeats, suggest they record it in `mistakes.md` with its category.
