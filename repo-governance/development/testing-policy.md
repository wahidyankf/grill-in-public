---
tldr: "Defines cacheable quick tests and explicit uncached integration tests."
when_to_use: "Use when adding, changing, or running project test and quality targets."
---

# Testing Policy

## Scope

This policy applies to every Nx project in this repository.

## Quick Tests

Each project must expose a cacheable `test:quick` target. It runs deterministic unit tests with collaborators mocked, the project's static type check when its language provides one, and linting. It reaches the last two through `dependsOn` rather than by reimplementing them, so `lint` and `typecheck` must exist as their own targets and be named there. The pre-push hook runs `test:quick` only for projects affected relative to `origin/main`; it intentionally uses Nx's local task cache.

## Integration Tests

Use an uncached `test:integration` target for behavior that crosses project or external-system boundaries. It may use real collaborators and is not enforced by pre-push because it can be slow or environment-dependent. Run it explicitly:

```sh
npm run test:integration
```

A project whose behavior crosses no boundary defines no `test:integration` target at all. Its absence is the signal that the project has nothing to integration-test, so a placeholder that echoes and exits earns a passing run without testing anything and hides the same absence it claims to report.

## Tooling

TypeScript projects use TypeScript 6's `tsc --noEmit`, Biome, and project-local ESLint commentary checks. Go projects use `go vet` plus the Go-module-pinned golangci-lint, which runs `gofmt`, Revive, and `nolintlint`. Keep dependency versions exact; audit npm dependencies and scan Go module dependencies with the commands [workspace commands](workspace-commands.md#repository-checks) lists.

## Verification

Run `npm test`, `npm run lint`, `npm run typecheck`, and any applicable `npm run test:integration` before handing off behavior changes. To inspect the pre-push selection without pushing, run:

```sh
npm exec nx -- affected -t test:quick --base=origin/main --head=HEAD
```
