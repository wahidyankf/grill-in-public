---
tldr: "Executes a plan phase by phase, delivering to main at each gate, then archives it."
when_to_use: "Use when starting or resuming execution of a plan under plans/."
---

# Plan Execution

## Purpose

Execute a plan's `delivery.md` phase by phase until every item is ticked, every gate has passed, and the plan is archived in `plans/done/`.

## When to Use

Use it when a plan is ready to run. Resuming an interrupted plan uses this workflow too, starting from [resuming](plan-execution/04-finalization.md#resuming).

## Prerequisites

The plan has passed the [plan-quality-gate](plan-quality-gate.md) workflow, and any open finding has been accepted by the owner in writing. The plan sits in `plans/in-progress/`: move it out of `backlog/`, update both indexes, and push that move before running any checklist item, per the [lifecycle rules](../conventions/plans-organization-policy/lifecycle-moves.md).

## Steps

1. [Materialize the checklist into the harness task list](plan-execution/01-task-list-sync.md), one task per checkbox.
2. Run Phase 0 and record the baseline: dependencies installed, gates green before anything changes.
3. [Work the current phase](plan-execution/02-phase-loop.md) item by item, in checklist order, ticking each only after its change exists.
4. [Pass the phase gate, then commit and push](plan-execution/03-gates-and-pushes.md) to `main`. Do not begin the next phase while a gate item fails.
5. Repeat until the last delivery phase is complete.
6. [Run Knowledge Capture and archive](plan-execution/04-finalization.md) the plan to `plans/done/`.

## Verification

```sh
npm test
npm run format:check
npm run check:markdown-links
```

Execution is complete when every checkbox is ticked, every gate passed, every `learnings.md` entry reached a terminal state, and the plan folder sits in `plans/done/` with a completion-date prefix and both indexes updated.

## Recovery

A failing gate is fixed inside its own phase, never by continuing. If the fix changes the approach, stop and amend the plan through [plan-planning](plan-planning.md), then re-run the [plan-quality-gate](plan-quality-gate.md) workflow before resuming: an executed plan that no longer describes what happened is worse than no plan.

If a step turns out to be impossible, do not tick it and do not silently drop it. Record what blocked it in `learnings.md`, mark the item blocked in `delivery.md` with one line of explanation, and finish every item that does not depend on it.
