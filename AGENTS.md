# Repository Guidelines

## Purpose and Working Style

Grind in Public is a personal lifelong-learning workspace for software engineering. Complete drills yourself before using tools for review, formatting, or feedback. Keep each exercise focused on one skill; record reasoning when useful. Track multi-step work in a granular task list, updated as each item resolves; see the [task tracking policy](repo-governance/conventions/task-tracking-policy.md). Resolve an open decision by grilling with options, not prose; see the [grilling-with-options policy](repo-governance/conventions/grilling-with-options-policy.md).

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public) and [ose-primer](https://github.com/wahidyankf/ose-primer) as read-only references; local rules govern. For CV work, read [cv/README.md](cv/README.md).

## Rule Changes and Audience

Before changing repository rules, follow the [`rules-propagation` workflow](repo-governance/workflows/rules-propagation.md); a change to what a harness reads also requires the [`harness-alignment` workflow](repo-governance/workflows/harness-alignment.md). `README.md` and `docs/` serve people; instruction files serve agents; `repo-governance/` serves both. `CLAUDE.md` must defer to this file; see the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md). Every `docs/`, `repo-governance/`, and harness directory requires an indexed README; see the [documentation index policy](repo-governance/documentation-index-policy.md).

## Project Structure

- `apps/` holds runnable workspace applications.
- `libs/` holds reusable workspace packages.
- `docs/` holds human-facing Diátaxis documentation.
- `repo-governance/` holds shared policies and workflows; `plans/` holds delivery plans and `specs/` holds Gherkin behavior.
- Root configs: `package.json`, `nx.json`, `tsconfig.base.json`.

Keep implementation and tests under `src/`; use lowercase-hyphenated project directories. Add assets only within the project needing them.

## Commands

`npm install`, `npm run build`, and `npm test` cover the common loop; `npm run format` is the formatting source of truth. The [workspace commands](repo-governance/development/workspace-commands.md) document is canonical for every command, check, and hook.

## Planning

Application, infrastructure, and substantial rule work is planned before it starts; drills are not. A plan is five documents in `plans/`, staged through `ideas/`, `backlog/`, `in-progress/`, and `done/`; see the [plans organization policy](repo-governance/conventions/plans-organization-policy.md). Author with [plan-planning](repo-governance/workflows/plan-planning.md), validate with [plan-quality-gate](repo-governance/workflows/plan-quality-gate.md), and run with [plan-execution](repo-governance/workflows/plan-execution.md). Plans deliver directly to `main`.

## Nx and Coding Conventions

Use Nx only as a raw task runner with `command` targets. Do not add Nx plugins, plugin-specific executors, or generators without explicit owner direction; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

Use strict TypeScript with CommonJS-compatible Node output. Badak Mini is the standard-library Go CLI for repository-local validation; follow the [Badak Mini policy](repo-governance/development/badakmini-cli-policy.md) before extending it. Prettier is the source of truth; Markdown uses unwrapped paragraphs and terminal-first ASCII diagrams—see the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md). Use two-space indentation, `camelCase` variables and functions, `PascalCase` classes, and descriptive file names. Import internal libraries by package name, not relative cross-project paths.

Comments must explain intent, flow, and non-obvious decisions without narrating syntax; see the [code commentary policy](repo-governance/development/code-commentary-policy.md).

## Testing and Commits

Each Nx project must define a cacheable `test:quick` target and an uncached `test:integration` target; see the [testing policy](repo-governance/development/testing-policy.md). Behavior is specified as Gherkin in `specs/` and implemented test-first, one scenario per red-green-refactor cycle; see the [specs policy](repo-governance/development/specs-policy.md) and the [TDD policy](repo-governance/development/tdd-policy.md).

Use Conventional Commits and split unrelated work into thematic commits. The [commit hook policy](repo-governance/development/commit-hook-policy.md) governs commit messages, pull-request content, attribution, and `--no-verify`. Do not commit secrets, `node_modules/`, or unreviewed dependency updates.
