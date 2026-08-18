---
tldr: "Records the supported agent harnesses and where each one reads instructions, config, and subagents."
when_to_use: "Use when adding, configuring, or changing support for an agent harness such as Claude Code, Codex, or opencode."
---

# Agent Harness Support Policy

## Supported Harnesses

This repository supports three harnesses, in the sense the [agent vocabulary](agent-vocabulary.md) defines: the tool that runs the model. Each reads a different instruction file, which is why `AGENTS.md` is canonical and `CLAUDE.md` exists as its derivative.

| Harness | Instructions read | Project config |
| --- | --- | --- |
| Claude Code | `CLAUDE.md` only, with no fallback | `.claude/settings.json` |
| Codex | `AGENTS.md` | `.codex/config.toml` |
| opencode | `AGENTS.md`, falling back to `CLAUDE.md` | `opencode.json` |

Where each harness loads its subagents, skills, and commands belongs to the [harness capability parity policy](harness-capability-parity-policy.md).

## Rules

Do not create a `CODEX.md` or `OPENCODE.md`. Both tools read `AGENTS.md` already, and a third instruction file would add a copy to keep aligned for no gain. Add an instruction file only when its harness cannot read `AGENTS.md`, and then follow the [agent instruction alignment policy](agent-instruction-alignment-policy.md).

Keep the shared subagents, skills, and commands at parity across harnesses, as the [harness capability parity policy](harness-capability-parity-policy.md) requires. That policy owns the capability directories and the check that enforces them.

Keep each project config to settings that the tool documents and that the repository actually needs. A config file is not a place for rules; rules belong in `AGENTS.md` or `repo-governance/`. Do not re-list an auto-discovered file as an extra instruction; `opencode.json` therefore pins only its schema.

Write every agent description as two things: a short statement of what the agent does, then a short sentence saying when it should be used. Harnesses route work by description, so an agent that omits its trigger will be picked at the wrong time or never. Keep the wording equivalent across formats.

Index a harness directory with a `README.md`, as the [documentation index policy](../documentation-index-policy.md) requires, and update it in the same change that adds, renames, or removes a file there. Verified behavior for an index inside a registering directory:

| Directory | Index behavior |
| --- | --- |
| `.claude/agents/` | Ignored without agent frontmatter |
| `.claude/commands/` | Becomes `/README`; no README, index from parent |
| `.codex/agents/` | Ignored; only `*.toml` is read |
| `.opencode/agents/` | Becomes an agent unless frontmatter sets `disable: true` |

A permission control that one harness lacks is recorded where it is missing. Codex has no per-agent shell switch, so its read-only explorer relies on `sandbox_mode` plus an explicit instruction; that divergence is noted in the agent file rather than left implicit.

## Verification

Run the [Align Agent Harnesses](../workflows/align-agent-harnesses.md) workflow after changing any instruction file, harness config, or subagent. Test a new tool's discovery behavior yourself before relying on it, because support changes between releases.
