# 1.1 Set up Go and the test workflow — Lab `a-add`

## Task

Create the repository and initialize a Go module. Implement `Add(a, b int) int` so the instructor-provided table-driven tests pass.

- [x] Create the module and make `go test ./...` run successfully.
- [x] Implement `Add` so the provided tests for positive, negative, and zero values pass.
- [x] Intentionally introduce a bug in the implementation, confirm a provided test fails, and diagnose why.
- [x] Fix the bug and confirm the tests pass again.

## Contract

None beyond the signatures.

## Required edge cases

- zero plus zero
- mixed signs
- both negative

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

The module builds, the regression against the provided tests is demonstrated, and you can explain the difference between a test failure and a compilation failure.

## If it fails

Read the `go help test` and `go help testflag` summaries, rerun the test with `-v` and `-run`, then redo the bug-introduction step with a different bug.

## Run

```
go test ./phase-1/1-1-setup/a-add/ -v
```
