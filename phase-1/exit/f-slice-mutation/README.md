# Phase 1 exit — Lab `f-slice-mutation`

## Task

- [ ] Implement `RotateLeft(nums []int)` — rotate the slice **in place** one position to the left.
- [ ] Implement `DropWhile(nums []int, drop func(int) bool) []int` — return a **new** slice that drops a leading run of elements for which `drop` is true.
- [ ] Before running the demonstration tests in `demo_test.go`, predict in chat (or local scratch):
  - whether `RotateLeft` changes the backing array address for a typical non-empty slice;
  - whether mutating the result of `DropWhile` can change the input.

Then run the tests and reconcile any wrong prediction.

## Contract

**`RotateLeft`**

- Moves every element one index left; the former first element becomes the last.
- Empty and single-element slices are no-ops.
- Must mutate `nums` in place (same backing array for non-empty inputs with `len >= 2`).

**`DropWhile`**

- Drops the longest prefix of `nums` where every element satisfies `drop`.
- Returns a new slice with the remaining elements in order (may be empty).
- Must not modify `nums`.
- The result must not alias `nums`.

## Required edge cases

- empty / nil input
- single element
- rotate of two and many elements
- DropWhile drops nothing / drops all / drops a proper prefix
- mutation and aliasing checks

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

All instructor-provided tests pass. You predicted storage/aliasing behavior before confirming with the demo tests.

## If it fails

Draw the slice header (pointer, length, capacity) for every variable in the failed example, redo 1.6 if needed, then retry.

## Run

```
go test ./phase-1/exit/f-slice-mutation/ -v
```
