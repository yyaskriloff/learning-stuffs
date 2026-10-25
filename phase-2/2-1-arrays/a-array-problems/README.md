# 2.1 Arrays and slices — Lab `a-array-problems`

## Task

- [ ] Find the second-largest distinct number.
- [ ] Move zeroes to the end while preserving the order of nonzero values.
- [ ] Merge two sorted slices.
- [ ] Remove duplicates from a sorted slice in place, returning the valid logical length.

## Contract

Decide what second-largest returns when fewer than two distinct values exist. Document your decision in a doc comment, then ask the instructor to complete the skipped contract tests.

## Required edge cases

- empty input
- all duplicates
- all zeroes
- negative numbers
- unequal slice lengths

## Assessed by

Implementation, Edge cases, Complexity.

## Pass criteria

All tests pass, and `notes.md` says whether each implementation modifies its input and justifies its complexity.

## If it fails

Redo 1.6, then retry after listing the failing problem's edge cases in `notes.md` before coding.

## Run

```
go test ./phase-2/2-1-arrays/a-array-problems/ -v
```
