---
tldr: "Defines the ordered, cacheable quick gate, strict coverage, and explicit integration tests."
when_to_use: "Use when adding, changing, or running project test and quality targets."
---

# Testing Policy

## Scope

This policy applies to every Nx project in this repository.

## Quick Tests

Each project must expose cacheable `typecheck`, `lint`, `test:unit`, `test:coverage`, and `test:quick` targets. `test:quick` is an ordered `nx:run-commands` aggregate with parallel execution disabled; it invokes those first four Nx target entry points in this order without copying their underlying commands:

```text
typecheck -> lint -> test:unit -> test:coverage
```

Unit tests are deterministic and replace external collaborators with test doubles. `test:coverage` reruns the unit suite with the language's native coverage instrumentation and fails below 95% aggregate statement coverage. Do not duplicate that executable threshold as metadata, omit runtime code, lower the threshold, or add broad exclusions to make the gate pass; an unavoidable generated-code exclusion requires the repository owner's explicit approval and a documented reason.

Pre-push invokes Nx affected with `origin/main` as the base and each pushed local commit as the head. Nx uses the project graph to run `test:quick` only for affected projects under `apps/` and `libs/`, so unrelated documentation changes do not run project tests and shared changes still reach every project they affect. The hook intentionally uses Nx's local task cache.

## Integration Tests

Use an uncached `test:integration` target for behavior that crosses project or external-system boundaries. It may use real collaborators and is not enforced by pre-push because it can be slow or environment-dependent. Run it explicitly:

```sh
npm run test:integration
```

A project whose behavior crosses no boundary defines no `test:integration` target at all. Its absence is the signal that the project has nothing to integration-test, so a placeholder that echoes and exits earns a passing run without testing anything and hides the same absence it claims to report.

## Tooling

TypeScript projects use TypeScript 6's `tsc --noEmit`, Biome, and project-local ESLint commentary checks. Go projects use `go vet` plus an exact-pinned GolangCI-Lint v2 tool configuration. That configuration starts from every linter, enables strict formatting, treats every finding as blocking, and disables only deprecated, duplicative, conflicting, or inapplicable checks with a nearby reason. A linter suppression must be narrow, specific, and explained; broad exclusions and issue caps are forbidden. Keep dependency versions exact; audit npm dependencies and scan Go module dependencies with the commands [workspace commands](workspace-commands.md#repository-checks) lists.

## Verification

Run `npm run typecheck`, `npm run lint`, `npm run test:unit`, `npm run test:coverage`, `npm run test:quick`, and any applicable `npm run test:integration` before handing off behavior changes. To inspect the pre-push selection without pushing, run:

```sh
npm exec nx -- affected -t test:quick --base=origin/main --head=HEAD
```
