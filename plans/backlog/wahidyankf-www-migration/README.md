# wahidyankf-www Migration

Move the personal website `wahidyankf-www` from `ose-public` into this repository, together with the projects it cannot run without.

## Context

`wahidyankf-www` is a Next.js 16 site living in `ose-public/apps/wahidyankf-www`, alongside a Playwright end-to-end project, nine Gherkin feature files, a Docker Compose development file, and a scheduled GitHub Actions workflow that deploys it. It is a personal site in a repository whose subject is something else, and this repository is where the owner's personal work now belongs.

The migration is not a copy. The site depends on three `ose-public` workspace libraries and on a Rust CLI for spec validation, and it pins a TypeScript version two majors behind this workspace. Each of those is a decision, not a file move.

## Scope

**In scope:** `apps/wahidyankf-www`, `apps/wahidyankf-www-fe-e2e`, `specs/apps/wahidyankf`, `infra/dev/wahidyankf-www`, and the library code the site imports.

**Out of scope:** every other `ose-public` app, the shared `ose-public` CI workflows beyond the site's own, and any redesign of the site. This plan moves working software without improving it; a migration that also refactors cannot say which change broke what.

**Affected projects in this repository:** `package.json` workspaces, `nx.json`, `tsconfig.base.json`, and the root toolchain. No existing project's behavior changes.

## Approach

Six phases, each ending green and delivering to `main`: baseline, libraries, application, specs and end-to-end, development infrastructure, then deployment. Libraries land before the app that imports them, and specs land with the code they describe. Deployment is last because it is the only phase whose failure is visible to the public.

## Documents

- [brd.md](brd.md) — why this is worth doing, and what would make it a mistake.
- [prd.md](prd.md) — user stories and Gherkin acceptance criteria.
- [tech-docs.md](tech-docs.md) — architecture, decisions, file impact, risks, rollback.
- [delivery.md](delivery.md) — the phased checklist.

## Open Decisions

These block promotion to `in-progress/` and need the owner's answer. Each is written as a grilling question in [tech-docs.md](tech-docs.md).

1. **Library boundary** — vendor the three `@open-sharia-enterprise/*` libraries into `libs/`, or copy only the components the site actually uses.
2. **Spec structure validation** — `specs:structure-validation` shells out to `ose-public`'s Rust `rhino-cli`, which does not exist here. Port the check to Badak Mini, or drop it.
3. **Linting** — the site runs `oxlint`; this workspace runs Biome plus ESLint commentary rules. One of them wins.
4. **TypeScript version** — this workspace is on 6.0.3, the site pins 5.8.3.
5. **Deployment** — this repository has no `.github/` at all. Recreate the scheduled deploy workflow, deploy another way, or leave the site undeployed until later.

## Quality Gate

Not yet run. This plan is authored but unvalidated; run [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) before promoting it.
