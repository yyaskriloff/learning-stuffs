# Syllabus: DSA + Go Interview Training

This is the executable curriculum for [README.md](README.md). Work through it top to bottom. Every checkbox is a task with a result that can be checked. A box is checked only when its lesson's pass criteria are met. Doing the work is not enough.

---

## How to use this syllabus

### Structure

- **Phase:** a stage of the program with a goal and an exit assessment.
- **Lesson:** one topic inside a phase, such as `2.6 — Linked-list patterns`. Each lesson has its own directory.
- **Lab:** one self-contained piece of work inside a lesson. Each lab has its own directory.

Every lesson lists the following:

| Field | Meaning |
| --- | --- |
| Lab | The task and its deliverable |
| Checkboxes | The individual verifiable steps |
| Contract | Behavior you must decide and document before coding, when the problem leaves it open |
| Required edge cases | Boundary conditions the instructor's tests must cover |
| Assessed by | Which assessment types apply (see below) |
| Pass criteria | The definition of done and the pass/fail rule |
| If it fails | The remediation path before you retry |

### Directory layout

```
phase-1/
  1-1-setup/
    a-add/
  1-2-loops/
    a-sum-count-even/
  ...
  exit/
phase-2/
  ...
mistakes.md
```

Each lab directory is self-contained:

- `README.md`: the lab's task, contract, edge cases, pass criteria, and remediation, copied from this file.
- `<name>.go`: your implementation. Use one Go package per lab.
- `<name>_test.go`: provided by the instructor (the agent). Table-driven tests, edge-case tests, and property or randomized tests where the lab requires them. Do not rewrite these to make a failing solution look correct.
- `notes.md`: written answers for complexity and reasoning items.
- `bench_test.go`: only when the lab asks for a benchmark; the instructor provides the skeleton or full benchmark.

Labs never import other labs. If a later lab needs an earlier helper, such as a tree builder, copy it in. That keeps each lab independently runnable and keeps the old work unchanged.

The repository has one Go module at the root, created in lesson 1.1. `go test ./...` runs everything, and `go test ./phase-3/...` runs one phase.

### Who writes which tests

- **The instructor (agent)** writes the tests for every lab when the lab is scaffolded: normal cases, required edge cases, and property or randomized checks where the syllabus asks for them.
- **You** write the implementation (and `notes.md` answers) that satisfy those tests. Your job is the solution, not the test suite.
- When a lab leaves a contract open (for example empty-input return values), document your decision in a doc comment or `notes.md`. The instructor then completes any skipped contract tests to match that decision before the lab can pass.
- **Exit assessments** use unseen problems and tests that the instructor provides. Those go in the phase's `exit/` directory.
- A lab whose code fails its tests is not passed. Do not edit the test files to hide failures.

### Assessment types

| Assessment type | Verification |
| --- | --- |
| Implementation | Automated tests verify outputs and behavior. |
| Edge cases | Dedicated tests verify boundary conditions. |
| Algorithm correctness | Property-based or randomized tests where appropriate, usually comparing against a brute-force reference. |
| Complexity | Written analysis in `notes.md`, benchmarks where useful, graded with the complexity rubric. |
| Reasoning | A written or oral answer graded with the reasoning rubric. |
| Interview simulation | A timed task scored with the interview rubric. |

A passing test suite does not satisfy a Reasoning or Complexity item. Those are graded separately.

### Shared rubrics

**Complexity rubric.** Score each answer from 0 to 4, one point per criterion:

1. The time bound is correct.
2. It names the dominant operation and the variables involved (n, k, V, E, and so on).
3. The justification points to specific parts of the code, not only a label.
4. Auxiliary space is stated separately from the input and output.

An answer passes at 3/4 if the bound is correct. A wrong bound fails regardless of score.

**Reasoning rubric.** An explanation passes when it does all of the following:

1. States what the algorithm computes.
2. Gives the invariant or key insight that makes it correct.
3. Explains why it terminates (for loops and recursion).
4. Shows how the required edge cases are handled.
5. Does not rely on "the tests pass" as the argument.

**Interview rubric.** Score each category 0, 1, or 2, for a maximum of 14:

| Category | 2 means |
| --- | --- |
| Understanding | Restated the problem and clarified constraints before coding |
| Approach | Gave a brute-force solution, then a justified improvement |
| Correctness | Solution passes the provided tests |
| Code quality | Readable, idiomatic Go |
| Testing | Identified and ran edge cases unprompted |
| Complexity | Correct time and space, justified |
| Communication | Reasoning was followable throughout |

A mock passes at 11/14 or higher, with Correctness above 0.

### Remediation rules

- On a failure, record the cause in `mistakes.md` with its category: algorithm, Go, debugging, or reasoning.
- Do the remediation listed in the lesson, then retry with a **new** input set or variation. Do not re-run the same failing attempt.
- After two failed retries, drop back to the earlier lesson named in the remediation path.

### Retention

Every phase exit includes a retest of a lab from an earlier phase, at least one week after it was first passed. Re-solve it from an empty directory. A topic counts as **Retained** only after it passes that retest.

---

## Phase 1 — Foundations

**Goal:** become comfortable writing Go solutions that satisfy provided tests and reasoning about algorithmic cost.

### 1.1 — Set up Go and the test workflow

Directory: `phase-1/1-1-setup/`

**Lab** `a-add`**:** Create the repository and initialize a Go module. Implement `Add(a, b int) int` so the instructor-provided table-driven tests pass.

- [x] Create the module and make `go test ./...` run successfully.
- [x] Implement `Add` so the provided tests for positive, negative, and zero values pass.
- [x] Intentionally introduce a bug in the implementation, confirm a provided test fails, and diagnose why.
- [x] Fix the bug and confirm the tests pass again.

**Required edge cases:** zero plus zero, mixed signs, both negative.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** the module builds, the regression against the provided tests is demonstrated, and you can explain the difference between a test failure and a compilation failure.

