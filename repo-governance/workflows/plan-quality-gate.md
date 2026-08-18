---
tldr: "Runs plan-checker and plan-fixer in a loop until a plan has no findings left."
when_to_use: "Use after authoring or changing a plan, and before executing one."
---

# Plan Quality Gate

## Purpose

Validate that a plan is complete, accurate, and executable, then fix what is not. The gate exists because a plan is executed literally: an ambiguous checkbox does not produce a question at execution time, it produces a wrong action.

## When to Use

Use it after [plan-planning](plan-planning.md) authors a plan, after any edit to a plan under `backlog/` or `in-progress/`, and before [plan-execution](plan-execution.md) begins. A plan that has never passed this gate is not ready to start.

## Prerequisites

The plan exists with all five documents and has been through the [structural review](plan-planning/04-structural-review.md). Choose a severity level; see [severity and modes](plan-quality-gate/severity-and-modes.md). Strict is the default.

## Steps

1. Run `plan-checker` against the plan folder. It reads every document and reports findings by severity, citing `file:line`; see the [check and fix loop](plan-quality-gate/check-fix-loop.md).
2. Stop if no finding meets the chosen level. Otherwise continue.
3. Run `plan-fixer` on the findings at or above that level. It edits the plan documents only; it never touches the code the plan describes.
4. Re-run `plan-checker`. A fix that introduced a new finding is caught here rather than at execution.
5. Repeat until two consecutive runs report nothing at the chosen level, or seven cycles have passed.
6. Record the outcome in the plan's `README.md`: the level used, the cycles run, and the final status; see the [findings report](plan-quality-gate/findings-report.md).

## Verification

The gate passes when two consecutive `plan-checker` runs report no finding at the chosen level. One clean run is not enough: a single pass can reflect a checker that stopped early rather than a plan that is sound.

## Recovery

If the loop reaches seven cycles, stop and read the remaining findings yourself. Repeated failure to converge usually means the plan's approach is wrong rather than its wording, and `plan-fixer` cannot fix an approach. Rewrite the affected phase through [plan-planning](plan-planning.md) instead of iterating further.

If `plan-fixer` changes the meaning of a decision rather than its expression, revert that edit and resolve the decision with the owner. The fixer's mandate is clarity, not authority.
