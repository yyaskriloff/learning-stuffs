# DSA + Go Interview Training Program

## 1. How Our Training Will Work

We're going to treat this as a practical engineering apprenticeship, not a course where I explain algorithms and you passively read along.

The goal is to become confident solving unfamiliar data structures and algorithms (DSA) problems in interviews, while learning Go well enough to write clean, idiomatic code under pressure.

Since I'm already an experienced developer, we won't teach programming from scratch. Instead, we'll focus on algorithmic problem-solving, understanding the underlying mechanics, recognizing patterns, and translating those patterns into Go.

**The central rule: I will do the work myself.** ChatGPT will teach, ask questions, give exercises, review code, challenge reasoning, and adapt lessons based on demonstrated understanding. It won't immediately hand me solutions just because a problem is difficult.

### The Five-Step Learning Loop

1. **Understand the problem**
   - Explain the inputs, outputs, constraints, edge cases, and what a correct solution must do.

2. **Develop my own approach**
   - Propose a brute-force solution first.
   - Reason about its complexity.
   - Look for ways to improve it.

3. **Implement it in Go**
   - Write the code myself.
   - Run tests and debug failures.
   - Use Go's standard library where appropriate.

4. **Defend my solution**
   - Explain why it works.
   - Analyze time and space complexity.
   - Identify boundary cases and potential improvements.

5. **Prove I learned it**
   - Solve a variation without following the original solution.
   - Revisit difficult problems later to build lasting recall.

### Teaching Rules

- **Hints before answers:** Start with questions or small hints, escalating only when necessary.
- **No fake mastery:** Understanding an explanation isn't the same as independently solving a problem.
- **Tests are mandatory:** Use examples, edge cases, and randomized or property-based checks when useful.
- **Complexity must be justified:** Explain why a solution has a particular time and space complexity.
- **Go is part of the curriculum:** Learn slices, maps, pointers, structs, interfaces, and other relevant language features through actual problems.
- **Mistakes are diagnostic:** Identify whether the difficulty comes from the algorithm, Go, debugging, or reasoning, and adjust the exercises accordingly.

I can ask for explanations at any time. After an explanation, however, I must demonstrate that I can apply what I learned.

---

## 2. The Curriculum

The program progresses through six stages. Each stage builds on the previous one. Advancement depends on demonstrated competence, not merely finishing the material.

### Stage 1: Foundations and Go Fluency

Topics:

- Big O: time and space complexity
- Go functions, slices, maps, structs, and testing
- Loops, nested loops, and basic problem decomposition
- Brute-force solutions and optimization
- Writing tests and debugging systematically

**Milestone:** Implement simple solutions from scratch and explain their complexity.

### Stage 2: Core Data Structures

Topics:

- Arrays and slices
- Strings and byte/rune handling
- Hash maps and sets
- Linked lists
- Stacks and queues

**Milestone:** Recognize when a data structure simplifies a problem and implement the underlying mechanics yourself.

### Stage 3: Problem-Solving Patterns

Topics:

- Two pointers
- Sliding windows
- Prefix sums
- Binary search
- Sorting and intervals
- Frequency counting

**Milestone:** Identify a useful approach in an unfamiliar problem without being told which pattern to use.

### Stage 4: Trees, Heaps, and Graphs

Topics:

- Recursion and tree traversal
- Binary search trees and heaps
- Breadth-first search (BFS) and depth-first search (DFS)
- Graph representations
- Topological sorting
- Union-find

**Milestone:** Reason about traversal, connectivity, ordering, and recursive state.

### Stage 5: Advanced Patterns

Topics:

- Backtracking
- Greedy algorithms
- Dynamic programming
- Monotonic stacks and queues
- Tries and other specialized structures

**Milestone:** Break down problems that initially seem too complicated to solve directly.

### Stage 6: Interview Simulation

Topics:

- Timed, unseen problems
- Verbal problem-solving
- Go implementation without excessive references
- Complexity analysis and edge cases
- Follow-up questions and alternative solutions
- Review of recurring mistakes

**Milestone:** Consistently solve representative problems independently under interview conditions.

### Adaptive Pacing

The amount of time spent on each stage depends on demonstrated performance.

Some topics may take one session, while others may require several. Topics such as dynamic programming should not be rushed simply to check them off the list.

---

## 3. The Syllabus Becomes Executable Before We Start

[SYLLABUS.md](SYLLABUS.md) is the ordered lesson list for this program. It stays as written for now. The goal, before any lesson is started, is to revise every checkbox into a verifiable task.

A lesson can contain several labs. A lab can contain several tests. Each lab is one checkbox on the syllabus: a specific piece of work with a result that can be checked.

Explaining correctness and recognizing patterns are assessed on their own. A passing unit test verifies behavior. It does not, by itself, show that I can explain the algorithm or choose the pattern.

### How a checkbox is verified

Every checkbox uses one or more of these assessment types. The lab states which ones apply.

| Assessment type | Verification |
| --- | --- |
| Implementation | Automated tests verify outputs and behavior. |
| Edge cases | Dedicated tests verify boundary conditions. |
| Algorithm correctness | Property-based or randomized tests where appropriate. |
| Complexity | Written analysis, benchmark experiments where useful, and a grading rubric. |
| Reasoning | A concrete written or oral response evaluated against explicit criteria. |
| Interview simulation | Timed task scored against a defined rubric. |

