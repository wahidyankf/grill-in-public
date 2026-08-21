# Repository Guidelines

## Purpose and Working Style

Grind in Public is a personal lifelong-learning workspace for software engineering. Complete drills yourself before using tools for review, formatting, or feedback. Keep each exercise focused on one skill; record reasoning when useful. Before starting any task, create a granular task list and keep it synchronized as each item resolves; see the [task tracking policy](repo-governance/conventions/task-tracking-policy.md). Treat maintained code and tooling as liabilities that must prove net recurring value; see [maintenance value](repo-governance/principles/maintenance-value.md). Resolve an open decision by grilling with options, not prose; see the [grilling-with-options policy](repo-governance/conventions/grilling-with-options-policy.md).

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public), [ose-primer](https://github.com/wahidyankf/ose-primer), and [beaver-nest](https://github.com/wahidyankf/beaver-nest) as read-only references; local rules govern. For CV work, read [cv/README.md](cv/README.md).

## Rule Changes and Audience

Before changing repository rules, follow the [`rules-propagation` workflow](repo-governance/workflows/rules-propagation.md); a change to what a harness reads also requires the [`harness-alignment` workflow](repo-governance/workflows/harness-alignment.md). `README.md` and `docs/` serve people; instruction files serve agents; `repo-governance/` serves both. `CLAUDE.md` must defer to this file; see the [agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md). Every `docs/`, `repo-governance/`, `cv/`, `scripts/`, and harness directory requires an indexed README; see the [documentation index policy](repo-governance/documentation-index-policy.md).

## Project Structure

- `apps/` runnable applications; `libs/` reusable packages.
- `docs/` human-facing Diátaxis documentation.
- `repo-governance/` shared policies and workflows; `plans/` delivery plans; `specs/` Gherkin behavior.
- Root configs: `package.json`, `nx.json`, `tsconfig.base.json`.

Keep implementation and tests under `src/`; use lowercase-hyphenated project directories. Add assets only within the project needing them.

## Commands

[Workspace commands](repo-governance/development/workspace-commands.md) is canonical for every command, check, and hook: the common loop, the narrower runs, and what each Git hook does.

## Planning

Create plan documents only when the owner explicitly requests a plan. A requested plan is five documents in `plans/`, promoted from an idea in `ideas/` and staged through `backlog/`, `in-progress/`, and `done/`; see the [plans organization policy](repo-governance/conventions/plans-organization-policy.md). Author with [plan-planning](repo-governance/workflows/plan-planning.md), validate with [plan-quality-gate](repo-governance/workflows/plan-quality-gate.md), and run with [plan-execution](repo-governance/workflows/plan-execution.md). Plans deliver directly to `main`.

## Nx and Coding Conventions

Use Nx only as a raw task runner with `command` targets. Add no plugin, executor, or generator without explicit owner direction; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

Prettier is the source of truth; Markdown uses unwrapped paragraphs and terminal-first ASCII diagrams—see the [Markdown style policy](repo-governance/conventions/markdown-style-policy.md). The language target, naming, indentation, and import style follow the [code style policy](repo-governance/development/code-style-policy.md).

Comments must explain intent, flow, and non-obvious decisions without narrating syntax; see the [code commentary policy](repo-governance/development/code-commentary-policy.md).

## Testing and Commits

Each Nx project follows the [testing policy](repo-governance/development/testing-policy.md). Behavior is specified as Gherkin in `specs/` and implemented test-first, one scenario per red-green-refactor cycle; see the [specs policy](repo-governance/development/specs-policy.md) and the [TDD policy](repo-governance/development/tdd-policy.md).

Use Conventional Commits and split unrelated work into thematic commits. Follow the [commit hook policy](repo-governance/development/commit-hook-policy.md) for public-repository safety, attribution, and hook requirements. Do not commit `node_modules/` or unreviewed dependency updates.