**If it fails:** read the `go help test` and `go help testflag` summaries, rerun the test with `-v` and `-run`, then redo the bug-introduction step with a different bug.

### 1.2 — Loops and basic state

Directory: `phase-1/1-2-loops/`

**Lab** `a-sum-count-even`**:** Implement:

- [x] `Sum(nums []int) int`
- [x] `CountEven(nums []int) int`
- [x] Explain what changes during each loop iteration and why the final result is correct.

**Required edge cases:** empty slice, one element, all even, all odd, mixed values, negative values.

**Assessed by:** Implementation, Edge cases, Complexity, Reasoning.

**Pass criteria:** all tests pass, and `notes.md` explains why both algorithms take O(n) time and O(1) auxiliary space, scoring 3/4 or better on the complexity rubric.

**If it fails:** trace both loops by hand on `[3, -2, 0]` and write down the state after each iteration, then retry the explanation.

### 1.3 — Time complexity

Directory: `phase-1/1-3-time-complexity/`

**Lab** `a-snippets`**:** Analyze short snippets containing:

- [x] A single loop.
- [x] Nested loops.
- [x] Two consecutive loops.
- [x] A loop that halves its input each iteration.
- [x] A loop whose number of iterations depends on a nested condition.

**Deliverable:** a written answer in `notes.md` for each snippet that identifies the dominant operation and the time complexity.

**Assessed by:** Complexity.

**Pass criteria:** at least 80% correct, each with a justification rather than only a Big-O label.

**If it fails:** for each missed snippet, count iterations exactly for n = 1, 2, 4, 8, and 16, then fit the pattern. Retry on a fresh set of snippets.

### 1.4 — Auxiliary space

Directory: `phase-1/1-4-auxiliary-space/`

**Lab** `a-classify`**:** Classify the memory usage of algorithms that:

- [x] Use a fixed number of variables.
- [x] Allocate a slice proportional to input size.
- [x] Build a map from the input.
- [x] Use recursion with input-dependent depth.

**Deliverable:** explain in `notes.md` why auxiliary space differs from the total memory occupied by the input and output.

**Assessed by:** Complexity, Reasoning.

**Pass criteria:** correctly classify at least 4 of 5 unseen examples.

**If it fails:** for each miss, list every allocation and every stack frame alive at the peak, then retry on new examples.

### 1.5 — Go functions, errors, and contracts

Directory: `phase-1/1-5-errors/`

**Lab** `a-divide`**:** Implement integer division that reports division by zero through an error.

- [x] Define the function signature.
- [x] Implement the normal case.
- [x] Return an appropriate error for a zero divisor.
- [x] Make the provided tests pass for both success and failure outcomes.
- [x] Explain why a function contract should make invalid inputs and failure behavior explicit.

**Required edge cases:** zero divisor, zero dividend, negative operands, the truncation direction for negative results, and `math.MinInt` divided by `-1`. For the last case, decide the contract and document it.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** callers can tell success from failure without parsing an error message. Use a sentinel error checked with `errors.Is`.

**If it fails:** study how `strconv.Atoi` and `errors.Is` report failure, then redesign the signature.

### 1.6 — Slices, mutation, and capacity

Directory: `phase-1/1-6-slices/`

**Lab** `a-reverse-filter`**:** Implement two functions:

- [x] Reverse a slice in place.
- [x] Filter a slice into a new result without changing the input.

**Lab** `b-aliasing`**:** The instructor provides demonstration tests. Before running them, predict in `notes.md`:

- [x] How `len` and `cap` differ in the given examples.
- [x] How `append` may reuse or replace a backing array.
- [x] How slice aliasing can make one function affect another variable.

Then run the provided tests and reconcile any wrong predictions.

**Required edge cases:** empty slices, single elements, duplicates, and mutation checks that assert the filter input is unchanged.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** you predict which slices share underlying storage in several small examples. Predict first in writing, then confirm by running the provided tests.

**If it fails:** draw the slice header (pointer, length, capacity) for every variable in the failed example, then retry with new examples.

### 1.7 — Debugging against failing tests

Directory: `phase-1/1-7-debugging/`

**Lab** `a-faulty`**:** Take a deliberately faulty implementation of a simple algorithm. The instructor supplies both the faulty code and a failing test suite (including regression and boundary cases).

- [x] Run the provided tests and reproduce the failure.
- [x] Identify the smallest useful failing input from the test output.
- [x] Fix the underlying defect in the implementation.
- [x] Confirm the provided regression and boundary tests pass.
- [ ] Explain why the original implementation failed and why those tests catch the regression.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** explain why the original implementation failed and why the provided tests would catch the regression.

**If it fails:** write your hypothesis before each edit, change one thing at a time, and retry on a second faulty implementation.

### 1.8 — Iteration with contracts

Directory: `phase-1/1-8-contracts/`

**Lab** `a-min-max-sorted`**:** Implement:

- [ ] `MinMax(nums []int) (min, max int, ok bool)`
- [ ] `IsSortedNonDecreasing(nums []int) bool`
- [ ] `IndexOfMax(nums []int) (int, bool)`, returning the first occurrence when the maximum repeats.

**Contract:** define what `MinMax` and `IndexOfMax` return for an empty slice.

**Required edge cases:** empty input, duplicates, negative values, already-sorted input, reverse-sorted input.

**Assessed by:** Implementation, Edge cases, Complexity.

**Pass criteria:** the provided tests cover every required edge case, your implementation passes them, and complexity is documented for each function.

**If it fails:** write the contract as a doc comment before the code, then fix the implementation to match that contract (and any completed contract tests).

### 1.9 — Brute force versus improvement

Directory: `phase-1/1-9-brute-force/`

**Lab** `a-duplicates`**:** Implement duplicate detection twice: once with pairwise comparisons and once with a map.

