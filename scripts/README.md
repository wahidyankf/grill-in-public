# Scripts

This directory is reserved for small, repository-local automation scripts that
do not belong to an Nx project or a Git hook.

Keep scripts focused, portable, and well commented. Prefer adding repeatable
development tasks as Nx `command` targets; place hook orchestration in
`.husky/`. The `.gitkeep` file preserves this directory while it has no scripts.

Before adding a script, check whether [Badak Mini](../apps/badakmini-cli/README.md)
already provides the needed repository validation.
