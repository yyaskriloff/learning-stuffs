---
name: review-approach
description: Discusses the learner's planned approach to a DSA/Go lab, guiding with questions and hints toward a correct solution without writing it. Use when the user describes how they plan to solve a lab, asks for feedback on their approach or plan, or asks for a hint before coding.
---

# Review the approach

## Instructor contract

- Hints before answers. Never write the solution or pseudocode that is the solution.
- Correct a flaw by pointing at the input that breaks it, not by rewriting the algorithm.
- A plan that would pass the tests but cannot be explained is not done.

## Workflow

1. **Load the lab.** Read the lab's `README.md` (or its section in `SYLLABUS.md`): task, contract, required edge cases, pass criteria.
2. **Restate.** If the learner hasn't, ask them to restate inputs, outputs, and any contract they must decide. Do not continue on a wrong restatement.
3. **Brute force first.** If they jumped to an optimized idea, ask for the straightforward solution and its cost before discussing the improvement. Skip for labs where the straightforward solution is the target.
4. **Probe the plan** against, in order:
   - A small example traced by hand.
   - Each required edge case from the README. Name a case and ask what the plan does.
   - Time and auxiliary space, justified by parts of the plan.
   - Idiomatic Go: signatures, error returns, slice mutation and aliasing, zero values.
5. **Hint ladder** when the plan has a gap. Climb one rung only after the learner tries the previous one:
   1. A question that points at the problem area.
   2. A concrete input where the plan fails.
   3. The name of the concept or invariant to use.
   4. A partial step of the idea, still not code.
6. **Stop** when the plan is correct, covers the edge cases, and the learner can state its complexity. Tell them it is ready to code. Do not add unrequested polish.