- [ ] Make both implementations pass the same provided test inputs.
- [ ] Pass the provided randomized comparison between the two implementations.
- [ ] Run the provided benchmarks at several input sizes.
- [ ] Explain the time and space trade-offs.
- [ ] Identify when the simpler solution may be preferable.

**Assessed by:** Implementation, Algorithm correctness, Complexity, Reasoning.

**Pass criteria:** both implementations agree across the provided fixed and random inputs, and you can explain why their complexity differs.

**If it fails:** if the implementations disagree, shrink the random failing input by hand to the smallest one that still fails, and ask the instructor to add it as a fixed test case.

### Phase 1 exit assessment

Directory: `phase-1/exit/`

- [ ] Solve three unseen beginner problems, each passing the instructor's tests.
- [ ] Score at least 8/10 on a complexity quiz.
- [ ] Explain a bug and how the provided regression test prevents it from returning.
- [ ] Pass a fresh slice-mutation exercise.

**If it fails:** repeat the lesson that matches the missed item (1.3 or 1.4 for complexity, 1.6 for slices, 1.7 for debugging), then retake that item with a new problem.

---

## Phase 2 — Core data structures

**Goal:** understand how the way data is represented affects operations, complexity, and correctness.

### 2.1 — Arrays and slices

Directory: `phase-2/2-1-arrays/`

**Lab** `a-array-problems`**:**

- [ ] Find the second-largest distinct number.
- [ ] Move zeroes to the end while preserving the order of nonzero values.
- [ ] Merge two sorted slices.
- [ ] Remove duplicates from a sorted slice in place, returning the valid logical length.

**Contract:** decide what second-largest returns when fewer than two distinct values exist.

**Required edge cases:** empty input, all duplicates, all zeroes, negative numbers, unequal slice lengths.

**Assessed by:** Implementation, Edge cases, Complexity.

**Pass criteria:** all tests pass, and `notes.md` says whether each implementation modifies its input and justifies its complexity.

**If it fails:** redo 1.6, then retry after listing the failing problem's edge cases in `notes.md` before coding.

### 2.2 — Strings, bytes, and runes

Directory: `phase-2/2-2-strings/`

**Lab** `a-runes`**:**

- [ ] Reverse a string by Unicode rune.
- [ ] Count rune frequencies.
- [ ] Check whether a string is a palindrome under a clearly defined normalization rule.
- [ ] Find the first unique rune.

**Contract:** write the palindrome normalization rule, such as case folding and which characters are ignored.

**Required edge cases:** empty string, single rune, ASCII-only input, multibyte Unicode characters.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** tests include ASCII and multibyte Unicode characters, and you show with an example why Go string byte indexing and rune iteration are not interchangeable.

**If it fails:** print `len(s)`, `[]byte(s)`, and `[]rune(s)` for a string such as `"héllo"`, explain each result, and retry.

### 2.3 — Hash maps

Directory: `phase-2/2-3-maps/`

**Lab** `a-frequency`**:**

- [ ] Implement a frequency map.
- [ ] Detect duplicates.
- [ ] Find the first unique element.
- [ ] Check whether two strings are anagrams under a stated normalization contract.

**Required edge cases:** empty input, all identical, no duplicates, strings of different lengths.

**Assessed by:** Implementation, Edge cases, Complexity, Reasoning.

**Pass criteria:** all tests pass. You explain expected average-case map performance, and why it does not guarantee constant worst-case behavior for every operation or implementation.

**If it fails:** explain what a hash collision is and what happens to lookup time when many keys collide, then retry the explanation.

### 2.4 — Sets and membership

Directory: `phase-2/2-4-sets/`

Go has no built-in set type, so use a map-based representation.

**Lab** `a-set-ops`**:**

- [ ] Return unique values.
- [ ] Find values shared by two collections.
- [ ] Check whether one collection is a subset of another.
- [ ] Decide whether result order matters and implement the contract accordingly.

**Required edge cases:** duplicates, empty collections, identical collections, disjoint collections.

**Assessed by:** Implementation, Edge cases, Complexity.

**Pass criteria:** tests cover duplicates and empty collections. Complexity and ordering guarantees are documented. Tests do not depend on map iteration order.

**If it fails:** if a test is flaky because of map iteration order, fix the implementation to return a deterministic result (or document a set-equality contract for the instructor to encode in the tests).

### 2.5 — Linked-list fundamentals

Directory: `phase-2/2-5-linked-lists/`

**Lab** `a-list-basics`**:**

- [ ] Define a singly linked-list node.
- [ ] Build a list from a slice.
- [ ] Traverse a list and convert it back to a slice.
- [ ] Insert at the head and tail.
- [ ] Delete the first occurrence of a value.

**Required edge cases:** empty list, one node, deleting the head, deleting the tail, missing value.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** all tests pass, and `notes.md` includes a drawing of the pointer changes for one insertion and one deletion.

**If it fails:** draw the nodes and arrows before and after each operation on paper, then rewrite the failing function from the drawing.

### 2.6 — Linked-list patterns

Directory: `phase-2/2-6-list-patterns/`

**Lab** `a-reverse-middle`**:**

- [ ] Reverse a linked list.
- [ ] Find its middle node.

**Lab** `b-cycle-merge`**:**

- [ ] Detect a cycle using Floyd's algorithm.
- [ ] Merge two sorted linked lists.

**Contract:** decide which middle node is returned when the list has an even number of nodes.

**Required edge cases:** empty list, one node, two nodes, cycle back to the head, cycle back to a later node, one empty input to merge.

**Assessed by:** Implementation, Edge cases, Reasoning.

**Pass criteria:** all tests pass, and you explain the pointer invariant for reversal and why cycle detection works, meeting the reasoning rubric.

**If it fails:** redo 2.5. Then trace reversal on a three-node list, writing `prev`, `curr`, and `next` after each step.

### 2.7 — Stacks

Directory: `phase-2/2-7-stacks/`

