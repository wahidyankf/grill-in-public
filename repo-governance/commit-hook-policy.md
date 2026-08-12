# Commit Hook Policy

This policy applies to every contributor and AI agent working in this
repository.

## Required Hooks

Do not bypass Git hooks with `--no-verify`, including for `git commit` and
`git push`. Hooks enforce formatting, Conventional Commit messages, and
governance-document limits; a failure must be investigated and resolved before
continuing.

## Exception

A bypass is allowed only with explicit approval from the repository owner. Record
the approval and reason in the commit or pull request so the skipped validation
is visible and can be rerun promptly.
