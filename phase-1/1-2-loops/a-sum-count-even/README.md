# 1.2 Loops and basic state — Lab `a-sum-count-even`

## Task

Implement:

- [ ] `Sum(nums []int) int`
- [ ] `CountEven(nums []int) int`
- [ ] Explain what changes during each loop iteration and why the final result is correct.

## Contract

None beyond the signatures.

## Required edge cases

- empty slice
- one element
- all even
- all odd
- mixed values
- negative values

## Assessed by

Implementation, Edge cases, Complexity, Reasoning.

## Pass criteria

All tests pass, and `notes.md` explains why both algorithms take O(n) time and O(1) auxiliary space, scoring 3/4 or better on the complexity rubric.

## If it fails

Trace both loops by hand on `[3, -2, 0]` and write down the state after each iteration, then retry the explanation.

## Run

```
go test ./phase-1/1-2-loops/a-sum-count-even/ -v
```