**Lab** `a-stack`**:**

- [ ] Implement a stack API with push, pop, and peek operations.
- [ ] Define what happens when the stack is empty.

**Lab** `b-stack-problems`**:**

- [ ] Check balanced delimiters.
- [ ] Evaluate postfix arithmetic expressions.

**Required edge cases:** empty input, only opening delimiters, only closing delimiters, interleaved mismatches like `([)]`, a postfix expression with too few operands, division by zero in postfix.

**Assessed by:** Implementation, Edge cases, Complexity, Reasoning.

**Pass criteria:** all tests pass. You explain why a stack fits nested expressions naturally, and what makes its core operations efficient in your implementation.

**If it fails:** trace the stack contents by hand for `{[()]}` and `([)]`, then retry.

### 2.8 — Queues and deques

Directory: `phase-2/2-8-queues/`

**Lab** `a-queue`**:**

- [ ] Implement a FIFO queue.
- [ ] Avoid repeatedly shifting all remaining elements when dequeuing.

**Lab** `b-deque`**:**

- [ ] Implement a deque or a clearly specified subset of deque operations, using a circular buffer.
- [ ] Make the provided tests pass for empty behavior and repeated enqueue/dequeue cycles.

**Required edge cases:** dequeue from empty, wraparound in the circular buffer, growth when full, long alternating enqueue/dequeue sequences.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** a randomized test that compares your deque against a simple slice model passes. `notes.md` compares a simple slice-based queue, a head-index queue, and a circular buffer.

**If it fails:** draw the buffer, head, and tail indexes after each operation in the smallest failing sequence.

### 2.9 — Choosing a data structure

Directory: `phase-2/2-9-choosing/`

**Lab** `a-scenarios`**:** For at least ten scenarios supplied by the instructor, choose a suitable data structure and justify the choice.

- [ ] Cover lookup-heavy, insertion-heavy, ordered, sequential, and priority-based workloads.
- [ ] State the operations that dominate each workload.
- [ ] Identify the complexity trade-off.
- [ ] Implement one scenario with the structure you chose.

**Assessed by:** Reasoning, Complexity, Implementation.

**Pass criteria:** at least 8/10 choices are defensible, and the implementation meets its contract.

**If it fails:** for each miss, list the workload's operations and their costs in every candidate structure, then retry on new scenarios.

### Phase 2 exit assessment

Directory: `phase-2/exit/`

- [ ] Solve one unseen array problem.
- [ ] Solve one unseen map/set problem.
- [ ] Solve one unseen stack/queue problem.
- [ ] Solve one unseen linked-list problem.
- [ ] Score at least 80% on data-structure selection.
- [ ] Retention: re-solve one Phase 1 lab from an empty directory.

**If it fails:** repeat the lesson for the missed structure, then retake that item with a new problem.

---

## Phase 3 — Problem-solving patterns

**Goal:** find a useful approach to an unfamiliar problem without being told which pattern to use.

### 3.1 — Two pointers

Directory: `phase-3/3-1-two-pointers/`

**Lab** `a-opposite-ends`**:**

- [ ] Two Sum on sorted input, returning both indexes and an `ok` flag.
- [ ] Valid palindrome, ignoring non-alphanumeric ASCII characters and case.
- [ ] Container with most water.

**Lab** `b-same-direction`**:**

- [ ] Remove all occurrences of a value in place, returning the new logical length.
- [ ] 3Sum, returning every unique triplet that sums to zero with no duplicate triplets.

**Required edge cases:** empty input, exactly two elements, no solution, many duplicates, negative values.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** a randomized test that compares Two Sum and 3Sum against brute-force versions passes. You state the invariant that justifies each pointer move.

**If it fails:** write the brute-force version first and use it to find a failing input. Shrink that input, then trace the pointers on it.

### 3.2 — Sliding window

Directory: `phase-3/3-2-sliding-window/`

**Lab** `a-fixed-window`**:**

- [ ] Maximum sum of any subarray of length `k`.

**Lab** `b-variable-window`**:**

- [ ] Longest substring without repeating characters.
- [ ] Minimum-length subarray whose sum is at least a target.
- [ ] Longest substring with at most `k` distinct characters.

**Contract:** define behavior for `k <= 0` and `k > len`, and state whether input is bytes or runes. The minimum-length problem assumes positive integers. Document that assumption.

**Required edge cases:** empty input, `k` equal to length, all identical characters, no valid window.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity, Reasoning.

**Pass criteria:** randomized tests against O(n²) brute-force versions pass. You explain why each index enters and leaves the window at most once, and why the minimum-length solution breaks when negative numbers are allowed.

**If it fails:** print the window boundaries and state at each step for the failing input, and identify which update happened in the wrong order.

### 3.3 — Prefix sums

Directory: `phase-3/3-3-prefix-sums/`

**Lab** `a-range-sum`**:**

- [ ] A `RangeSum` type with a constructor and `Query(l, r)`.
- [ ] A 2D prefix sum with rectangle queries.

**Lab** `b-prefix-map`**:**

- [ ] Count subarrays whose sum equals `k`, with negative numbers allowed.
- [ ] Apply a list of range increments with a difference array and return the final slice.

**Contract:** decide whether ranges are inclusive or half-open and how invalid ranges are reported.

**Required edge cases:** single element, full range, negative values, `k = 0`, empty matrix.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** randomized tests against brute-force sums pass, and `notes.md` states the preprocessing and per-query costs.

**If it fails:** write out the prefix array for the failing input by hand and check the index arithmetic for off-by-one errors.

### 3.4 — Binary search

Directory: `phase-3/3-4-binary-search/`

**Lab** `a-search`**:**

- [ ] Classic search returning `(index, ok)`.
- [ ] Lower bound (first index `>= x`) and upper bound (first index `> x`).
- [ ] Search in a rotated sorted array of distinct values.

**Lab** `b-search-answer`**:**

