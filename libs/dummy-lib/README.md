# Dummy Library

`@grill-in-public/dummy-lib` is a small TypeScript library that exports `createGreeting(name)`. It provides the reusable dependency consumed by [Dummy App](../../apps/dummy-app/README.md).

## Use

Within this workspace, import the package by name:

```ts
import { createGreeting } from "@grill-in-public/dummy-lib";

const message = createGreeting("Wahidyan");
```

## Quality Checks

From the repository root, run the cacheable unit, type-check, and lint suite:

```sh
npx nx run dummy-lib:test:quick
```

See [Run the Nx Workspace](../../docs/how-to/run-nx-workspace.md) for shared setup and workspace commands.
