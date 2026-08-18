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

`AGENTS.md` states when work is planned and which workflow runs each stage. What is Claude Code-specific: the quality gate spawns the `plan-checker` and `plan-fixer` subagents from `.claude/agents/`, and the same two are mirrored for the other harnesses.

## Architecture

```text
apps/dummy-app  --depends on-->  libs/dummy-lib   (@grind-in-public/dummy-lib)
apps/badakmini-cli                                (Go validation CLI)
```

Every Nx target is a raw `command` target: no plugins, generators, or executors; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

Tests run against compiled output rather than `src/`, so running one test file directly needs a build first; [workspace commands](repo-governance/development/workspace-commands.md#build-and-test) shows the invocation, and the [testing policy](repo-governance/development/testing-policy.md) owns what `test:quick` depends on.

`apps/badakmini-cli` owns repository-local checks, including the limit above. `cv/` holds career evidence; read [cv/README.md](cv/README.md) before touching it.

## Commit Attribution

The [commit hook policy](repo-governance/development/commit-hook-policy.md) forbids AI attribution in commits and pull requests, and `.claude/settings.json` carries the `attribution` settings that policy names. Never add a trailer, footer, or session link by hand.

## Quality Gates

Pre-commit formats staged files and announces the [Rules Propagation](repo-governance/workflows/rules-propagation.md) workflow when a staged path carries rules, and [Harness Alignment](repo-governance/workflows/harness-alignment.md) when a harness reads that path; a `PreToolUse` hook says the same before an edit. Commit messages go through commitlint. [Workspace commands](repo-governance/development/workspace-commands.md#hooks) lists what each hook runs and the caveats of running those checks locally; the [commit hook policy](repo-governance/development/commit-hook-policy.md) governs bypasses.

## Writing Here

Follow the [code commentary policy](repo-governance/development/code-commentary-policy.md): linters enforce a minimum shape only, so review the reasoning a learner needs. Follow the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md) for unwrapped prose and ASCII diagrams, and [root cause orientation](repo-governance/principles/root-cause-orientation.md) when something fails.
