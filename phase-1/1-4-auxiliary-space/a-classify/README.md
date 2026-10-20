# 1.4 Auxiliary space — Lab `a-classify`

## Task

Classify the memory usage of algorithms that:

- [ ] Use a fixed number of variables.
- [ ] Allocate a slice proportional to input size.
- [ ] Build a map from the input.
- [ ] Use recursion with input-dependent depth.

## Contract

None. Analysis only — do not change the example functions to alter their space use.

## Required edge cases

Not applicable. When remediating, list every allocation and every stack frame alive at the peak for each miss.

## Assessed by

Complexity, Reasoning.

## Pass criteria

Correctly classify at least 4 of 5 examples (A–E in `classify.go`). For each, state auxiliary space in Big-O (in terms of `n` = input length where relevant) and justify from allocations / stack. Also explain orally why auxiliary space differs from the total memory occupied by the input and output.

## If it fails

For each miss, list every allocation and every stack frame alive at the peak, then retry on new examples.

## Examples

| ID | Function | What to classify |
| --- | --- | --- |
| A | `ExampleA` | auxiliary space vs input size `n` |
| B | `ExampleB` | same |
| C | `ExampleC` | same |
| D | `ExampleD` | same |
| E | `ExampleE` | same |

Match each example to one of the patterns in the task (or say if it mixes patterns), give **O(…)** auxiliary space, and point at the code that causes it.

## Run

There are no correctness tests for this lab. To compile the package:

```
go test ./phase-1/1-4-auxiliary-space/a-classify/ -v
```
