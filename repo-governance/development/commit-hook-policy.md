---
tldr: "Defines required Git hooks and the narrow exception for bypassing them."
when_to_use: "Use before committing, pushing, changing hooks, or considering a hook bypass."
---

# Commit Hook Policy

This policy applies to every contributor and AI agent working in this repository.

## Required Hooks

Do not bypass Git hooks with `--no-verify`, including for `git commit` and `git push`. Hooks enforce formatting, Conventional Commit messages, governance-document limits, and repository-local Markdown links. A failure must be investigated and resolved before continuing.

## Markdown Link Validation

Pre-push validates every Git-tracked repository Markdown file, not only changed documents, so moved or deleted files cannot leave a dangling internal reference. It checks repository-local file paths and Markdown heading fragments; external URLs are outside this validation. Run `npm run check:markdown-links` to diagnose a failure before pushing.

## Commit Attribution

Commits and pull requests must carry no AI attribution. Do not add a `Co-Authored-By` trailer for an assistant, a generated-with footer, or a session link, by hand or by tool default. The author and committer remain the repository owner. Configure each agent to suppress its own attribution; for Claude Code, `.claude/settings.json` sets `attribution.commitTrailers` to `false` with empty `commit` and `pr` text.

## Exception

A bypass is allowed only with explicit approval from the repository owner. Record the approval and reason in the commit or pull request so the skipped validation is visible and can be rerun promptly.
