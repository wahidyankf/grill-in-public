# Repository Guidelines

## Purpose and Working Style

SWE Grilling is a personal workspace for software-engineering interview
preparation. Favor deliberate, hands-on problem solving: complete coding drills
yourself before using tools for review, formatting, or feedback. Keep each
exercise focused on one skill and record reasoning when it adds value.

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public) and
[ose-primer](https://github.com/wahidyankf/ose-primer) as read-only references;
local rules govern. For CV work, read [cv/README.md](cv/README.md).

## Rule Changes and Audience

Before changing repository rules, read and follow the
[`propagate-rules` workflow](repo-governance/workflows/propagate-rules.md).
`README.md` and `docs/` serve people; this and related instruction files serve
AI agents. `repo-governance/` applies to both.

## Project Structure

- `apps/` contains runnable workspace applications.
- `libs/` contains reusable workspace packages consumed by apps.
- `docs/` contains human-facing Diátaxis documentation.
- `repo-governance/` contains shared policies and workflows.
- Root `package.json`, `nx.json`, and `tsconfig.base.json` configure tooling.

Keep implementation and tests together under `src/`; use focused,
lowercase-hyphenated project directories. Add assets only within the project
that needs them.

## Commands

- `npm install` installs exact-pinned dependencies and enables Husky hooks.
- `npm run build`, `npm run typecheck`, and `npm run lint` run Nx targets.
- `npm test` runs cached `test:quick` targets; `npm run test:integration` does
  not run during pre-push.
- `npm run run:dummy` builds and runs the TypeScript demonstration CLI.
- `npm run format` and `npm run format:check` apply or verify Prettier.
- `npm run check:governance` enforces governance-document word limits.
- `npm run check:markdown-links` validates every repository-local Markdown link;
  it runs on every pre-push. See the [commit hook policy](repo-governance/commit-hook-policy.md).
- `npm audit --audit-level=low` checks the locked dependency tree.

## Nx and Coding Conventions

Use Nx only as a raw task runner with `command` targets. Do not add Nx plugins,
plugin-specific executors, or generators without explicit owner direction; see
the [Nx workspace policy](repo-governance/nx-workspace-policy.md).

Use strict TypeScript and CommonJS-compatible Node output for TypeScript code.
Badak Mini uses its pinned Go toolchain and standard library; do not add Go
dependencies without owner direction. Prettier is the formatting source of truth. Use two-space indentation,
`camelCase` variables and functions, `PascalCase` classes, and descriptive file
names. Import internal libraries by package name, not relative cross-project
paths.

## Testing, Commits, and Pull Requests

Each Nx project must define a cacheable `test:quick` target for unit tests,
type checks, and linting. Mock collaborators in unit tests; put real
cross-project or external interactions in uncached `test:integration` targets.
Pre-push runs affected `test:quick` targets against `origin/main`. See the
[testing policy](repo-governance/testing-policy.md).

Use Conventional Commits, enforced by Husky and commitlint. Split unrelated work
into thematic commits. Never use `--no-verify` without explicit owner approval;
see the [commit hook policy](repo-governance/commit-hook-policy.md). Keep pull
requests focused and include motivation, commands run, linked issues when
applicable, and screenshots only for visual changes. Do not commit secrets,
`node_modules/`, or unreviewed dependency updates.
