---
tldr: "Records the supported agent harnesses and where each one reads instructions, config, and subagents."
when_to_use: "Use when adding, configuring, or changing support for an agent harness such as Claude Code, Codex, or opencode."
---

# Agent Harness Support Policy

## Supported Harnesses

This repository supports three harnesses. Each reads a different instruction file, which is why `AGENTS.md` is canonical and `CLAUDE.md` exists as its derivative.

| Harness | Instructions read | Project config | Subagents |
| --- | --- | --- | --- |
| Claude Code | `CLAUDE.md` only, with no fallback | `.claude/settings.json` | `.claude/agents/*.md` |
| Codex | `AGENTS.md` | `.codex/config.toml` | `.codex/agents/*.toml` |
| opencode | `AGENTS.md`, falling back to `CLAUDE.md` | `opencode.json` | `.opencode/agents/*.md` |

## Rules

Do not create a `CODEX.md` or `OPENCODE.md`. Both tools read `AGENTS.md` already, and a third harness file would add a copy to keep aligned for no gain. Add a harness file only when its tool cannot read `AGENTS.md`, and then follow the [agent instruction alignment policy](agent-instruction-alignment-policy.md).

Keep the shared subagents at parity. The same role must exist in every harness that supports subagents, with the same name, the same purpose, and the same permission posture. Their instructions may be reworded to fit each format, but they must not diverge in what the agent is allowed to do.

Keep each project config to settings that the tool documents and that the repository actually needs. A config file is not a place for rules; rules belong in `AGENTS.md` or `repo-governance/`. Do not list an auto-discovered file again as an extra instruction, which is why `opencode.json` pins only its schema: opencode already reads `AGENTS.md` and `.opencode/agents/`.

A permission control that one harness lacks is recorded where it is missing. Codex has no per-agent shell switch, so its read-only explorer relies on `sandbox_mode` plus an explicit instruction; that divergence is noted in the agent file rather than left implicit.

## Verification

Run the [Align Agent Harnesses](../workflows/align-agent-harnesses.md) workflow after changing any harness file, config, or subagent. Confirm a new tool's discovery behavior against its own documentation before relying on it here, because instruction-file support changes between releases.
