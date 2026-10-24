# Phase 1 exit — Lab `a-count-in-range`

## Task

- [ ] Implement `CountInRange(nums []int, lo, hi int) int`.

## Contract

`CountInRange` returns how many elements of `nums` fall in the closed interval `[lo, hi]` (inclusive on both ends).

- If `lo > hi`, return `0`.
- A nil or empty `nums` returns `0`.
- Do not modify `nums`.

## Required edge cases

- empty / nil input
- `lo > hi`
- single element in range / out of range
- all elements in range
- no elements in range
- negative values and mixed signs
- `lo == hi` (exact match only)

## Assessed by

Implementation, Edge cases.

## Pass criteria

All instructor-provided tests pass.

## If it fails

Trace the count by hand on the smallest failing input, then fix and retry.

## Run

```
go test ./phase-1/exit/a-count-in-range/ -v
```
