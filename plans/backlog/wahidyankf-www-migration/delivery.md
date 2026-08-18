# Delivery

**Executor legend**: `[AI]` an agent performs it — the default; `[HUMAN]` only the owner; `[AI+HUMAN]` an agent prepares, the owner performs the final action.

> **Blocked on decisions.** D1 through D5 in [tech-docs.md](tech-docs.md) must be answered before Phase 1. Phase 0 may run without them.

## Phase 0: Baseline

- [ ] [AI] Record this workspace's current state: run `npm install`, then `npm test`, `npm run check:governance`, `npm run check:harness-parity`, and `npm run check:markdown-links`. Acceptance: every command exits 0, and the outputs are saved to `evidence/phase-0-baseline.txt`.
- [ ] [AI] Record the source site's current state: from `ose-public`, run `npx nx run wahidyankf-www:build`. Acceptance: exits 0, and the output is saved to `evidence/phase-0-source-build.txt`.
- [ ] [AI] Capture the source site's rendered home page text to `evidence/phase-0-source-home.txt`, for the comparison in Phase 3.
- [ ] [HUMAN] Answer D1 through D5 in [tech-docs.md](tech-docs.md), recording each decision and its reason.

### Phase 0 Gate

> Every check below passes before Phase 1 begins. A failure is fixed inside Phase 0.

- [ ] [AI] `npm test` — exits 0.
- [ ] [AI] `ls plans/backlog/wahidyankf-www-migration/evidence/` — the three baseline files exist.
- [ ] [HUMAN] Every decision D1 through D5 is answered in `tech-docs.md`.

> **Pause Safety**: nothing has changed except this plan and its evidence folder. Safe to stop. Resume with `npm test`.

## Phase 1: Libraries

- [ ] [AI] Create `libs/ts-env-loader/` per D1, renaming the package to `@grind-in-public/ts-env-loader`. Acceptance: `libs/ts-env-loader/package.json` exists with the new name.
- [ ] [AI] Add `libs/ts-env-loader/project.json` with `build`, `typecheck`, `lint`, `test:quick`, and `test:integration` targets as raw `command` targets, following `libs/dummy-lib/project.json`. Acceptance: `npx nx show project ts-env-loader` lists all five.
- [ ] [AI] RED — **Gherkin (underpins)** the env-loader's existing unit tests. Port them and run `npx nx run ts-env-loader:test:quick`. Acceptance: the suite runs and fails only where behavior genuinely differs.
- [ ] [AI] GREEN — resolve each failure until `npx nx run ts-env-loader:test:quick` exits 0.
- [ ] [AI] REFACTOR — remove any `ose-public`-specific path or name left in the ported code. Acceptance: `grep -rn "open-sharia-enterprise" libs/ts-env-loader` returns nothing.
- [ ] [AI] Repeat the four steps above for `web-ui-token`. Acceptance: `npx nx run web-ui-token:test:quick` exits 0.
- [ ] [AI] Repeat the four steps above for `web-ui`. Acceptance: `npx nx run web-ui:test:quick` exits 0.
- [ ] [AI] Register the new packages in the root `package.json` workspaces if D1 vendored them. Acceptance: `npm install` succeeds and `npx nx show projects` lists them.

### Phase 1 Gate

- [ ] [AI] `npm test` — exits 0 across every project.
- [ ] [AI] `npm run format:check` — no formatting differences.
- [ ] [AI] `grep -rn "@open-sharia-enterprise" libs/` — returns nothing.

> **Pause Safety**: the libraries exist and pass their own checks; no application consumes them yet, and no existing project changed. Safe to stop. Resume with `npm test`.

- [ ] [AI] Commit the phase with a Conventional Commits message and push to `origin main`.

## Phase 2: Application

