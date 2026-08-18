---
tldr: "Explains how to build, test, lint, and run the raw-Nx demonstration workspace."
when_to_use: "Use when working with Nx projects, their quality targets, or the Dummy App."
---

# Run the Nx Workspace

Use this guide to build, test, or run the TypeScript workspace after installing the repository's dependencies.

## Prerequisites

From the repository root, install the pinned dependency tree:

```sh
npm install
```

## Build and Test

Build every project, type-check and lint the workspace, then run cached unit tests:

```sh
npm run build
npm run typecheck
npm run lint
npm test
```

`npm test` runs the cacheable `test:quick` targets. They use mocked collaborators for units and include Badak Mini's Go tests.

## Run Integration Tests

Run real cross-project checks separately when needed:

```sh
npm run test:integration
```

For example, Dummy App's integration test verifies that it consumes the greeting exported by `libs/dummy-lib`. Integration tests are not part of the pre-push hook.

## Check Governance Guidance and Markdown Links

Run the governance check after changing the agent instruction files, `AGENTS.md` and `CLAUDE.md`, Markdown files under `repo-governance/`, or a `README.md` under `.claude/`, `.codex/`, or `.opencode/`:

```sh
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
```

The governance command uses the [Badak Mini](../../apps/badakmini-cli/README.md) Go CLI and runs automatically during a push that changes those paths. The parity command compares the subagents, skills, and commands each harness exposes, and runs during a push that changes a harness directory. The link command validates every Git-tracked Markdown file and runs during every push. A fourth command, `npm run check:rule-change`, runs during pre-commit and names the [Rules Propagation](../../repo-governance/workflows/rules-propagation.md) workflow when a staged change touches a rule path, plus [Harness Alignment](../../repo-governance/workflows/harness-alignment.md) when a harness reads that path; it reports without blocking. It cannot see an untracked new document, so run `git add -N <file>` before trusting a local run on newly created Markdown.

Before each push, Nx runs cached `test:quick` targets for projects affected relative to `origin/main`. See the shared [testing policy](../../repo-governance/development/testing-policy.md) for the target rules.

## Run the Demonstration App

```sh
npm run run:dummy
```

Expected output:

```text
Hello, Wahidyan!
```

Nx discovers the projects from their `project.json` files. Inspect them with `npx nx show projects`. This repository uses only raw Nx command targets; see the [Nx workspace policy](../../repo-governance/development/nx-workspace-policy.md) before adding Nx tooling.
