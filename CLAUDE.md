# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Canonical Guidance

[`AGENTS.md`](AGENTS.md) is authoritative and links to detailed policies in [`repo-governance/`](repo-governance/README.md). Read it first. This file adds only Claude Code-specific detail and must never contradict or restate canonical guidance. See the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md), and run the [Harness Alignment](repo-governance/workflows/harness-alignment.md) workflow whenever canonical guidance changes.

Like `AGENTS.md`, this file has a hard 500-word limit enforced by `npm run check:governance`. Link to a focused document rather than growing it.

Claude Code is one of three supported harnesses; Codex and opencode read `AGENTS.md`. See the [agent harness support policy](repo-governance/conventions/agent-harness-support.md) and the [agent vocabulary](repo-governance/conventions/agent-vocabulary.md). The subagents in `.claude/agents/` are mirrored for both, as the [harness capability parity policy](repo-governance/conventions/harness-capability-parity-policy.md) requires.

Review, format, and give feedback by default rather than solving the owner's drills in this hands-on lifelong-learning workspace.

## Commands

```sh
npm install                  # pinned deps and Husky hooks
npm run build                # nx run-many -t build
npm test                     # cached test:quick
npm run format               # Prettier, formatting source of truth
```

The [workspace commands](repo-governance/development/workspace-commands.md) document is canonical for every command, narrower run, repository check, and hook.

## Planning

Application, infrastructure, and rule work is planned; drills are not. Plans live in `plans/` as five documents and move through `ideas/`, `backlog/`, `in-progress/`, and `done/`; see the [plans organization policy](repo-governance/conventions/plans-organization-policy.md). Run [plan-planning](repo-governance/workflows/plan-planning.md), then [plan-quality-gate](repo-governance/workflows/plan-quality-gate.md), then [plan-execution](repo-governance/workflows/plan-execution.md). The gate uses the `plan-checker` and `plan-fixer` subagents in `.claude/agents/`; phases deliver directly to `main`.

## Architecture

```text
apps/dummy-app  --depends on-->  libs/dummy-lib   (@grind-in-public/dummy-lib)
apps/badakmini-cli                                (Go validation CLI)
```

Every Nx target is a raw `command` target: no plugins, generators, or executors; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

TypeScript projects compile with `tsc` into `<project>/dist`; tests run against that compiled JavaScript with `node --test`, never `src/`. `test:quick` therefore declares `dependsOn: ["build", "typecheck", "lint"]`, and running one test file directly needs a build first.

`apps/badakmini-cli` owns repository-local checks, including the limit above. `cv/` holds career evidence; read [cv/README.md](cv/README.md) before touching it.

## Commit Attribution

The [commit hook policy](repo-governance/development/commit-hook-policy.md) forbids AI attribution in commits and pull requests. `.claude/settings.json` enforces it: `attribution.commitTrailers` is `false`, `commit` and `pr` text are empty, and `sessionUrl` is off, so no `Co-Authored-By` trailer or generated-with footer appears. Never add either by hand.

## Quality Gates

Pre-commit formats staged files and announces the [Rules Propagation](repo-governance/workflows/rules-propagation.md) workflow when a staged path carries rules; a `PreToolUse` hook says the same before an edit. Pre-push runs affected tests, the governance and parity checks where they apply, and always validates Markdown links. Commit messages go through commitlint. The link check reads Git-tracked files, so `git add -N` a new document first. See the [commit hook policy](repo-governance/development/commit-hook-policy.md).

## Writing Here

Follow the [code commentary policy](repo-governance/development/code-commentary-policy.md): linters enforce a minimum shape only, so review the reasoning a learner needs. Follow the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md) for unwrapped prose and ASCII diagrams, and [root cause orientation](repo-governance/principles/root-cause-orientation.md) when something fails.
