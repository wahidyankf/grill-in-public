---
tldr: "Describes how one phase is executed, including behavior cycles and notes."
when_to_use: "Use while working the items of a delivery phase."
---

# Phase Loop

A phase is worked in checklist order. The order is not decoration: it was authored so the repository stays coherent between items.

## Working an Item

Read the whole checkbox before acting. It names the path, the command, and the acceptance criterion, so it should need no interpretation. If it does need interpretation, that is a plan defect: fix the plan, note it in `learnings.md`, and continue.

Respect the executor tag. An `[AI]` item is performed directly. A `[HUMAN]` item stops execution and hands off to the owner with a statement of what remains; it is never performed on their behalf, and never ticked as though it were. An `[AI+HUMAN]` item is prepared fully, then handed over for the final action.

## Behavior Cycles

A behavior change runs RED → GREEN → REFACTOR against exactly one Gherkin scenario, per the [TDD policy](../../development/tdd-policy.md):

1. Write the test and run it. Confirm it fails, and that it fails for the stated reason. A test that passes immediately proves the assertion is wrong, not that the work is done.
2. Make the smallest change that passes it.
3. Refactor with the test still green.

Each of the three is its own checkbox, so a half-finished cycle is visible.

## Implementation Notes

When reality differs from the plan — a different path, an extra step, a command that needed a flag — record it inline under the checkbox in one line. The archived plan is then a record of what happened rather than what was intended.

## Learnings

When something surprises you, append it to `learnings.md` in the moment. Written later from memory, it records what you already believed rather than what actually happened.

## Stopping

Stop at a phase boundary, never mid-item. If you must stop mid-phase, leave the item open, note the state in `learnings.md`, and make sure the working tree still builds.
