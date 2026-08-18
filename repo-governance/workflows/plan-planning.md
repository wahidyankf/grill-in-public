---
tldr: "Turns a described change into a validated five-document plan under plans/."
when_to_use: "Use when application, infrastructure, or substantial rule work needs a plan before anyone starts it."
---

# Plan Planning

## Purpose

Turn a described change into a five-document plan that another session could execute without asking a question. The plan lands in `plans/backlog/` or `plans/in-progress/` and leaves this workflow only after the [plan-quality-gate](plan-quality-gate.md) workflow reports no findings.

## When to Use

Use it for application work, infrastructure work, and substantial rule work; the [plans organization policy](../conventions/plans-organization-policy.md) states the trigger and the line below which a rule change skips planning. Do not use it for a drill: practice is tracked in a harness task list, not planned.

## Prerequisites

Know which stage the plan targets. `backlog/` means prepared but not started, and is the default. `in-progress/` means execution follows immediately, and the plan is pushed to `main` before the first checklist item runs.

## Steps

1. [Explore before asking](plan-planning/01-exploration.md). Read the code, the governance documents, and the Git history that bear on the change. A question the repository already answers must not reach the owner.
2. [Grill the open decisions](plan-planning/02-grilling.md) with structured options, as the [grilling-with-options policy](../conventions/grilling-with-options-policy.md) requires. Unresolved decisions become guesses baked into a checklist.
3. [Author the five documents](plan-planning/03-plan-authoring.md) into `plans/<stage>/<identifier>/`, following the [five-document structure](../conventions/plans-organization-policy/five-document-structure.md).
4. [Review the structure](plan-planning/04-structural-review.md) against the checks listed there before handing off.
5. Run the [plan-quality-gate](plan-quality-gate.md) workflow. Fix what it reports; re-run until it is clean.
6. Update the stage's `README.md` index, then commit and push the plan to `main`. A plan that exists only locally cannot be picked up by a later session.

## Verification

```sh
npm run format:check
npm run check:markdown-links
```

The plan is ready when every document exists, every checklist item names a path, a command, and an acceptance criterion, every phase ends with a gate, and the quality gate reports no findings at strict level.

## Recovery

If the change turns out to be too small for five documents, it did not need a plan: delete the folder, do the work, and track it in the task list. If a decision cannot be resolved, leave the plan in `backlog/` with the open question written into `README.md` rather than guessing past it.
