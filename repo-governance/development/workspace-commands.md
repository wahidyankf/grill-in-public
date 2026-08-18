---
tldr: "Lists the npm and Nx commands that build, test, and validate this workspace."
when_to_use: "Use when running, testing, or validating any part of the workspace."
---

# Workspace Commands

This document is the canonical command reference. `AGENTS.md` and `CLAUDE.md` link here rather than restating it, so the list cannot drift between them.

## Setup

- `npm install` installs pinned dependencies and enables Husky hooks.

## Build and Test

- `npm run build`, `npm run typecheck`, and `npm run lint` run the matching Nx targets.
- `npm test` runs cached `test:quick` targets.
- `npm run test:integration` runs the uncached integration targets; pre-push skips them.
- `npm run run:dummy` runs the demo CLI.

Narrower runs:

```sh
npx nx run dummy-lib:test:quick
node --test apps/dummy-app/dist/index.test.js   # build first
go -C apps/badakmini-cli test ./internal/governance -run TestName
npx nx affected -t test:quick --base=origin/main --head=HEAD
```

TypeScript projects compile with `tsc` into `<project>/dist`, and tests run against that compiled JavaScript, never `src/`. `test:quick` therefore declares `dependsOn: ["build", "typecheck", "lint"]`, so running one test file directly needs a build first. See the [testing policy](testing-policy.md).

## Formatting

- `npm run format` and `npm run format:check` apply or verify Prettier, the formatting source of truth.

## Repository Checks

- `npm run check:governance` enforces the 500-word limit on instruction files, `repo-governance/` documents, and harness READMEs.
- `npm run check:harness-parity` compares the subagents, skills, and commands each harness exposes.
- `npm run check:markdown-links` validates repository-local Markdown links. It reads Git-tracked files, so `git add -N` a new document before trusting a local run.
- `npm run check:rule-change` announces the [rules-propagation](../workflows/rules-propagation.md) workflow for staged rule paths, and [harness-alignment](../workflows/harness-alignment.md) when a harness reads that path. It reports without blocking.
- `npm audit --audit-level=low` checks the locked dependency tree.

[Badak Mini](../../apps/badakmini-cli/README.md) implements the four `check:` commands. When a check fails, read the [Badak Mini policy](badakmini-cli-policy.md) before changing it; the usual fix is the document, not the checker.

## Hooks

Pre-commit formats staged files and announces the rule workflows. Pre-push requires `origin/main`, runs affected `test:quick`, runs the governance check when the push touches governance or a harness path, compares harness capabilities when a harness directory changes, and always validates Markdown links. See the [commit hook policy](commit-hook-policy.md).
