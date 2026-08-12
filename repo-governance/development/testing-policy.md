---
tldr: "Defines cacheable quick tests and explicit uncached integration tests."
when_to_use: "Use when adding, changing, or running project test and quality targets."
---

# Testing Policy

## Scope

This policy applies to every Nx project in this repository.

## Quick Tests

Each project must expose a cacheable `test:quick` target. It runs deterministic
unit tests with collaborators mocked, the project's static type check when its
language provides one, and linting. The pre-push hook runs `test:quick` only for
projects affected relative to `origin/main`; it intentionally uses Nx's local
task cache.

## Integration Tests

Use an uncached `test:integration` target for behavior that crosses project or
external-system boundaries. It may use real collaborators and is not enforced
by pre-push because it can be slow or environment-dependent. Run it explicitly:

```sh
npm run test:integration
```

## Tooling

TypeScript projects use `tsc --noEmit` plus Biome linting. Go projects use
`go vet` and fail when `gofmt -l` reports an unformatted source file. Keep
dependency versions exact and audit any added npm dependency.

## Verification

Run `npm test`, `npm run lint`, `npm run typecheck`, and any applicable
`npm run test:integration` before handing off behavior changes. To inspect the
pre-push selection without pushing, run:

```sh
npm exec nx -- affected -t test:quick --base=origin/main --head=HEAD
```
