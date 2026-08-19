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

## Run the Repository Checks

Run these checks after changing the agent instruction files, `AGENTS.md` and `CLAUDE.md`, Markdown files under `repo-governance/`, anything under `.agents/`, `.claude/`, `.codex/`, or `.opencode/`, or a project's `project.json` or `nx.json`:

```sh
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
npm run check:project-targets
```

What each of these checks enforces is stated once in [workspace commands](../../repo-governance/development/workspace-commands.md#repository-checks), together with a fifth command, `npm run check:rule-change`, that this guide does not run. Which hook runs each one, and on which pushes, is listed in the same reference under [hooks](../../repo-governance/development/workspace-commands.md#hooks). The rule-change and link commands read Git-tracked files, so `git add -N <file>` a new document before trusting a local run.

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
