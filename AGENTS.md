# Repository Guidelines

## Purpose and Working Style

Grill in Public is a personal interview-preparation workspace. Complete drills yourself before using tools for review, formatting, or feedback. Keep each exercise focused on one skill; record reasoning when useful. Track multi-step work in a granular task list, updated as each item resolves; see the [task tracking policy](repo-governance/conventions/task-tracking-policy.md). Resolve an open decision by grilling with options, not prose; see the [grilling-with-options policy](repo-governance/conventions/grilling-with-options-policy.md).

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public) and [ose-primer](https://github.com/wahidyankf/ose-primer) as read-only references; local rules govern. For CV work, read [cv/README.md](cv/README.md).

## Rule Changes and Audience

Before changing repository rules, read and follow the [`propagate-rules` workflow](repo-governance/workflows/propagate-rules.md); a change to what a harness reads also requires the [`align-agent-harnesses` workflow](repo-governance/workflows/align-agent-harnesses.md). `README.md` and `docs/` serve people; instruction files serve agents; `repo-governance/` serves both. `CLAUDE.md` must defer to this file; see the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md). Every `docs/`, `repo-governance/`, and harness directory requires an indexed README; see the [documentation index policy](repo-governance/documentation-index-policy.md).

## Project Structure

- `apps/` holds runnable workspace applications.
- `libs/` holds reusable workspace packages consumed by apps.
- `docs/` holds human-facing Diátaxis documentation.
- `repo-governance/` holds shared policies and workflows.
- Root configs: `package.json`, `nx.json`, `tsconfig.base.json`.

Keep implementation and tests under `src/`; use lowercase-hyphenated project directories. Add assets only within the project needing them.

## Commands

- `npm install` installs pinned dependencies and enables Husky hooks.
- `npm run build`, `npm run typecheck`, and `npm run lint` run Nx targets.
- `npm test` runs cached `test:quick` targets; `npm run test:integration` is excluded from pre-push.
- `npm run run:dummy` runs the demo CLI.
- `npm run format` and `npm run format:check` apply or verify Prettier.
- `npm run check:governance`, `npm run check:harness-parity`, and `npm run check:markdown-links` enforce word limits, equal harness capabilities, and local links. See the [commit hook policy](repo-governance/development/commit-hook-policy.md).
- `npm audit --audit-level=low` checks the locked dependency tree.

## Nx and Coding Conventions

Use Nx only as a raw task runner with `command` targets. Do not add Nx plugins, plugin-specific executors, or generators without explicit owner direction; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

Use strict TypeScript with CommonJS-compatible Node output. Badak Mini is the standard-library Go CLI for repository-local validation; follow the [Badak Mini policy](repo-governance/development/badakmini-cli-policy.md) before extending it. Prettier is the source of truth; Markdown uses unwrapped paragraphs and terminal-first ASCII diagrams—see the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md). Use two-space indentation, `camelCase` variables and functions, `PascalCase` classes, and descriptive file names. Import internal libraries by package name, not relative cross-project paths.

Comments must explain intent, flow, and non-obvious decisions without narrating syntax; see the [code commentary policy](repo-governance/development/code-commentary-policy.md).

## Testing, Commits, and Pull Requests

Each Nx project must define a cacheable `test:quick` target and an uncached `test:integration` target for boundary-crossing behavior. See the [testing policy](repo-governance/development/testing-policy.md).

Use Conventional Commits, enforced by Husky and commitlint. Split unrelated work into thematic commits. Never use `--no-verify` without explicit owner approval; see the [commit hook policy](repo-governance/development/commit-hook-policy.md). Keep pull requests focused, with motivation, commands run, linked issues when applicable, and screenshots only for visual changes. Do not commit secrets, `node_modules/`, or unreviewed dependency updates.
