# Codex Harness

This directory holds the project-scoped Codex configuration for Grind in Public. Codex reads its repository rules from [`AGENTS.md`](../AGENTS.md), so no Codex-specific instruction file exists; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- `config.toml` — project settings. Codex loads this layer only for a trusted project, and it takes precedence over the user config.
- `hooks.json` — the `PreToolUse` hook that announces the rule-change workflows before an `apply_patch` edit; see the [rule change trigger policy](../repo-governance/development/rule-change-trigger-policy.md). Codex runs a project hook only after the owner trusts the project and approves the hook with `/hooks`.
- [`agents/`](agents/README.md) — the shared subagents available in this repository.

Codex discovers agents from `agents/*.toml` without registration, and ignores other file types there.

Codex skills live outside this directory, in [`.agents/skills/`](../.agents/README.md), because that is where Codex reads them.
