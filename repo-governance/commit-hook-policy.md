# Commit Hook Policy

This policy applies to every contributor and AI agent working in this
repository.

## Required Hooks

Do not bypass Git hooks with `--no-verify`, including for `git commit` and
`git push`. Hooks enforce formatting, Conventional Commit messages,
governance-document limits, and repository-local Markdown links. A failure must
be investigated and resolved before continuing.

## Markdown Link Validation

Pre-push validates every Git-tracked repository Markdown file, not only changed
documents, so moved or deleted files cannot leave a dangling internal reference.
It checks repository-local file paths and Markdown heading fragments; external
URLs are outside this validation. A target must be Git-tracked. Run
`npm run check:markdown-links` to diagnose a failure before pushing.

## Exception

A bypass is allowed only with explicit approval from the repository owner. Record
the approval and reason in the commit or pull request so the skipped validation
is visible and can be rerun promptly.
