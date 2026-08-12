---
tldr: "Defines Badak Mini's role as the repository-local validation CLI."
when_to_use: "Use when adding or changing recurring repository validation checks."
---

# Badak Mini Policy

## Scope

Badak Mini is this repository's small, standard-library-only Go CLI for repository-local validation. Its production code owns recurring checks that protect repository health, such as instruction-size and Markdown-link validation; pinned Go `tool` dependencies support development validation only.

## Rules

- Add a new recurring repository-local check to Badak Mini rather than a new standalone shell checker, unless the repository owner directs otherwise.
- Keep validation deterministic and offline so it can run in pre-push. Inspect the Git-tracked repository state when the check concerns committed content.
- Keep Badak Mini's production imports standard-library-only. Owner-approved, exact-pinned Go `tool` dependencies may support build, lint, test, or vulnerability checks, but must not become runtime dependencies.
- For each new check, add a focused command, an Nx target, unit tests, and human-facing usage documentation. Wire it into pre-push only when the check must block every push.

See [Badak Mini's README](../../apps/badakmini-cli/README.md) for its current command surface and local verification commands.
