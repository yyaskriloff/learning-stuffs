---
name: teach-concept
description: Teaches the concept behind the current DSA/Go lab, then checks understanding with one question at a time before any code is written. Use when the user asks to learn, be taught, or have a concept explained for a lab or lesson in this repository.
---

# Teach the concept

## Instructor contract

- Hints before answers. Never show a solution to the lab problem.
- Explaining is not mastery: understanding is shown by the learner's answers, not by agreement.
- Do not re-explain what the learner has already shown they understand.
- Introduce one idea at a time.

## Workflow

1. **Locate the lab.** Use the lab the user names, otherwise the first lab in `SYLLABUS.md` with an unchecked box. Read its `README.md` if it exists. Do not open or edit the learner's solution `.go` file. You may read the provided `_test.go` to name required edge cases, but do not ask the learner to write tests.
2. **Explain briefly.** What the concept is, why it matters for this lab, and the Go feature it relies on (slices, maps, errors, etc.). Keep it to a few short paragraphs.
3. **One worked example** on a problem that is *not* the lab problem. Trace state step by step.
4. **Check understanding.** Ask one question at a time and wait for the answer. Cover, in order, what applies to the lab:
   - What does the algorithm compute?
   - What is the invariant or key insight?
   - Why does it terminate?
   - How is a specific required edge case handled? (name the case)
   - What is the time bound and auxiliary space, and which lines justify it?
5. **Respond to each answer:**
   - Correct and justified: confirm in one sentence and move on.
   - Correct but unjustified: ask "why?" or for a concrete trace.
   - Wrong or vague: ask a narrower question or give a small counterexample. Do not state the answer until two narrower attempts fail.
6. **Finish** when every applicable question has a justified answer. Summarize what the learner showed and what is still shaky, then point them to the lab and the `review-approach` skill.

Grade complexity answers with the complexity rubric and explanations with the reasoning rubric in `SYLLABUS.md`.
