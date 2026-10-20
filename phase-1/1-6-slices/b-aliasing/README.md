# 1.6 Slices, mutation, and capacity — Lab `b-aliasing`

## Task

Before running the tests, predict orally (in chat or out loud):

- [ ] How `len` and `cap` differ in the given examples.
- [ ] How `append` may reuse or replace a backing array.
- [ ] How slice aliasing can make one function affect another variable.

Then run the provided tests and reconcile any wrong predictions.

## Contract

Nothing to implement. The tests demonstrate real Go slice behavior. Predict first, then confirm.

## Required edge cases

Covered by the demonstration scenarios (full slice, reslice, append with spare capacity, append that grows, shared headers).

## Assessed by

Reasoning.

## Pass criteria

You predict which slices share underlying storage in the scenarios below. Predict first orally, then confirm by running the tests. Reconcile mismatches by drawing ptr/len/cap for each variable.

## If it fails

Draw the slice header (pointer, length, capacity) for every variable in the failed example, then retry with new examples.

## Scenarios to predict (before `go test`)

Open `aliasing_test.go` and for each test name, answer before running:

1. **`TestLenAndCap`** — For `make([]int, 2, 5)` and for `s[1:2]`, what are `len` and `cap`?
2. **`TestAppendReusesBacking`** — After `b := a[:2]` and `b = append(b, 99)`, does `a` change? Do `a` and `b` share storage?
3. **`TestAppendReplacesBacking`** — After appending past capacity, does the original slice’s array change? Do the headers still share a pointer?
4. **`TestAliasingAcrossVariables`** — After `y := x` and `y[1] = 9`, what is `x`?
5. **`TestAliasingThroughFunction`** — If a function writes `nums[0] = 7`, does the caller’s slice change?

## Run

```
go test ./phase-1/1-6-slices/b-aliasing/ -v
```