- [ ] Find the minimum ship capacity that delivers all packages within `D` days.

**Required edge cases:** empty input, one element, target below all values, target above all values, many duplicates, a mid calculation that cannot overflow.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** randomized tests against a linear scan and `sort.SearchInts` pass. You state the loop invariant, why the loop terminates, and the monotonic predicate in lab `b`.

**If it fails:** write the invariant as a comment above the loop, then check every branch against it on the failing input.

### 3.5 — Sorting and intervals

Directory: `phase-3/3-5-sorting-intervals/`

**Lab** `a-sorts`**:**

- [ ] Implement insertion sort and merge sort yourself.
- [ ] Pass the provided stability tests for merge sort (structs with equal keys).
- [ ] Sort records by multiple keys with `slices.SortFunc` or `sort.Slice`.

**Lab** `b-intervals`**:**

- [ ] Merge overlapping intervals.
- [ ] Insert an interval into a sorted, non-overlapping list.
- [ ] Find the minimum number of meeting rooms needed.

**Contract:** decide whether intervals are closed or half-open, and whether touching intervals merge.

**Required edge cases:** empty input, single interval, fully nested intervals, touching intervals, unsorted input, identical intervals.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** randomized tests compare your sorts against `slices.Sort`. `notes.md` explains why merge sort is O(n log n) and the cost of sort-then-scan.

**If it fails:** draw the intervals on a number line for the failing input and compare them with the output.

### 3.6 — Frequency counting

Directory: `phase-3/3-6-frequency/`

**Lab** `a-counting`**:**

- [ ] Find the majority element with a map, then with Boyer-Moore voting.
- [ ] Return the top `k` most frequent elements by sorting the counts.
- [ ] Group anagrams.
- [ ] Check anagrams with a `[26]int` array and with a map, and compare the two.

**Contract:** define the tie-breaking order for top `k`, and what majority returns when no majority exists.

**Required edge cases:** empty input, all identical, every element unique, `k` equal to the number of distinct elements.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** a randomized test compares Boyer-Moore against the map version on inputs that have a majority. You explain why Boyer-Moore is correct and when an array beats a map for counting.

**If it fails:** trace the Boyer-Moore candidate and count by hand on the failing input.

### 3.7 — Pattern recognition

Directory: `phase-3/3-7-recognition/`

**Lab** `a-classify`**:** For 12 unseen problem statements from the instructor, before writing any code:

- [ ] Name the pattern you would use, or say that you would start with brute force.
- [ ] Give the brute-force complexity and the target complexity.
- [ ] Name the signal in the problem statement that pointed to the pattern.
- [ ] Implement two of the twelve, with tests.

**Assessed by:** Reasoning, Implementation.

**Pass criteria:** at least 10 of 12 classifications are defensible, and both implementations pass the instructor's tests.

**If it fails:** for each miss, write the signal you should have noticed, then retry on a new set.

### Phase 3 exit assessment

Directory: `phase-3/exit/`

- [ ] Solve four unseen problems where the pattern is not named, with 45 minutes per problem. At least three must pass the instructor's tests.
- [ ] For each problem, write the pattern and complexity before coding.
- [ ] Retention: re-solve one Phase 2 lab from an empty directory.

**If it fails:** repeat 3.7 with new problems, plus the lesson for each pattern you missed.

---

## Phase 4 — Trees, heaps, and graphs

**Goal:** reason about traversal, connectivity, ordering, and recursive state.

### 4.1 — Recursion

Directory: `phase-4/4-1-recursion/`

**Lab** `a-recursion-basics`**:**

- [ ] Sum a slice recursively.
- [ ] Compute `x^n` with fast exponentiation in O(log n) multiplications.
- [ ] Implement naive Fibonacci and count how many calls it makes.
- [ ] Reverse a linked list recursively.

**Deliverable:** in `notes.md`, draw the call tree for `fib(5)` and state the base case, recursive case, and maximum stack depth for each function.

**Contract:** define behavior for negative `n`.

**Required edge cases:** `n = 0`, `n = 1`, empty slice, empty list.

**Assessed by:** Implementation, Edge cases, Complexity, Reasoning.

**Pass criteria:** all tests pass, and you explain why recursion depth counts as auxiliary space.

**If it fails:** write out the stack frames by hand for the smallest failing input.

### 4.2 — Binary tree fundamentals

Directory: `phase-4/4-2-trees/`

**Lab** `a-traversals`**:**

- [ ] Define a `TreeNode`. The instructor provides a test helper that builds a tree from a level-order slice with `nil` markers.
- [ ] Preorder, inorder, and postorder traversal, both recursive and iterative with an explicit stack.
- [ ] Level-order traversal with a queue, returning one slice per level.

**Lab** `b-tree-properties`**:**

- [ ] Maximum depth.
- [ ] Invert a tree.
- [ ] Check whether two trees are the same.
- [ ] Check whether a tree is symmetric.

**Required edge cases:** nil tree, single node, left-skewed tree, right-skewed tree, complete tree.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** recursive and iterative traversals agree on randomly generated trees. You explain what each traversal order guarantees about when a node is visited.

**If it fails:** redo 4.1, then trace the iterative stack on a three-node tree.

### 4.3 — Tree recursion that returns state

Directory: `phase-4/4-3-tree-state/`

**Lab** `a-tree-state`**:**

- [ ] Path sum: does any root-to-leaf path add up to a target?
- [ ] Diameter of a binary tree.
- [ ] Check whether a tree is height-balanced in O(n).
- [ ] Lowest common ancestor in a binary tree.

**Required edge cases:** nil tree, single node, target reached at an internal node but not a leaf, diameter path that does not pass through the root, one node being an ancestor of the other.

**Assessed by:** Implementation, Edge cases, Complexity, Reasoning.

**Pass criteria:** all tests pass. For each function you explain what a recursive call returns versus what is tracked across calls, and why the naive balance check is O(n²).

