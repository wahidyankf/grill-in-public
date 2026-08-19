---
tldr: "Lists the instruction file and project config each supported harness reads."
when_to_use: "Use when configuring a harness or checking which file it reads."
---

# Supported Harnesses

| Harness | Instructions read | Project config |
| --- | --- | --- |
| Claude Code | `CLAUDE.md` only, with no fallback | `.claude/settings.json` |
| Codex | `AGENTS.md` | `.codex/config.toml` |
| opencode | `AGENTS.md`, falling back to `CLAUDE.md` | `opencode.json` |

Where each harness loads its subagents, skills, and commands belongs to the [harness capability parity policy](../harness-capability-parity-policy.md).
