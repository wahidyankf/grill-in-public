# Repository Guidelines

## Purpose and Working Style

SWE Grilling is a personal workspace for software-engineering interview
preparation. Favor deliberate, hands-on problem solving: complete coding drills
yourself before using tools for review, formatting, or feedback. Keep each
exercise focused on one skill and record reasoning where it adds value.

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public) and
[ose-primer](https://github.com/wahidyankf/ose-primer) as optional, read-only
references. Local guidance remains authoritative.

## Rule Changes

Before changing repository rules, immediately read and follow the
[`propagate-rules` workflow](repo-governance/workflows/propagate-rules.md).

## Project Structure and Module Organization

The repository is intentionally small:

- `README.md` explains the preparation goals and workflow.
- `package.json` defines local tooling; `package-lock.json` locks exact versions.
- `.husky/` contains Git hooks; `commitlint.config.cjs` defines commit rules.

As exercises are added, group them under `exercises/<topic>/<problem>/`. Keep
the implementation and its tests together, for example
`exercises/arrays/two-sum/solution.js` and `solution.test.js`. Add assets only
when an exercise needs them; keep them within that exercise directory.

## Build, Test, and Development Commands

- `npm install` installs the pinned dependencies and enables Husky hooks.
- `npm run format` applies Prettier to supported project files.
- `npm run format:check` verifies formatting without changing files.
- `npm run check:governance` enforces governance-document word limits.
- `npm audit` checks the locked dependency tree for known vulnerabilities.

`npm test` is currently a placeholder and intentionally fails. When adding the
first automated test, configure a test runner and replace that script in the
same change.

## Coding Style and Naming Conventions

Use CommonJS unless a change deliberately introduces another module system.
Prettier is the source of truth for formatting; use its default two-space
indentation and run `npm run format` before committing. Prefer clear,
lowercase-hyphenated directory names, `camelCase` for variables and functions,
and `PascalCase` for classes. Name files after the problem or responsibility,
not generic terms such as `helpers`.

## Testing Guidelines

Add a colocated `*.test.js` file for each exercise once test tooling is
introduced. Cover the expected solution, edge cases, and invalid or empty
inputs when relevant. A change that introduces behavior should include its
test command and result in the pull request description. No coverage threshold
exists yet; prioritize meaningful cases over a percentage target.

## Commit and Pull Request Guidelines

Use Conventional Commits, enforced by the Husky `commit-msg` hook and
commitlint. Examples: `feat(arrays): add two-sum drill` and
`docs: clarify practice workflow`. Valid types include `feat`, `fix`, `docs`,
`test`, `refactor`, `chore`, `build`, `ci`, `perf`, `style`, and `revert`.

Keep pull requests focused. Include a concise summary, the motivation, commands
run, and linked issues when applicable. Include screenshots only for visual
changes. Do not commit secrets, `node_modules/`, or unreviewed dependency
updates; retain exact dependency versions and run `npm audit` after changes.
