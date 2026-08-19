---
tldr: "Sets the language target, naming, indentation, and import style for source code."
when_to_use: "Use when writing or reviewing code in any project, or when adding a language to the workspace."
---

# Code Style Policy

## Scope

This policy covers how source code is written: the language target, and the naming, indentation, and import conventions every project shares. Formatting is not covered — Prettier is the source of truth for that, and the [Markdown style policy](../conventions/markdown-style-policy.md) covers prose. What a comment must explain belongs to the [code commentary policy](code-commentary-policy.md).

## Language Target

Use strict TypeScript with CommonJS-compatible Node output. Badak Mini is the standard-library Go CLI for repository-local validation; follow the [Badak Mini policy](badakmini-cli-policy.md) before extending it.

## Naming and Layout

Use two-space indentation, `camelCase` variables and functions, `PascalCase` classes, and descriptive file names.

A descriptive file name is one a reader can act on without opening the file. `parse-project-file.ts` says what it holds; `utils.ts` says only that someone had nowhere to put it.

## Imports

Import internal libraries by package name, not by relative cross-project paths. A relative path across a project boundary compiles, so nothing stops it, and it hides the dependency from Nx — which then cannot tell that the importing project is affected when the imported one changes.

## Verification

```sh
npm run lint
npm run typecheck
```

Linting and type checking catch the mechanical part. Naming is reviewed by a person, because a name is only wrong relative to what the code does.
