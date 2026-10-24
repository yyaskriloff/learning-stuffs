# 1.9 Brute force versus improvement — Lab `a-duplicates`

## Task

Implement duplicate detection twice: once with pairwise comparisons and once with a map.

- [ ] Make both implementations pass the same provided test inputs.
- [ ] Pass the provided randomized comparison between the two implementations.
- [ ] Run the provided benchmarks at several input sizes.
- [ ] Explain the time and space trade-offs.
- [ ] Identify when the simpler solution may be preferable.

## Contract

`HasDuplicatePairwise(nums []int) bool` and `HasDuplicateMap(nums []int) bool` both report whether any value appears more than once in `nums`. An empty or single-element input has no duplicates (`false`). Neither function may modify `nums`.

- Pairwise: compare elements using nested loops (no map).
- Map: use a `map` to track values seen so far.

## Required edge cases

- empty / nil input
- single element
- no duplicates
- adjacent duplicates
- non-adjacent duplicates
- all equal
- negative values

## Assessed by

Implementation, Algorithm correctness, Complexity, Reasoning.

## Pass criteria

Both implementations agree across the provided fixed and random inputs, and you can explain why their complexity differs (oral check in chat after the tests pass).

## If it fails

If the implementations disagree, shrink the random failing input by hand to the smallest one that still fails, and ask the instructor to add it as a fixed test case.

## Run

```
go test ./phase-1/1-9-brute-force/a-duplicates/ -v
go test ./phase-1/1-9-brute-force/a-duplicates/ -bench=. -benchmem
```
