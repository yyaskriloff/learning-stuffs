# 1.8 Iteration with contracts — Lab `a-min-max-sorted`

## Task

Implement:

- [ ] `MinMax(nums []int) (min, max int, ok bool)`
- [ ] `IsSortedNonDecreasing(nums []int) bool`
- [ ] `IndexOfMax(nums []int) (int, bool)`, returning the first occurrence when the maximum repeats.

## Contract

Define what `MinMax` and `IndexOfMax` return for an empty slice. Document your decisions in doc comments, then ask the instructor to complete the skipped contract tests.

## Required edge cases

- empty input
- duplicates
- negative values
- already-sorted input
- reverse-sorted input

## Assessed by

Implementation, Edge cases, Complexity.

## Pass criteria

The provided tests cover every required edge case, your implementation passes them, and complexity is documented for each function (oral check in chat after the tests pass).

## If it fails

Write the contract as a doc comment before the code, then fix the implementation to match that contract (and any completed contract tests).

## Run

```
go test ./phase-1/1-8-contracts/a-min-max-sorted/ -v
```
