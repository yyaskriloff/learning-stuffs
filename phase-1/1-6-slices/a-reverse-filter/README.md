# 1.6 Slices, mutation, and capacity — Lab `a-reverse-filter`

## Task

Implement two functions:

- [ ] Reverse a slice in place.
- [ ] Filter a slice into a new result without changing the input.

## Contract

- `Reverse(nums []int)` reverses `nums` in place (same underlying array). A nil or empty slice is a no-op.
- `Filter(nums []int, keep func(int) bool) []int` returns a new slice of elements for which `keep` is true, in the same order. It must not modify `nums`. Filtering a nil or empty input returns an empty non-nil slice or nil — either is acceptable if `len` is 0; the tests only require `len(got) == 0` and that the input is unchanged.

## Required edge cases

- empty slices
- single elements
- duplicates
- mutation checks that assert the filter input is unchanged

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

All instructor-provided tests pass. Be ready to explain orally how in-place reverse differs from filter’s new slice (mutation / aliasing).

## If it fails

Draw the slice header (pointer, length, capacity) for the failing case, then retry.

## Run

```
go test ./phase-1/1-6-slices/a-reverse-filter/ -v
```
