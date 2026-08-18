# Technical Documentation

## What Exists Today

```text
ose-public/
├── apps/wahidyankf-www/            Next.js 16.2.6, React 19.2.6, Tailwind 4.2.1, port 3201, 36 source files
├── apps/wahidyankf-www-fe-e2e/     Playwright + playwright-bdd, generated features, step definitions
├── specs/apps/wahidyankf/          9 .feature files plus product, containers, components, system-context
├── infra/dev/wahidyankf-www/       docker-compose.yml for local development
└── .github/workflows/              wahidyankf-www-test-local-deploy-prod.yml, scheduled twice daily,
                                    delegating to _reusable-www-test-local-deploy.yml, deploying by
                                    pushing to the prod-wahidyankf-www branch
```

## Dependency Reality

The site imports three workspace libraries, not two:

```text
wahidyankf-www
├── @open-sharia-enterprise/web-ui          12 imports   9.6 MB   depends on web-ui-token + 7 Radix packages
├── @open-sharia-enterprise/web-ui-token     2 imports            design tokens
└── @open-sharia-enterprise/ts-env-loader     3 imports    40 KB   wraps dotenv
```

It also depends on tooling this repository does not have: `specs:structure-validation` shells out to `ose-public`'s Rust `rhino-cli`, and `lint` runs `npx oxlint@latest`.

## Toolchain Gap

| Concern    | This repository                  | The site as written       |
| ---------- | -------------------------------- | ------------------------- |
| Node       | 24.16.0, engines pinned          | unpinned                  |
| Nx         | 23.1.1                           | 22.5.4                    |
| TypeScript | 6.0.3                            | 5.8.3                     |
| Lint       | Biome 2.5.8 + ESLint 10          | oxlint (latest)           |
| Test       | `node --test` on compiled output | Vitest 4.1.0 on source    |
| CI         | none                             | GitHub Actions, scheduled |

The test row matters most: this workspace's [testing policy](../../../repo-governance/development/testing-policy.md) has TypeScript projects compile to `dist/` and test the compiled JavaScript, while the site tests source with Vitest and a JSDOM environment. A React component suite cannot reasonably run under `node --test` against `dist/`, so either the policy gains an explicit exception for browser-rendered projects or the site's suite is rewritten. The exception is the honest answer, and it is a rule change, so it runs through [rules-propagation](../../../repo-governance/workflows/rules-propagation.md).

## Target Shape

```text
grind-in-public/
├── apps/wahidyankf-www/
├── apps/wahidyankf-www-fe-e2e/
├── libs/web-ui/            (decision 1)
├── libs/web-ui-token/      (decision 1)
├── libs/ts-env-loader/     (decision 1)
├── specs/apps/wahidyankf/
└── infra/dev/wahidyankf-www/
```

Package names change from `@open-sharia-enterprise/*` to `@grind-in-public/*`, matching `@grind-in-public/dummy-lib` and the existing import-by-package-name rule.

## Decisions to Grill

Each of these changes what gets built and none can be defaulted safely.

**D1 — Library boundary.** Vendor all three libraries into `libs/`, or copy only the components the site uses into the app. Vendoring keeps the site's imports untouched and imports 9.6 MB of design system this repository otherwise has no use for; copying keeps the repository small and forks the components away from their upstream.

**D2 — Spec structure validation.** Port `specs structure validate` to Badak Mini, or drop the target. Badak Mini already owns repository-local checks and is the natural home; porting is real Go work for a check the repository has lived without.

**D3 — Lint.** Adopt oxlint for this project, run Biome across the site instead, or run both. The Nx workspace policy allows any raw `command` target, so all three are mechanically fine and the question is which one the owner wants to maintain.

**D4 — TypeScript.** Move the site to 6.0.3, or pin 5.8.3 for this project. Next.js 16 against TypeScript 6 is the risk; a per-project pin is allowed but leaves two compilers in one workspace.

**D5 — Deployment.** Recreate the scheduled GitHub Actions workflow here, deploy another way, or leave the site undeployed for now. This repository has no `.github/` directory at all, so this is a new capability rather than a port.

## Risks and Mitigation

- **Build fails on the toolchain gap.** Phase 1 builds the libraries and Phase 2 the app, each behind its own gate, so a failure names its cause.
- **Rendered output drifts silently.** Phase 3 compares rendered text against the source repository rather than trusting a green build.
- **The design system is larger than the site needs.** D1 is answered before Phase 1 begins.

## Rollback

Every phase is a separate commit on `main`. Rolling back is reverting the phase's commits; nothing is deleted from `ose-public` until the final phase, so the source copy remains the fallback for the entire migration.