**If it fails:** write each helper's return value as a one-sentence contract before rewriting it.

### 4.4 — Binary search trees

Directory: `phase-4/4-4-bst/`

**Lab** `a-bst-ops`**:**

- [ ] Insert and search.
- [ ] Delete, covering a leaf, one child, and two children.
- [ ] Validate a BST using bounds, not only parent comparisons.
- [ ] Find the `k`th smallest value.

**Contract:** decide how duplicates are handled.

**Required edge cases:** empty tree, deleting the root, deleting a node with two children, skewed tree, values at `math.MinInt` and `math.MaxInt`.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** randomized tests confirm that inorder traversal is sorted after random inserts and deletes, and that validation agrees. `notes.md` gives O(h) and explains balanced versus skewed height.

**If it fails:** build a validator that fails on the smallest broken tree, then trace the delete case that produced it.

### 4.5 — Heaps

Directory: `phase-4/4-5-heaps/`

**Lab** `a-binary-heap`**:**

- [ ] Implement a binary min-heap on a slice: push, pop, peek, and heapify.

**Lab** `b-container-heap`**:**

- [ ] Find the `k`th largest element using `container/heap`.
- [ ] Find the top `k` frequent elements with a heap, revisiting 3.6.
- [ ] Merge `k` sorted slices.

**Required edge cases:** pop from empty, duplicates, `k = 1`, `k = n`, empty slices among the `k` inputs.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** repeated pops from your heap match `slices.Sort` on random inputs. `notes.md` explains push/pop at O(log n), heapify at O(n), and heap versus sort for top `k`.

**If it fails:** draw the heap as a tree and as an array after each sift for the failing input.

### 4.6 — Graph representations

Directory: `phase-4/4-6-graphs/`

**Lab** `a-representations`**:**

- [ ] Build an adjacency list from an edge list, directed and undirected.
- [ ] Build an adjacency matrix from the same edges.
- [ ] Write a grid neighbor function with bounds checks, treating the grid as an implicit graph.

**Contract:** decide how self-loops and parallel edges are handled.

**Required edge cases:** empty graph, isolated vertex, self-loop, parallel edges, 1×1 grid, grid corners.

**Assessed by:** Implementation, Edge cases, Complexity.

**Pass criteria:** the list and matrix agree on every edge for random graphs. `notes.md` compares O(V + E) and O(V²) space and says when each wins.

**If it fails:** draw the failing graph and write out both representations by hand.

### 4.7 — Depth-first search

Directory: `phase-4/4-7-dfs/`

**Lab** `a-dfs`**:**

- [ ] Reachability, recursive and iterative.
- [ ] Count connected components.
- [ ] Number of islands in a grid.
- [ ] Detect a cycle in an undirected graph.

**Required edge cases:** disconnected graph, single vertex, all water, all land, 1×N grid, two vertices joined by one edge (not a cycle).

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity, Reasoning.

**Pass criteria:** recursive and iterative versions agree on random graphs. You explain O(V + E), and why marking visited at the wrong time can loop or double-count.

**If it fails:** print the visit order and visited set for the failing input.

### 4.8 — Breadth-first search

Directory: `phase-4/4-8-bfs/`

**Lab** `a-bfs`**:**

- [ ] Shortest path in an unweighted graph, returning the distance and the reconstructed path.
- [ ] Multi-source BFS: the time for rot to spread through a grid of oranges.
- [ ] Shortest path in a binary matrix with eight-directional moves.

**Required edge cases:** unreachable target, source equals target, blocked start, no fresh oranges, grid with no path.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** BFS distances match a brute-force path search on small random graphs. You explain why BFS gives shortest paths in unweighted graphs and why it fails with weights.

**If it fails:** print the queue contents for each layer on the failing input.

### 4.9 — Topological sort

Directory: `phase-4/4-9-topological-sort/`

**Lab** `a-topo`**:**

- [ ] Kahn's algorithm, returning an order and an `ok` flag that is false when there is a cycle.
- [ ] A DFS version using three colors.
- [ ] Course schedule: can every course be finished?

**Contract:** many valid orders can exist. Tests check that the output is a valid order, not that it matches one exact order.

**Required edge cases:** no edges, a cycle, a self-loop, a disconnected DAG.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** on random DAGs, every edge `u → v` puts `u` before `v`. When a back edge is added, both versions report a cycle. You explain what indegree and the three colors mean.

