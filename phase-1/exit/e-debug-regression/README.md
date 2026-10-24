# Phase 1 exit — Lab `e-debug-regression`

## Task

The instructor supplies a deliberately faulty implementation and a failing test suite.

- [ ] Run the provided tests and reproduce the failure.
- [ ] Identify the smallest useful failing input from the test output.
- [ ] Fix the underlying defect in the implementation.
- [ ] Confirm the provided regression and boundary tests pass.
- [ ] Explain why the original implementation failed and how the regression test prevents that bug from returning.

## Contract

`CumulativeMax(nums []int) []int` returns a new slice of the same length where `out[i]` is the maximum among `nums[0]` through `nums[i]` (inclusive).

- A nil or empty input returns a result with `len` 0.
- Do not modify `nums`.

## Required edge cases

- empty / nil input
- single element
- strictly increasing / decreasing
- all equal
- negatives (including a maximum that is negative)
- input left unmodified

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

Tests pass, and you can explain the original defect and why the failing case would catch a regression.

## If it fails

Write a hypothesis before each edit, change one thing at a time, then retry. If needed, redo 1.7 first.

## Run

```
go test ./phase-1/exit/e-debug-regression/ -v
```
