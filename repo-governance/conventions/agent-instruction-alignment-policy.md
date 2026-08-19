---
tldr: "Keeps assistant-specific instruction files deferring to AGENTS.md without contradiction."
when_to_use: "Use when creating, editing, or reviewing CLAUDE.md or any other assistant-specific instruction file."
---

# Agent Instruction Alignment Policy

## Scope

This policy applies to every assistant-specific instruction file in the repository, at any depth, including `CLAUDE.md`, `GEMINI.md`, `COPILOT.md`, `.cursorrules`, and any equivalent file a future tool introduces. It does not apply to `AGENTS.md`, which these files depend on, or to `repo-governance/`.

The [agent harness support policy](agent-harness-support.md) records which harness reads which file, and when a new instruction file is warranted at all. The [agent vocabulary](agent-vocabulary.md) fixes what each of those words means.

## Canonical Source

[`AGENTS.md`](../../AGENTS.md) is the canonical agent instruction file, and it links to the detailed policies in `repo-governance/`. An assistant-specific file is a derivative: it may add only the operational detail its tool needs, such as command invocations, tool-specific workflows, or architecture notes that help that assistant work faster. It must state that `AGENTS.md` is authoritative and link to it.

## Word Limit

`CLAUDE.md` is one of the documents the [document word limit policy](document-word-limit-policy.md) governs; that policy sets the limit, names its scope, and states what to do when a document reaches it. Keep every instruction file equally concise. Badak Mini treats `AGENTS.md` and `CLAUDE.md` as required, so removing one fails the check.

## Required Alignment

An assistant-specific file must not contradict `AGENTS.md` or any document in `repo-governance/`. When guidance would conflict, resolve the conflict at the canonical source using the [Rules Propagation](../workflows/rules-propagation.md) workflow instead of writing a divergent local instruction.

Do not restate a rule that `AGENTS.md` or `repo-governance/` already owns. Link to it instead, so a single edit at the canonical source stays true everywhere. A short pointer that names the rule and links to its home is acceptable when the assistant needs the rule at the point of work; a paraphrase that could drift out of date is not.

## Maintenance

When a change alters `AGENTS.md` or a governance document, review every assistant-specific file in the same change and update the affected references, links, and derived detail, following the [Harness Alignment](../workflows/harness-alignment.md) workflow. A stale derivative is a defect, not a harmless copy.

When a rule in an assistant-specific file turns out to apply to every agent, move it to its canonical home first, then link to it.

## Verification

Confirm that each assistant-specific file names `AGENTS.md` as authoritative, contains no rule that contradicts or silently duplicates canonical guidance, and links only to existing targets. Run the [Harness Alignment](../workflows/harness-alignment.md) workflow for the inventory command, the per-item review, and the required checks.