- [ ] [AI] Copy `apps/wahidyankf-www/` from `ose-public`, excluding `node_modules`, `.next`, and build output. Acceptance: `apps/wahidyankf-www/package.json` exists here.
- [ ] [AI] Rewrite the site's imports from `@open-sharia-enterprise/*` to `@grind-in-public/*`. Acceptance: `grep -rn "@open-sharia-enterprise" apps/wahidyankf-www/src` returns nothing.
- [ ] [AI] Apply D3 to the `lint` target and D4 to the TypeScript version. Acceptance: `npx nx run wahidyankf-www:lint` and `npx nx run wahidyankf-www:typecheck` both exit 0.
- [ ] [AI] Add `apps/wahidyankf-www/project.json` with the site's targets as raw `command` targets, omitting `specs:structure-validation` until D2 is implemented. Acceptance: `npx nx show project wahidyankf-www` lists `build`, `dev`, `typecheck`, `lint`, `test:unit`, `test:coverage`, `test:quick`.
- [ ] [AI] RED — **Gherkin (binds)** "The site builds from this repository"

```gherkin
Scenario: The site builds from this repository
  Given the workspace has installed its dependencies
  When "npx nx run wahidyankf-www:build" runs
  Then the command exits 0
  And a Next.js build output exists under "apps/wahidyankf-www/.next"
```

Run the build and record the failure. Acceptance: the build fails, and the reason is recorded in `learnings.md`.

- [ ] [AI] GREEN — resolve the build failures until `npx nx run wahidyankf-www:build` exits 0.
- [ ] [AI] REFACTOR — remove any workaround the build needed that the plan did not anticipate, or record why it must stay.
- [ ] [AI] RED — **Gherkin (binds)** "The site's quick checks pass under this workspace"

```gherkin
Scenario: The site's quick checks pass under this workspace
  Given the site project defines a "test:quick" target
  When "npx nx run wahidyankf-www:test:quick" runs
  Then the command exits 0
```

Run it and record the failures. Acceptance: the target runs and its failures are listed.

- [ ] [AI] GREEN — resolve each failure until `npx nx run wahidyankf-www:test:quick` exits 0.
- [ ] [AI] Raise the testing-policy exception for browser-rendered projects through the [rules-propagation](../../../repo-governance/workflows/rules-propagation.md) workflow, since this project tests source with Vitest rather than compiled output with `node --test`. Acceptance: the exception is written into the [testing policy](../../../repo-governance/development/testing-policy.md).

### Phase 2 Gate

- [ ] [AI] `npx nx run wahidyankf-www:build` — exits 0.
- [ ] [AI] `npx nx run wahidyankf-www:test:quick` — exits 0.
- [ ] [AI] `npm test` — every other project still exits 0.
- [ ] [AI] `npm run check:governance` — exits 0, including the amended testing policy.

> **Pause Safety**: the site builds and tests here, and the source repository is untouched. The site is not yet served from here. Safe to stop. Resume with `npx nx run wahidyankf-www:test:quick`.

- [ ] [AI] Commit the phase and push to `origin main`.

## Phase 3: Specs and End-to-End

- [ ] [AI] Copy `specs/apps/wahidyankf/` from `ose-public` into `specs/apps/wahidyankf/`. Acceptance: all nine `.feature` files are present.
- [ ] [AI] Update `specs/README.md` to list the new subject. Acceptance: the index links to `specs/apps/wahidyankf/README.md`.
- [ ] [AI] Verify every scenario satisfies the one-Given, one-When, one-Then rule in the [specs policy](../../../repo-governance/development/specs-policy.md). Acceptance: each violation is listed and fixed, or recorded as a deliberate exception.
- [ ] [AI] Implement D2. Acceptance: either `npx nx run wahidyankf-www:specs:structure-validation` exits 0 against a Badak Mini implementation, or the target is absent and the decision is recorded.
- [ ] [AI] Copy `apps/wahidyankf-www-fe-e2e/` and repoint its configuration at this workspace. Acceptance: `npx nx run wahidyankf-www-fe-e2e:test:e2e` runs to completion.
- [ ] [AI] RED — **Gherkin (binds)** "The rendered pages match the source repository"

```gherkin
Scenario: The rendered pages match the source repository
  Given both copies of the site have been built
  When the migrated home page and the source home page are compared
  Then their rendered text content is identical
```

Capture the migrated home page text to `evidence/phase-3-migrated-home.txt` and diff it against `evidence/phase-0-source-home.txt`. Acceptance: the diff is empty, or every difference is explained.

- [ ] [AI] GREEN — resolve each unexplained difference until the diff is empty.

