---
tldr: "Indexes repeatable repository procedures."
when_to_use: "Use when a task has a defined sequence, required checks, or recovery steps."
---

# Repository Workflows

This directory contains repeatable procedures for working in SWE Grilling. Use a workflow when a task has a defined sequence, required checks, or recovery steps that should be performed consistently by contributors and agents.

## Adding a Workflow

Create one Markdown file per procedure using a descriptive, lowercase-hyphenated name, such as `add-coding-exercise.md` or `update-dependency.md`. Keep each workflow narrowly scoped; link to related governance guidance instead of duplicating it.

## Available Workflows

- [Align Agent Harnesses](align-agent-harnesses.md) verifies that every supported harness receives the same rules through its instruction file, config, and subagents.
- [Propagate Rules](propagate-rules.md) integrates changed repository rules without duplication or contradictions.
- [Refresh READMEs](refresh-readmes.md) keeps root, project, documentation, and governance READMEs accurate before a thematic commit.

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
