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
- [Document Naming Policy](document-naming-policy.md) — how Markdown documents and their child directories are named. Use it when creating, renaming, or splitting a document.
- [Task Tracking Policy](task-tracking-policy.md) — how granular a task list must be and when it must be updated. Use it when starting or reviewing work that takes more than one step.
- [Grilling-With-Options Policy](grilling-with-options-policy.md) — the structured form an agent must use to resolve an open decision with the owner. Use it before asking the owner to decide anything.
- [Grilling Harness Binding](grilling-harness-binding.md) — the question tool each harness uses and the Markdown fallback. Use it when asking a decision question or adding a harness.
- [Plans Organization Policy](plans-organization-policy.md) — how a plan is staged, named, structured, and archived under `plans/`. Use it when creating, executing, or archiving a plan; its detail lives in [`plans-organization-policy/`](plans-organization-policy/README.md).

## Adding a Convention

Add a focused convention here when a standard applies broadly across the repository and is not a foundational principle, executable-development policy, or repeatable procedure. Keep the canonical rule in one document, link to it from concise entry points, and use the [Rules Propagation](../workflows/rules-propagation.md) workflow before changing governance.
