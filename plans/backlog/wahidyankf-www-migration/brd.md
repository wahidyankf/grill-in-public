# Business Requirements

## Why This Matters

The owner's personal website sits in `ose-public`, a repository about Open Sharia Enterprise. The site is neither part of that product nor governed by its purpose, so every change to it costs a trip into an unrelated repository and inherits rules written for something else.

This repository is the owner's personal workspace. Moving the site here puts the personal work under the owner's own governance, and gives this repository a real deployed application — which is the thing that makes its planning, specs, and TDD policies apply to something with consequences rather than to a greeting function.

This is a judgment call, not a measured outcome. Nothing here is failing; the site works where it is. The argument is about where the work belongs, and it should be labeled as such rather than dressed up in numbers.

## Who It Affects

The owner, as the only maintainer and the only person whose attention the split repository currently costs. Site visitors are affected only if the migration breaks the deployment, which is why the deployment phase is last and separately gated.

## What Success Means

- The site builds, tests, and runs from this repository with the same behavior it has today.
- Its nine Gherkin feature files live in this repository's `specs/` tree and still bind to tests.
- The `ose-public` copy is removed only after the migrated copy is verified, never before.
- No existing project in this repository changes behavior.

## Non-Goals

- **No redesign.** Not one visual or content change lands in this plan. A migration that also improves things cannot attribute a regression.
- **No dependency upgrades** beyond what the move strictly requires. Version alignment is a decision recorded in `tech-docs.md`, not an opportunistic bump.
- **No new deployment platform.** If the deploy pipeline is recreated, it is recreated as it was.
- **No retrofit of specs onto existing projects.** `dummy-app`, `dummy-lib`, and Badak Mini gain specs when a plan touches them, not because this plan introduces `specs/`.

## Risks

- **The site stops deploying.** The current pipeline is a scheduled GitHub Actions workflow, and this repository has no `.github/` directory. Until that is resolved the migrated site is code that nobody serves.
- **The library dependency is larger than it looks.** `web-ui` is 9.6 MB and pulls in Radix and a token package; vendoring it imports a design system this repository has no other use for.
- **Toolchain drift.** This workspace runs TypeScript 6.0.3 and Nx 23; the site was built against TypeScript 5.8.3 and Nx 22. The build may fail for reasons unrelated to the move.
- **Partial migration.** A migration abandoned halfway leaves the site in two repositories, which is worse than either end state.

## What Would Make This a Mistake

If the deploy pipeline cannot be recreated here and the site would go unserved for an extended period, the move should wait. A personal site that nobody can visit is not an improvement in governance.
