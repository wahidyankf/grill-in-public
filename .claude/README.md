# Claude Code Harness

This directory holds the Claude Code configuration for Grind in Public. Claude Code reads its repository rules from [`CLAUDE.md`](../CLAUDE.md), which derives from the authoritative [`AGENTS.md`](../AGENTS.md), not from this directory; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- `settings.json` — project settings. It disables commit and pull-request attribution, as the [commit hook policy](../repo-governance/development/commit-hook-policy.md) requires, and registers the `PreToolUse` hook that announces the rule-change workflows before an edit to a rule file; see [harness pre-edit triggers](../repo-governance/development/harness-pre-edit-triggers.md).
- [`agents/`](agents/README.md) — the shared subagents available in this repository.
- [`skills/`](skills/README.md) — the shared skills available in this repository.

Claude Code also writes `settings.local.json` here for personal overrides. That file is ignored by Git and must not hold repository rules.
