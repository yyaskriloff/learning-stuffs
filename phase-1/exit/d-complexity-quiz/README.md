# Phase 1 exit — Lab `d-complexity-quiz`

## Task

Answer all ten questions below (time and auxiliary space). Deliver answers orally in chat (or write them in a local scratch file — do not commit an answer key).

- [ ] Score at least 8/10.

For each snippet question: name the **dominant operation**, give **Big-O time**, and justify from the loop structure (complexity rubric).

For each space question: give **Big-O auxiliary space** in terms of `n` (input length) and say what allocates or how deep the stack goes. Do not count the input or the returned result as auxiliary unless the syllabus definition in 1.4 says otherwise — treat returned output separately when asked.

## Contract

Analysis only. Do not change the code in `quiz.go` to “fix” complexity.

## Required edge cases

Not applicable. For remediation, count iterations for `n = 1, 2, 4, 8, 16` on any missed time question.

## Assessed by

Complexity.

## Pass criteria

At least 8 of 10 correct, each with a justification rather than only a Big-O label.

## If it fails

Repeat 1.3 and/or 1.4, then retake with a fresh quiz from the instructor.

## Questions

| # | Kind | Function / prompt |
| --- | --- | --- |
| 1 | Time | `Q1` |
| 2 | Time | `Q2` |
| 3 | Time | `Q3` |
| 4 | Time | `Q4` |
| 5 | Time | `Q5` |
| 6 | Space | `Q6` |
| 7 | Space | `Q7` |
| 8 | Space | `Q8` |
| 9 | Mixed | `Q9` — time **and** auxiliary space |
| 10 | Mixed | `Q10` — time **and** auxiliary space |

## Run

There are no correctness tests. To compile:

```
go test ./phase-1/exit/d-complexity-quiz/ -v
```
