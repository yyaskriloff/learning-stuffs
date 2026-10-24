# Phase 1 exit — Lab `b-any-pair-sum`

## Task

- [ ] Implement `AnyPairSum(nums []int, target int) bool`.

## Contract

`AnyPairSum` reports whether there exist two **distinct indexes** `i` and `j` such that `nums[i] + nums[j] == target`.

- A nil, empty, or single-element input returns `false`.
- The same value may appear twice in the slice; that counts as a valid pair when indexes differ.
- Do not modify `nums`.
- A brute-force (nested loop) solution is fine.

## Required edge cases

- empty / nil input
- single element
- two elements that sum to target / do not
- pair that reuses the same value at different indexes
- negative values
- zeros
- no solution among many elements

## Assessed by

Implementation, Edge cases, Algorithm correctness.

## Pass criteria

All instructor-provided tests pass, including the randomized check against a brute-force reference.

## If it fails

Shrink the failing input to the smallest case that still fails, then fix and retry.

## Run

```
go test ./phase-1/exit/b-any-pair-sum/ -v
```
