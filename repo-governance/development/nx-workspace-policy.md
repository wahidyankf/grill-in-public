---
tldr: "Restricts Nx to raw commands and explicitly ordered command aggregates."
when_to_use: "Use when changing Nx configuration, targets, dependencies, generators, or executors."
---

# Nx Workspace Policy

## Scope

This policy applies whenever repository work adds or changes Nx configuration, targets, dependencies, generators, or executors.

## Required Approach

Use Nx as a raw task runner. Define a target that owns one command with the `command` shorthand. Define an aggregate that must run several existing targets in order with the built-in `nx:run-commands` executor, an explicit `options.commands` list, and `options.parallel` set to `false`. Invoke the existing Nx target entry points from that list instead of copying their underlying commands, so only the individual targets own those commands.

Use ordinary, exact-pinned npm or language-native dependencies for compilation, testing, and execution. The built-in command runner is not a technology plugin and does not justify adding an `@nx/*` package.

Do not add framework-, language-, or platform-specific Nx plugins; do not use plugin-specific generators or executors. In particular, avoid adding direct `@nx/*` technology packages merely to scaffold or run a project.

## Exceptions

An exception requires explicit repository-owner direction that identifies the needed plugin and the capability it provides. Record the decision in the change that introduces it, keep its dependency version exact, and run the full locked dependency audit afterward.

## Verification

Run `npx nx show projects` to confirm project discovery, then run the affected Nx targets and `npm audit --audit-level=low`. Preserve `.nx/` and generated project build directories in `.gitignore`.
