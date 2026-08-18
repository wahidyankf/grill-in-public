# Dummy App

Dummy App is a small TypeScript command-line application used to practice the raw-Nx workspace setup. It consumes [`@grind-in-public/dummy-lib`](../../libs/dummy-lib/README.md) and prints a deterministic greeting.

## Run

From the repository root:

```sh
npm run run:dummy
```

Expected output:

```text
Hello, Wahidyan!
```

## Quality Checks

Run the cacheable unit, type-check, and lint suite with:

```sh
npx nx run dummy-app:test:quick
```

The separate integration suite verifies the real app-to-library boundary:

```sh
npx nx run dummy-app:test:integration
```

See [Run the Nx Workspace](../../docs/how-to/run-nx-workspace.md) for shared setup and workspace commands.
