# 1.7 Debugging against failing tests — Lab `a-faulty`

## Task

Take a deliberately faulty implementation of a simple algorithm. The instructor supplies both the faulty code and a failing test suite (including regression and boundary cases).

- [ ] Run the provided tests and reproduce the failure.
- [ ] Identify the smallest useful failing input from the test output.
- [ ] Fix the underlying defect in the implementation.
- [ ] Confirm the provided regression and boundary tests pass.
- [ ] Explain why the original implementation failed and why those tests catch the regression.

## Contract

`RunningSum(nums []int) []int` returns a new slice of the same length where `out[i]` is the sum of `nums[0]` through `nums[i]` (inclusive). A nil or empty input returns a result with `len` 0. The function must not modify `nums`.

## Required edge cases

- empty / nil input
- single element
- two elements
- mixed positive and negative values
- zeros
- input left unmodified (mutation check)

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

Explain why the original implementation failed and why the provided tests would catch the regression.

## If it fails

Write your hypothesis before each edit, change one thing at a time, and retry on a second faulty implementation.

## Run

```
go test ./phase-1/1-7-debugging/a-faulty/ -v
```
