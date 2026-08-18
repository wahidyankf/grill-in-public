---
tldr: "Indexes stable, cross-cutting standards for repository work."
when_to_use: "Use when creating, changing, or reviewing work covered by a shared convention."
---

# Governance Conventions

This directory contains stable, cross-cutting standards that make repository work consistent. Conventions implement the foundational principles without replacing focused development policies or repeatable workflows.

## Available Conventions

- [Agent Harness Support Policy](agent-harness-support.md) — which harnesses are supported and where each reads its instructions and config. Use it when adding or configuring a harness.
- [Harness Capability Parity Policy](harness-capability-parity-policy.md) — the subagents, skills, and commands every harness must expose alike, and where each lives. Use it when adding, renaming, or removing one.
- [Agent Vocabulary](agent-vocabulary.md) — what harness, agent, instruction file, and subagent mean here. Use it when writing or reviewing any text about agents.
- [Agent Instruction Alignment Policy](agent-instruction-alignment-policy.md) — how assistant-specific instruction files defer to `AGENTS.md`. Use it when creating, editing, or reviewing `CLAUDE.md` or a similar file.
- [Markdown Style Policy](markdown-style-policy.md) — source formatting for every repository Markdown file. Use it when creating, editing, reviewing, or formatting Markdown.
- [Task Tracking Policy](task-tracking-policy.md) — how granular a task list must be and when it must be updated. Use it when starting or reviewing work that takes more than one step.

## Adding a Convention

Add a focused convention here when a standard applies broadly across the repository and is not a foundational principle, executable-development policy, or repeatable procedure. Keep the canonical rule in one document, link to it from concise entry points, and use the [Propagate Rules](../workflows/propagate-rules.md) workflow before changing governance.
