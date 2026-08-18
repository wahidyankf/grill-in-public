# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Canonical Guidance

[`AGENTS.md`](AGENTS.md) is authoritative and links to detailed policies in [`repo-governance/`](repo-governance/README.md). Read it first. This file adds only Claude Code-specific detail and must never contradict or restate canonical guidance. See the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md), and run the [Align Agent Harnesses](repo-governance/workflows/align-agent-harnesses.md) workflow whenever canonical guidance changes.

Like `AGENTS.md`, this file has a hard 500-word limit enforced by `npm run check:governance`. Link to a focused document rather than growing it.

Claude Code is one of three supported harnesses; Codex and opencode read `AGENTS.md`. See the [agent harness support policy](repo-governance/conventions/agent-harness-support.md) and the [agent vocabulary](repo-governance/conventions/agent-vocabulary.md). The subagents in `.claude/agents/` are mirrored for both, as the [harness capability parity policy](repo-governance/conventions/harness-capability-parity-policy.md) requires.

Review, format, and give feedback by default rather than solving the owner's drills in this hands-on lifelong-learning workspace.

## Commands

```sh
npm install                  # pinned deps and Husky hooks
npm run build                # nx run-many -t build
npm test                     # cached test:quick
npm run test:integration     # uncached, skipped by pre-push
npm run run:dummy            # run demo CLI
npm run format               # Prettier, formatting source of truth
npm run check:governance     # 500-word limits
npm run check:harness-parity # equal subagents, skills, commands
npm run check:markdown-links # link validation
npm run check:rule-change    # announce propagate-rules for staged rules
```

Narrower runs:

```sh
npx nx run dummy-lib:test:quick
node --test apps/dummy-app/dist/index.test.js   # build first
go -C apps/badakmini-cli test ./internal/governance -run TestName
npx nx affected -t test:quick --base=origin/main --head=HEAD
```

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

Pre-commit formats staged files and announces the [Propagate Rules](repo-governance/workflows/propagate-rules.md) workflow when a staged path carries rules; a `PreToolUse` hook says the same before an edit. Pre-push requires `origin/main`, runs affected `test:quick`, runs the governance check when the push touches governance or a harness path, compares harness capabilities when a harness directory changes, and always validates Markdown links. Commit messages go through commitlint; see the [commit hook policy](repo-governance/development/commit-hook-policy.md). The link check reads Git-tracked files, so `git add -N` a new document first.

## Writing Here

Follow the [code commentary policy](repo-governance/development/code-commentary-policy.md): linters enforce a minimum shape only, so review the reasoning a learner needs. Follow the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md) for unwrapped prose and ASCII diagrams, and [root cause orientation](repo-governance/principles/root-cause-orientation.md) when something fails.
