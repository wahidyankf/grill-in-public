# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Canonical Guidance

[`AGENTS.md`](AGENTS.md) is authoritative and links to detailed policies in [`repo-governance/`](repo-governance/README.md). Read it first. This file adds only Claude Code-specific detail and must never contradict or restate canonical guidance. See the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md), and run the [Harness Alignment](repo-governance/workflows/harness-alignment.md) workflow whenever canonical guidance changes.

Like `AGENTS.md`, this file lives under the [document word limit policy](repo-governance/conventions/document-word-limit-policy.md), which sets the limit and states how a document that reaches it is fixed.

Claude Code is one of three supported harnesses; Codex and opencode read `AGENTS.md`. See the [agent harness support policy](repo-governance/conventions/agent-harness-support.md) and the [agent vocabulary](repo-governance/conventions/agent-vocabulary.md). The subagents in `.claude/agents/` are mirrored for both, as the [harness capability parity policy](repo-governance/conventions/harness-capability-parity-policy.md) requires.

Default to reviewing and giving feedback rather than solving: `AGENTS.md` states the drill rule this follows from.

## Commands

The [workspace commands](repo-governance/development/workspace-commands.md) document is canonical for every command, narrower run, repository check, and hook. It is not summarized here, because a summary is what drifts.

## Planning

`AGENTS.md` states when work is planned and which workflow runs each stage. What is Claude Code-specific: the quality gate spawns the `plan-checker` and `plan-fixer` subagents from `.claude/agents/`, and the same two are mirrored for the other harnesses.

## Architecture

```text
apps/dummy-app  --depends on-->  libs/dummy-lib   (@grind-in-public/dummy-lib)
apps/badakmini-cli                                (Go validation CLI)
```

Every Nx target is a raw `command` target; the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md) owns what that forbids and the one exception.

Tests run against compiled output rather than `src/`, so running one test file directly needs a build first; [workspace commands](repo-governance/development/workspace-commands.md#build-and-test) shows the invocation, and the [testing policy](repo-governance/development/testing-policy.md) owns what `test:quick` depends on.

`apps/badakmini-cli` owns repository-local checks, including the limit above. `cv/` holds career evidence; read [cv/README.md](cv/README.md) before touching it.

## Commit Attribution

The [commit hook policy](repo-governance/development/commit-hook-policy.md) forbids AI attribution in commits and pull requests, and `.claude/settings.json` carries the `attribution` settings that policy names. Never add a trailer, footer, or session link by hand.

## Quality Gates

Pre-commit formats staged files and announces the [Rules Propagation](repo-governance/workflows/rules-propagation.md) workflow, plus [Harness Alignment](repo-governance/workflows/harness-alignment.md) where it applies; a `PreToolUse` hook says the same before an edit. The [rule change trigger policy](repo-governance/development/rule-change-trigger-policy.md) owns which paths announce which. Commit messages go through commitlint. [Workspace commands](repo-governance/development/workspace-commands.md#hooks) lists what each hook runs and the caveats of running those checks locally; the [commit hook policy](repo-governance/development/commit-hook-policy.md) governs bypasses.

## Writing Here

Follow the [code commentary policy](repo-governance/development/code-commentary-policy.md): linters enforce a minimum shape only, so review the reasoning a learner needs. Follow the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md) for unwrapped prose and ASCII diagrams, and [root cause orientation](repo-governance/principles/root-cause-orientation.md) when something fails.
