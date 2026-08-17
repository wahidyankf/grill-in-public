# Codex Harness

This directory holds the project-scoped Codex configuration for SWE Grilling. Codex reads its repository rules from [`AGENTS.md`](../AGENTS.md), so no Codex-specific instruction file exists; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- `config.toml` — project settings. Codex loads this layer only for a trusted project, and it takes precedence over the user config.
- [`agents/`](agents/README.md) — the shared subagents available in this repository.

Codex discovers agents from `agents/*.toml` without registration, and ignores other file types there.
