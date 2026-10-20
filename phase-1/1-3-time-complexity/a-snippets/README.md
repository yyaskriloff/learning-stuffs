# 1.3 Time complexity — Lab `a-snippets`

## Task

Analyze short snippets containing:

- [ ] A single loop.
- [ ] Nested loops.
- [ ] Two consecutive loops.
- [ ] A loop that halves its input each iteration.
- [ ] A loop whose number of iterations depends on a nested condition.

## Contract

None. This lab is analysis only — do not change the snippet code to “fix” complexity.

## Required edge cases

Not applicable (no implementation contract). When counting iterations for remediation, use `n = 1, 2, 4, 8, 16`.

## Assessed by

Complexity.

## Pass criteria

At least 80% correct, each with a justification rather than only a Big-O label. Explain each snippet orally to the instructor (dominant operation, variables such as `n`, and the bound).

## If it fails

For each missed snippet, count iterations exactly for `n = 1, 2, 4, 8, and 16`, then fit the pattern. Retry on a fresh set of snippets.

## Snippets

Work from `snippets.go`. For each function, identify the **dominant operation** (what runs in the innermost body that matters), how many times it runs in terms of `n`, and the **Big-O time** bound. Justify from the loop structure.

| ID | Function | Pattern |
| --- | --- | --- |
| A | `SnippetA` | single loop |
| B | `SnippetB` | nested loops |
| C | `SnippetC` | two consecutive loops |
| D | `SnippetD` | loop that halves each iteration |
| E | `SnippetE` | iterations depend on a nested condition |

## Run

There are no correctness tests for this lab. To compile the package:

```
go test ./phase-1/1-3-time-complexity/a-snippets/ -v
```

Optional: call a snippet with your own `op` that increments a counter to measure iterations for small `n`.
