# 1.5 Go functions, errors, and contracts — Lab `a-divide`

## Task

Implement integer division that reports division by zero through an error.

- [ ] Define the function signature.
- [ ] Implement the normal case.
- [ ] Return an appropriate error for a zero divisor.
- [ ] Make the provided tests pass for both success and failure outcomes.
- [ ] Explain why a function contract should make invalid inputs and failure behavior explicit.

## Contract

- `Divide(a, b int) (int, error)` performs Go integer division (truncates toward zero).
- Division by zero returns the sentinel `ErrDivideByZero` (callers check with `errors.Is`).
- For `math.MinInt / -1`, decide the contract, document it on the function, then ask the instructor to complete the skipped test case.

## Required edge cases

- zero divisor
- zero dividend
- negative operands
- truncation direction for negative results
- `math.MinInt` divided by `-1` (open — you decide and document)

## Assessed by

Implementation, Edge cases, Reasoning.

## Pass criteria

Callers can tell success from failure without parsing an error message. Use a sentinel error checked with `errors.Is`. Explain orally why the contract should make invalid inputs and failure behavior explicit.

## If it fails

Study how `strconv.Atoi` and `errors.Is` report failure, then redesign the signature.

## Run

```
go test ./phase-1/1-5-errors/a-divide/ -v
```
