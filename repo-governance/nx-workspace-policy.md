# Nx Workspace Policy

## Scope

This policy applies whenever repository work adds or changes Nx configuration,
targets, dependencies, generators, or executors.

## Required Approach

Use Nx as a raw task runner. Define project targets explicitly with the
`command` shorthand for Nx's command runner and use ordinary, exact-pinned npm
dependencies for compilation, testing, and execution.

Do not add framework-, language-, or platform-specific Nx plugins; do not use
plugin-specific generators or executors. In particular, avoid adding direct
`@nx/*` technology packages merely to scaffold or run a project.

## Exceptions

An exception requires explicit repository-owner direction that identifies the
needed plugin and the capability it provides. Record the decision in the change
that introduces it, keep its dependency version exact, and run the full locked
dependency audit afterward.

## Verification

Run `npx nx show projects` to confirm project discovery, then run the affected
Nx targets and `npm audit --audit-level=low`. Preserve `.nx/` and generated
project build directories in `.gitignore`.
