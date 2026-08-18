---
tldr: "Defines how a plan is staged, named, structured, and archived in plans/."
when_to_use: "Use when creating, executing, reviewing, or archiving a plan under plans/."
---

# Plans Organization Policy

## Scope

This policy governs `plans/`, the repository's working record of change. A plan explains why work exists, what it depends on, and what evidence proves it finished. Plans are temporary and belong to delivery; `docs/` serves readers and `repo-governance/` holds rules, so neither is a home for a plan. The three plan workflows — [plan-planning](../workflows/plan-planning.md), [plan-quality-gate](../workflows/plan-quality-gate.md), and [plan-execution](../workflows/plan-execution.md) — carry out what this policy defines.

## When a Plan Is Required

A plan is required for application work, infrastructure work, and repository rule changes. Drills and study are not planned: the owner practices by hand and tracks the session in a harness task list, as the [task tracking policy](task-tracking-policy.md) requires. A drill becomes plan-worthy only when it stops being practice and starts being repository work.

## Rules

Read the rule you need rather than the whole set:

- [Folder Structure](plans-organization-policy/folder-structure.md) — the four lifecycle stages.
- [Plan Naming](plans-organization-policy/plan-naming.md) — stage-aware folder names.
- [Two-Pager Template](plans-organization-policy/two-pager-template.md) — what an idea contains.
- [Five-Document Structure](plans-organization-policy/five-document-structure.md) — the plan files.
- [Delivery Checklists](plans-organization-policy/delivery-checklists.md) — granularity, clarity, executor tags.
- [Phases and Gates](plans-organization-policy/phases-and-gates.md) — natural pauses.
- [Knowledge Capture](plans-organization-policy/knowledge-capture.md) — draining `learnings.md`.
- [Lifecycle Moves](plans-organization-policy/lifecycle-moves.md) — starting, completing, reopening.

## Delivery

Plans deliver directly to `main`. This repository runs no pull-request flow, no worktrees, and no delivery modes: a phase ends, its gate passes, and the work is committed and pushed. The [commit hook policy](../development/commit-hook-policy.md) still governs every commit.

## Diagrams and Secrets

Plans follow the [Markdown style policy](markdown-style-policy.md) without exception, so diagrams are terminal-first ASCII, never Mermaid. Plan documents are committed, so they never carry a secret: name the variable and its location instead of its value.

## Verification

`plans/` is outside `repo-governance/`, so no word limit applies to a plan and no automated check reads one. The [plan-quality-gate](../workflows/plan-quality-gate.md) workflow is the verification: `plan-checker` reports findings against these rules and `plan-fixer` resolves them.