### What every checkbox includes

1. **A specific task and deliverable.** What I build or produce, and where it lives.
2. **A clear definition of done.** The observable result that finishes the lab.
3. **A test plan or assessment rubric.** The automated tests, the written criteria, or both.
4. **Required edge cases.** The boundary conditions the lab must cover.
5. **A rule for passing or failing.** What must pass, and what a failure means.
6. **A remediation path if it fails.** The next drill or earlier lab to repeat before I retry.

### Directory layout

Each lesson has its own directory. Each lab inside that lesson has its own directory. A lab directory is self-contained: the task, my solution, the tests or rubric, the edge cases, the pass/fail rule, and the remediation path all live there.

```
<lesson>/
  <lab>/
  <lab>/
```

---

## 4. Weekly Training Schedule

Start with five sessions per week, approximately 45–60 minutes each.

| Session | Focus                 | Required Output                                 |
| ------- | --------------------- | ----------------------------------------------- |
| 1       | Learn a concept       | Working implementation or small experiment      |
| 2       | Guided problem        | My own Go solution                              |
| 3       | Independent problem   | Solution without step-by-step help              |
| 4       | New variation         | Apply the technique to a different problem      |
| 5       | Review and assessment | Retest an old problem and explain the reasoning |

On busy days, a 20-minute session is acceptable. Consistency matters more than having a perfect schedule.

### Progress Tracking

Every topic will have one of five statuses:

- **Not introduced:** We haven't studied it.
- **Learning:** I need guidance to apply it.
- **Developing:** I can solve familiar problems but struggle with variations.
- **Independent:** I can solve and explain it without help.
- **Retained:** I can solve a new variation after a delay.

This prevents confusing exposure to a topic with actual competence.

---

## 5. Learning Go Alongside Algorithms

I don't need to finish an entire Go course before starting DSA. We'll learn language features when they become useful.

### Early Lessons

Learn:

- Functions
- Slices and maps
- Zero values
- Loops
- Errors
- Go's testing conventions

### As Data Structures Become More Complex

Learn:

- Structs and pointers
- Values versus references
- Linked nodes
- Interfaces
- Generic data structures where useful

### As Algorithms Become More Complex

Learn:

- Recursion
- Closures where appropriate
- Sorting APIs
- Heaps
- Graph representations
- Memory allocation considerations

### Throughout the Program

Practice:

- Table-driven tests
- Benchmarks when useful
- Race-free code where relevant
- Idiomatic error handling
- Readable implementations

### Language and Dependency Policy

Generally, use Go's standard library rather than adding dependencies.

Understand the data structures being used rather than blindly relying on a library. However, don't waste time reimplementing every standard-library feature in production-style exercises.

The goal is to learn Go through algorithmic problem-solving rather than separating language study from DSA.

---

## 6. The Standard Problem-Solving Checklist

Every substantial problem should follow this process. The objective is to make it habitual, even under interview pressure.

### Before Coding

- [ ] Restate the problem and clarify constraints.
- [ ] Work through a small example by hand.
- [ ] Describe a straightforward solution.
- [ ] Analyze time and space complexity.
- [ ] Identify a better approach, if one exists.

### Implementation

- [ ] Implement the solution in Go.
- [ ] Test normal cases and edge cases.
- [ ] Explain correctness and complexity aloud.

I am not expected to find the optimal solution immediately. A correct brute-force solution is often the best starting point because it provides something concrete to reason about and improve.

---

## 7. How ChatGPT Should Teach Me

Act as my instructor, coding mentor, and interview coach.

For each lesson:

1. Introduce the concept and explain why it matters.
2. Demonstrate only what is necessary to establish the fundamentals.
3. Give me a practical task that requires me to think and write code.
4. Let me attempt the task before revealing a solution.
5. Review my reasoning and implementation.
6. Ask follow-up questions to test understanding.
7. Give hints before providing more direct guidance.
8. Require me to test and analyze my implementation.
9. Provide a variation to verify that I can transfer what I learned.
10. Track my progress and revisit weak areas.

### Avoid

- Giving complete solutions before I've had a chance to attempt the problem.
- Treating a correct answer as proof that I understand the algorithm.
- Introducing too many concepts at once.
- Overexplaining concepts I have already demonstrated that I understand.
- Advancing simply because we've completed a lesson.
- Making every exercise artificially difficult.

### Prioritize

- Hands-on exercises.
- Independent reasoning.
- Clear explanations of complexity.
- Debugging and test-driven verification.
- Realistic interview questions.
- Progressive difficulty.
- Retention and transfer of knowledge.

---

## 8. Definition of Success

By the end of the program, I should be able to:

- Solve unfamiliar DSA problems independently.
- Recognize common algorithmic patterns.
- Choose appropriate data structures and justify the choice.
- Analyze time and space complexity accurately.
- Write clean, idiomatic Go under interview conditions.
- Test edge cases and debug incorrect solutions.
- Explain my reasoning clearly to an interviewer.
- Adapt a known approach to a new problem.
- Solve problems under realistic time constraints.
- Retain concepts well enough to use them after a delay.

The program should be adapted to my upcoming interview timeline, strengths, weaknesses, and demonstrated progress.

**The ultimate objective is not to memorize solutions. It is to develop the ability to solve problems I have never seen before.**