### Phase 3 Gate

- [ ] [AI] `diff evidence/phase-0-source-home.txt evidence/phase-3-migrated-home.txt` — no output.
- [ ] [AI] `npx nx run wahidyankf-www-fe-e2e:test:e2e` — exits 0.
- [ ] [AI] `npm run check:markdown-links` — exits 0.

> **Pause Safety**: the migrated site renders identically to the source and its end-to-end suite passes. The source copy is still live. Safe to stop. Resume with the diff command above.

- [ ] [AI] Commit the phase and push to `origin main`.

## Phase 4: Development Infrastructure

- [ ] [AI] Copy `infra/dev/wahidyankf-www/docker-compose.yml`, adjusting any path that referenced `ose-public`. Acceptance: `grep -rn "ose-public" infra/` returns nothing.
- [ ] [AI] Add an `infra/README.md` index describing the directory. Acceptance: `npm run check:markdown-links` exits 0.
- [ ] [AI+HUMAN] Bring the compose stack up locally and confirm the site serves on port 3201. Acceptance: the home page responds 200; the owner confirms the page looks right.

### Phase 4 Gate

- [ ] [AI] `npm run check:markdown-links` — exits 0.
- [ ] [HUMAN] The local stack served the site and the owner confirmed it.

> **Pause Safety**: local development works from this repository. Deployment is unchanged and still runs from `ose-public`. Safe to stop.

- [ ] [AI] Commit the phase and push to `origin main`.

## Phase 5: Deployment and Source Removal

- [ ] [AI] Implement D5. Acceptance: either `.github/workflows/` contains the site's deploy workflow and a dry run succeeds, or the decision to defer deployment is recorded in `README.md`.
- [ ] [HUMAN] Confirm the deployed site is serving from the new pipeline before anything is removed. Acceptance: the owner states the live site is correct.
- [ ] [AI] RED — **Gherkin (binds)** "The source copy is removed only after verification"

```gherkin
Scenario: The source copy is removed only after verification
  Given the migrated site passes every gate in this plan
  When the removal step runs in the source repository
  Then the migrated copy is already serving and verified
```

Assert the precondition before removing anything. Acceptance: the owner's confirmation above is recorded in `evidence/phase-5-verification.txt`.

- [ ] [AI+HUMAN] Remove the site, its end-to-end project, its specs, its infra file, and its workflow from `ose-public` in one change, and open it for the owner to land there. Acceptance: `ose-public` no longer builds `wahidyankf-www`, and its indexes no longer list it.

### Phase 5 Gate

- [ ] [HUMAN] The live site serves from the new pipeline, or deferral is recorded.
- [ ] [AI] `evidence/phase-5-verification.txt` exists and records the confirmation.
- [ ] [AI] `npm test` — exits 0.

> **Pause Safety**: the site lives and serves from this repository, and the source copy is gone or explicitly retained. Safe to stop.

- [ ] [AI] Commit the phase and push to `origin main`.

## Phase 6: Knowledge Capture

- [ ] [AI] Triage every entry in `learnings.md` through the [knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md). Acceptance: every entry is routed or discarded with a one-line reason.
- [ ] [AI] Route each rule-shaped learning through [rules-propagation](../../../repo-governance/workflows/rules-propagation.md), and each harness-shaped one through [harness-alignment](../../../repo-governance/workflows/harness-alignment.md). Acceptance: each lands in exactly one canonical home.
- [ ] [AI] File any learning too large to land inline as a two-pager in `plans/ideas/`. Acceptance: each such entry has a corresponding file.

### Phase 6 Gate

- [ ] [AI] Every `learnings.md` entry is terminal, or the plan records `No generalizable learnings — <reason>`.
- [ ] [AI] `npm run check:governance` — exits 0.

> **Pause Safety**: every lesson has a durable home. The plan is ready to archive. Safe to stop.

## Phase 7: Archival

- [ ] [AI] Rename the folder to `YYYY-MM-DD__wahidyankf-www-migration` using the completion date and move it to `plans/done/`.
- [ ] [AI] Update `plans/in-progress/README.md` and `plans/done/README.md`.
- [ ] [AI] Commit the move and push to `origin main`.
