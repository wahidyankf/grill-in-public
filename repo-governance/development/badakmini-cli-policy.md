---
tldr: "Defines Badak Mini's role as the repository-local validation CLI."
when_to_use: "Use when adding or changing recurring repository validation checks."
---

# Badak Mini Policy

## Scope

Badak Mini is this repository's small, standard-library-only Go CLI for repository-local validation. Its production code owns recurring checks that protect repository health, such as instruction-size and Markdown-link validation; pinned Go `tool` dependencies support development validation only.

## Rules

- Do not create a Badak Mini command merely because a rule can be automated. First reuse an existing command, hook, standard tool, or clear manual review. A new command requires explicit repository-owner approval plus evidence that the failure is recurring, materially affects repository-wide correctness, and that centralized enforcement returns more value than its code, tests, documentation, execution time, false positives, upgrades, and eventual removal will cost.
- When an approved recurring repository-local check truly needs custom executable code, add it to Badak Mini rather than creating a standalone shell checker.
- Keep validation deterministic and offline so it can run in pre-push. Inspect the Git-tracked repository state when the check concerns committed content.
- Keep Badak Mini's production imports standard-library-only. Owner-approved, exact-pinned Go `tool` dependencies may support build, lint, test, or vulnerability checks, but must not become runtime dependencies.
- For each new check, add a focused command, an Nx target, unit tests, and human-facing usage documentation.
- Wire a check that blocks into pre-push, scoped to the paths that can break it, so a push that cannot fail the check does not pay for it. A check that no path narrows runs on every push.
- Wire a check that only reports into pre-commit, where the author can still act on what it says. A notice that arrives at push time asks for an amend rather than an edit.
- State the hook and its scope in the [workspace commands](workspace-commands.md) hook summary, which is canonical for what each hook runs.

See [Badak Mini's README](../../apps/badakmini-cli/README.md) for its current command surface and local verification commands.
