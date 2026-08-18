# Claude Code Harness

This directory holds the Claude Code configuration for SWE Grilling. Claude Code reads its repository rules from [`CLAUDE.md`](../CLAUDE.md), not from this directory; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- `settings.json` — project settings. It disables commit and pull-request attribution, as the [commit hook policy](../repo-governance/development/commit-hook-policy.md) requires, and registers the `PreToolUse` hook that announces the propagate-rules workflow before an edit to a rule file; see the [rule change trigger policy](../repo-governance/development/rule-change-trigger-policy.md).
- [`agents/`](agents/README.md) — the shared subagents available in this repository.

Claude Code also writes `settings.local.json` here for personal overrides. That file is ignored by Git and must not hold repository rules.
