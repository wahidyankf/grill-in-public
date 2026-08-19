# opencode Harness

This directory holds opencode's project configuration for Grind in Public. opencode reads its repository rules from [`AGENTS.md`](../AGENTS.md), falling back to `CLAUDE.md` only when no `AGENTS.md` exists, so no opencode-specific instruction file is needed; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- [`agents/`](agents/README.md) — the shared subagents available in this repository.
- [`plugin/`](plugin/README.md) — session plugins, currently the rule-change notice.
- [`skills/`](skills/README.md) — opencode's skill directory. It holds none: opencode reads the shared ones from `.claude/skills/` and `.agents/skills/`.

Project settings live in [`opencode.json`](../opencode.json) at the repository root, which is where opencode expects them.
