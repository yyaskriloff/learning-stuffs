# Phase 1 exit — Lab `c-partition`

## Task

- [ ] Implement `Partition(nums []int, pivot int) (less, ge []int)`.

## Contract

`Partition` returns two new slices:

- `less`: every element of `nums` that is **strictly less than** `pivot`, in the same relative order as in `nums`.
- `ge`: every element of `nums` that is **greater than or equal to** `pivot`, in the same relative order as in `nums`.

Together, `less` and `ge` contain every element of `nums` exactly once (as a multiset).

- A nil or empty `nums` returns two empty slices (`len` 0 each).
- Do **not** modify `nums`.
- The returned slices must not alias `nums` (mutating a result must not change the input).

## Required edge cases

- empty / nil input
- all less than pivot
- all greater than or equal to pivot
- duplicates equal to pivot
- negatives and zero
- single element

## Assessed by

Implementation, Edge cases.

## Pass criteria

All instructor-provided tests pass, including mutation and aliasing checks.

## If it fails

Redo 1.6 if the mutation/aliasing case fails, then retry.

## Run

```
go test ./phase-1/exit/c-partition/ -v
```
