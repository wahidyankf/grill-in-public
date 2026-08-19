---
tldr: "Indexes repeatable repository procedures."
when_to_use: "Use when a task has a defined sequence, required checks, or recovery steps."
---

# Repository Workflows

This directory contains repeatable procedures for working in Grind in Public. Use a workflow when a task has a defined sequence, required checks, or recovery steps that should be performed consistently by contributors and agents.

## Adding a Workflow

Create one Markdown file per procedure. The [document naming policy](../conventions/document-naming-policy.md) owns what to name it, and how a workflow that outgrows one file splits into a child directory. Keep each workflow narrowly scoped; link to related governance guidance instead of duplicating it.

## Available Workflows

- [Harness Alignment](harness-alignment.md) verifies that every supported harness receives the same rules through its instruction file, config, and subagents. Its detail lives in [`harness-alignment/`](harness-alignment/README.md).
- [Rules Propagation](rules-propagation.md) integrates changed repository rules without duplication or contradictions. Its detail lives in [`rules-propagation/`](rules-propagation/README.md).
- [README Refresh](readme-refresh.md) keeps root, project, documentation, and governance READMEs accurate before a thematic commit.
- [Rules Quality Gate](rules-quality-gate.md) runs `rules-checker` and `rules-fixer` over every rule-bearing file until no findings remain, composing Harness Alignment as a step. Its detail lives in [`rules-quality-gate/`](rules-quality-gate/README.md).
- [Plan Planning](plan-planning.md) turns a described change into a validated five-document plan under `plans/`.
- [Plan Quality Gate](plan-quality-gate.md) runs `plan-checker` and `plan-fixer` until a plan has no findings left.
- [Plan Execution](plan-execution.md) executes a plan phase by phase, delivering to `main` at each gate, then archives it.

The three plan workflows keep their detail in [`plan-planning/`](plan-planning/README.md), [`plan-quality-gate/`](plan-quality-gate/README.md), and [`plan-execution/`](plan-execution/README.md).

## Workflow Template

Each workflow should include:

1. **Purpose** — What outcome the procedure produces.
2. **When to use** — The task or condition that triggers it.
3. **Prerequisites** — Required tools, repository state, or access.
4. **Steps** — Ordered commands and actions, including expected results.
5. **Verification** — Checks that prove the outcome is complete.
6. **Recovery** — Safe next actions if a step fails, when applicable.

Use exact commands and paths where possible. Keep instructions current with the repository tooling, including the formatting, governance, and dependency checks defined in `package.json`.

## Maintenance

Update a workflow whenever its procedure changes. Move universally required, short rules to the root `AGENTS.md`; keep the detailed, conditional procedure here to preserve progressive disclosure.