**If it fails:** trace indegrees and the queue (Kahn's) or the color changes (DFS) on the failing graph.

### 4.10 — Union-find

Directory: `phase-4/4-10-union-find/`

**Lab** `a-union-find`**:**

- [ ] Implement find and union with path compression and union by size.
- [ ] Count provinces from an adjacency matrix.
- [ ] Find the redundant connection in a graph that has one extra edge.

**Required edge cases:** union of an element with itself, repeated unions, a single element, every element already connected.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** component counts match DFS from 4.7 on random graphs. `notes.md` explains what each optimization prevents and the near-constant amortized cost.

**If it fails:** print the parent array after each union on the failing input.

### Phase 4 exit assessment

Directory: `phase-4/exit/`

- [ ] Solve one unseen tree problem.
- [ ] Solve one unseen graph traversal problem.
- [ ] Solve one unseen ordering or connectivity problem.
- [ ] Explain the recursive state of one solution, meeting the reasoning rubric.
- [ ] Retention: re-solve one Phase 3 lab from an empty directory.

**If it fails:** repeat the lesson for the missed area, then retake that item with a new problem.

---

## Phase 5 — Advanced patterns

**Goal:** break down problems that at first seem too complicated to solve directly.

Dynamic programming spans four lessons and a transfer check. Do not advance past 5.6 until it passes.

### 5.1 — Backtracking

Directory: `phase-5/5-1-backtracking/`

**Lab** `a-enumerate`**:**

- [ ] Subsets.
- [ ] Permutations, including input with duplicates and no repeated permutations in the output.
- [ ] Combination sum.

**Lab** `b-constrained-search`**:**

- [ ] N-Queens: count the solutions.
- [ ] Word search in a grid.

**Required edge cases:** empty input, duplicate values, no solution, target `0`, `n = 1`.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity, Reasoning.

**Pass criteria:** property tests pass (subsets count `2^n`, permutations count `n!` for distinct input, no duplicates in the output). `notes.md` includes a decision tree with the choose/explore/unchoose steps, the pruning rule, and complexity in terms of output size.

**If it fails:** print the partial candidate at each recursive call for the smallest failing input.

### 5.2 — Greedy algorithms

Directory: `phase-5/5-2-greedy/`

**Lab** `a-greedy`**:**

- [ ] Maximum number of non-overlapping intervals.
- [ ] Jump game: can you reach the end?
- [ ] Gas station: find the starting index.

**Lab** `b-counterexample`**:**

- [ ] Explain in `notes.md` why greedy coin change fails for coins `{1, 3, 4}` and target `6`, using the provided test against a brute-force version.

**Required edge cases:** empty input, a single element, all intervals overlapping, a zero jump at the start, no valid gas station.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** greedy results match brute force on small random inputs. `notes.md` contains an exchange argument for the interval problem that meets the reasoning rubric.

**If it fails:** find a small input where your greedy choice disagrees with brute force. Either fix the choice or show why it cannot be fixed.

### 5.3 — Dynamic programming I: memoization

Directory: `phase-5/5-3-dp-memo/`

**Lab** `a-memo`**:**

- [ ] Climbing stairs.
- [ ] Fibonacci with memoization, comparing call counts against the naive version from 4.1.
- [ ] House robber.
- [ ] Coin change (fewest coins), top-down.

**Deliverable:** before coding each problem, write in `notes.md` the state definition, recurrence, base cases, and where the answer lives.

**Required edge cases:** `n = 0`, `n = 1`, an unreachable target in coin change, a single house.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** memoized results match the naive recursive versions on small inputs, and every state definition was written before its code.

**If it fails:** write the naive recursion first, find the repeated subproblem in its call tree, and use that as the state.

### 5.4 — Dynamic programming II: tabulation and space

Directory: `phase-5/5-4-dp-table/`

**Lab** `a-bottom-up`**:**

- [ ] Convert house robber and coin change from 5.3 to bottom-up.
- [ ] Reduce house robber to O(1) space.
- [ ] Longest increasing subsequence in O(n²).
- [ ] Word break.

**Required edge cases:** empty input, a strictly decreasing sequence, a word dictionary containing the empty string, a word that cannot be segmented.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** bottom-up results match the top-down versions on random inputs. `notes.md` explains the fill order and why the space reduction is safe.

**If it fails:** fill the table by hand for the failing input and check the order of dependencies.

### 5.5 — Dynamic programming III: two dimensions

Directory: `phase-5/5-5-dp-2d/`

**Lab** `a-grid-strings`**:**

- [ ] Unique paths in a grid with obstacles.
- [ ] Longest common subsequence, returning the subsequence itself as well as its length.
- [ ] Edit distance.
- [ ] 0/1 knapsack.

**Required edge cases:** an obstacle at the start or end, empty strings, identical strings, zero capacity, an item heavier than the capacity.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Reasoning.

**Pass criteria:** results match brute-force recursion on small random inputs. The state definitions are written before the code.

**If it fails:** draw the table for the failing input and mark which cell each value depends on.

### 5.6 — Dynamic programming transfer check

Directory: `phase-5/5-6-dp-transfer/`

**Lab** `a-unseen-dp`**:** Three unseen DP problems from the instructor.

- [ ] For each, write the state, recurrence, and base cases before coding.
- [ ] Implement each with tests.

**Assessed by:** Reasoning, Implementation.

**Pass criteria:** at least 2 of 3 are solved without hints, and all three state definitions are correct.

**If it fails:** repeat 5.3 through 5.5 with new problems, then retake this lab with three different problems. Do not advance until it passes.

### 5.7 — Monotonic stacks

Directory: `phase-5/5-7-monotonic-stack/`

**Lab** `a-monotonic-stack`**:**

- [ ] Next greater element.
- [ ] Daily temperatures.
- [ ] Largest rectangle in a histogram.

**Required edge cases:** empty input, strictly increasing input, strictly decreasing input, all equal values.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** results match O(n²) brute-force versions on random inputs. `notes.md` argues the O(n) bound from each index being pushed and popped at most once.

**If it fails:** write the stack contents after each index for the failing input.

### 5.8 — Monotonic deques

Directory: `phase-5/5-8-monotonic-deque/`

**Lab** `a-window-max`**:**

- [ ] Sliding window maximum.

**Required edge cases:** `k = 1`, `k` equal to length, all equal values, decreasing input.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** results match brute force on random inputs, and the amortized O(n) bound is justified.

**If it fails:** redo 2.8, then trace the deque contents on the failing input.

### 5.9 — Tries

Directory: `phase-5/5-9-tries/`

**Lab** `a-trie`**:**

- [ ] Insert, search, and starts-with.
- [ ] Delete a word, pruning nodes that are no longer needed.
- [ ] Autocomplete: return every word with a given prefix in sorted order.

**Contract:** decide whether keys are bytes or runes, and how the empty string is handled.

**Required edge cases:** empty string, a prefix that is also a word, deleting a word that is a prefix of another word, a missing word.

**Assessed by:** Implementation, Edge cases, Algorithm correctness, Complexity.

**Pass criteria:** a randomized test compares the trie against a map-of-strings model. `notes.md` gives O(L) per operation and compares space with the map.

**If it fails:** draw the trie for the failing word set and mark the end-of-word flags.

### Phase 5 exit assessment

Directory: `phase-5/exit/`

- [ ] Solve one unseen backtracking problem.
- [ ] Solve one unseen greedy problem, with a written justification of the greedy choice.
- [ ] Solve two unseen DP problems, each with the state written before the code.
- [ ] Solve one unseen monotonic-structure or trie problem.
- [ ] Retention: re-solve one Phase 4 lab from an empty directory.

**If it fails:** repeat the lesson for the missed area, then retake that item with a new problem.

---

## Phase 6 — Interview simulation

**Goal:** consistently solve representative problems on your own under interview conditions.

Every task in this phase is scored with the interview rubric unless it says otherwise. No references are allowed except the problem statement and Go's standard library documentation.

### 6.1 — Rubric calibration

Directory: `phase-6/6-1-calibration/`

**Lab** `a-self-score`**:**

- [ ] Pick two of your earlier solutions and score each against the interview rubric.
- [ ] Compare your scores with the instructor's.

**Assessed by:** Reasoning.

**Pass criteria:** your score is within one point of the instructor's in every category.

**If it fails:** discuss each category where you differed, then score a third solution.

### 6.2 — Verbal walkthrough

Directory: `phase-6/6-2-verbal/`

**Lab** `a-walkthrough`**:** For three unseen problems, work through the before-coding checklist out loud or as a written transcript before writing code:

- [ ] Restate the problem and ask clarifying questions.
- [ ] Work a small example by hand.
- [ ] Describe the brute-force solution and its complexity.
- [ ] Propose an improvement and its complexity.

**Assessed by:** Interview simulation (Understanding, Approach, Complexity, and Communication categories only).

**Pass criteria:** each walkthrough takes 10 minutes or less and scores at least 7/8 on those four categories.

**If it fails:** practice the checklist on two problems you have already solved, then retry with new problems.

### 6.3 — Timed easy problems

Directory: `phase-6/6-3-timed-easy/`

**Lab** `a-easy`**:**

- [ ] Solve three unseen easy problems, 20 minutes each.

**Assessed by:** Interview simulation.

**Pass criteria:** all three pass the instructor's tests within the time limit.

**If it fails:** log the cause in `mistakes.md`, do the remediation for that topic's lesson, and retry with new problems.

### 6.4 — Timed medium problems

Directory: `phase-6/6-4-timed-medium/`

**Lab** `a-medium`**:**

- [ ] Solve four unseen medium problems, 40 minutes each.

**Assessed by:** Interview simulation.

**Pass criteria:** at least three of four pass the instructor's tests within the time limit, each scoring 11/14 or better.

**If it fails:** log each failure in `mistakes.md`, repeat the lesson for its pattern, and retry with new problems.

### 6.5 — Follow-up questions

Directory: `phase-6/6-5-follow-ups/`

**Lab** `a-follow-ups`**:** After each of four solved problems, the instructor asks a follow-up: a larger input, a streaming input, less memory, or a changed constraint.

- [ ] Adapt the solution, or describe a correct alternative with its complexity.
- [ ] Say plainly when you would describe a solution rather than code it fully.

**Assessed by:** Reasoning, Complexity.

**Pass criteria:** at least three of four follow-ups are handled correctly.

**If it fails:** for each miss, write the alternative after the session and revisit the lesson that covers it.

### 6.6 — Go without references

Directory: `phase-6/6-6-go-fluency/`

**Lab** `a-from-memory`**:** In 15 minutes, with no documentation, implement the following so the instructor-provided tests pass:

- [ ] A `container/heap` implementation and its use.
- [ ] A sort with a custom multi-key comparator.
- [ ] BFS on a grid.
- [ ] A struct with methods, and a map of slices.

**Assessed by:** Implementation.

**Pass criteria:** the code compiles on the first or second attempt and the provided tests pass.

**If it fails:** list every syntax gap in `mistakes.md`, drill each one for five minutes, and retry the next day.

### 6.7 — Recurring mistakes

Directory: `phase-6/6-7-mistakes/`

**Lab** `a-review`**:**

- [ ] Confirm every entry in `mistakes.md` has the problem, category, fix, and remediation drill.
- [ ] Re-solve the three oldest unresolved problems from an empty directory.

**Assessed by:** Implementation, Reasoning.

**Pass criteria:** all three re-solves pass, and you explain what you now do differently.

**If it fails:** keep the entry open, repeat its remediation drill, and retest it in a later session.

### 6.8 — Mock interviews

Directory: `phase-6/6-8-mocks/`

**Lab** `a-mocks`**:** Full 45-minute mocks, each with one medium problem or two easy problems, plus follow-ups.

- [ ] Complete at least four mocks.
- [ ] Record the rubric score for each in the lab's `notes.md`.

**Assessed by:** Interview simulation.

**Pass criteria:** three consecutive mocks score 11/14 or higher.

**If it fails:** target the lowest-scoring rubric category with the matching earlier lesson, then continue the mocks.

### Program exit assessment

Directory: `phase-6/exit/`

Each line maps to the README's definition of success.

- [ ] Solve unfamiliar problems on your own: pass two unseen medium problems in a single session.
- [ ] Recognize patterns: score at least 10/12 on a fresh classification set, as in 3.7.
- [ ] Choose data structures and justify them: score at least 8/10 on fresh scenarios, as in 2.9.
- [ ] Analyze complexity: score at least 9/10 on a mixed complexity quiz.
- [ ] Write idiomatic Go under interview conditions: pass 6.6 again.
- [ ] Test and debug: fix an instructor-supplied faulty implementation, as in 1.7, within 20 minutes.
- [ ] Explain reasoning: one solution's explanation meets the reasoning rubric.
- [ ] Adapt a known approach: handle two follow-ups correctly, as in 6.5.
- [ ] Solve under time limits: pass one more mock at 11/14 or higher.
- [ ] Retain what you learned: re-solve one lab from each of Phases 2 through 5 from an empty directory.

The goal is not a memorized set of solutions. It is the ability to solve a problem you have never seen before.
