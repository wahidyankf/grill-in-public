# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Canonical Guidance

[`AGENTS.md`](AGENTS.md) is authoritative and links to the detailed policies in [`repo-governance/`](repo-governance/README.md). Read it first. This file is a derivative: it adds only Claude Code-specific detail and must never contradict or restate canonical guidance. See the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md), and run the [Align Agent Harnesses](repo-governance/workflows/align-agent-harnesses.md) workflow whenever canonical guidance changes.

Like `AGENTS.md`, this file has a hard 500-word limit enforced by `npm run check:governance`. Link to a focused document instead of growing it.

The repository also supports Codex and opencode, which read `AGENTS.md` directly; see the [agent harness support policy](repo-governance/conventions/agent-harness-support.md). The `drill-reviewer` and `repo-explorer` subagents in `.claude/agents/` are mirrored for both, and must stay at parity.

This is a hands-on interview-preparation workspace: review, format, and give feedback by default, rather than solving the owner's drills.

## Commands

```sh
npm install                  # exact-pinned deps and Husky hooks
npm run build                # nx run-many -t build
npm test                     # cached test:quick across projects
npm run test:integration     # uncached, excluded from pre-push
npm run run:dummy            # build and run the demo CLI
npm run format               # Prettier, the formatting source of truth
npm run check:governance     # 500-word harness and governance limits
npm run check:markdown-links # repository-local link validation
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
apps/dummy-app  --depends on-->  libs/dummy-lib   (@swe-grilling/dummy-lib)
apps/badakmini-cli                                (Go validation CLI)
```

Every Nx target is a raw `command` target, with no plugins, generators, or executors; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

TypeScript projects compile with `tsc` into `<project>/dist`, and tests run against that compiled JavaScript with `node --test`, never against `src/`. `test:quick` therefore declares `dependsOn: ["build", "typecheck", "lint"]`, and running one test file directly needs a build first. Cross-project imports use the package name.

`apps/badakmini-cli` owns recurring repository-local checks, including the word limit above. `cv/` holds career evidence; read [cv/README.md](cv/README.md) before touching it.

## Commit Attribution

The [commit hook policy](repo-governance/development/commit-hook-policy.md) forbids AI attribution in commits and pull requests. `.claude/settings.json` enforces it here: `attribution.commitTrailers` is `false`, with empty `commit` and `pr` text and `sessionUrl` disabled, which suppresses the `Co-Authored-By` trailer and the generated-with footer. Never add either by hand.

## Quality Gates

Pre-commit formats staged files. Pre-push requires `origin/main`, runs affected `test:quick`, runs the governance check when the push touches `AGENTS.md`, `CLAUDE.md`, or `repo-governance/`, and always validates Markdown links. Commit messages go through commitlint. See the [commit hook policy](repo-governance/development/commit-hook-policy.md); the link check reads Git-tracked files, so `git add -N` a new document before trusting a local run.

## Writing Here

Follow the [code commentary policy](repo-governance/development/code-commentary-policy.md): linters enforce only a minimum shape, so review the reasoning a learner would need. Follow the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md) for unwrapped prose and ASCII diagrams, and [root cause orientation](repo-governance/principles/root-cause-orientation.md) when something fails.
