# Run the Nx Workspace

Use this guide to build, test, or run the TypeScript workspace after installing
the repository's dependencies.

## Prerequisites

From the repository root, install the pinned dependency tree:

```sh
npm install
```

## Build and Test

Build every project, type-check the workspace, and run the app test:

```sh
npm run build
npm run typecheck
npm test
```

The test proves that `apps/dummy-app` consumes the greeting exported by
`libs/dummy-lib`.

## Run the Demonstration App

```sh
npm run run:dummy
```

Expected output:

```text
Hello, Wahidyan!
```

Nx discovers the projects from their `project.json` files. Inspect them with
`npx nx show projects`. This repository uses only raw Nx command targets; see
the [Nx workspace policy](../../repo-governance/nx-workspace-policy.md) before
adding Nx tooling.
