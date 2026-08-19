---
tldr: "Integrates changed repository rules without duplication or contradictions."
when_to_use: "Use before adding, moving, changing, or removing repository rules or agent guidance."
---

# Rules Propagation

## Purpose

Integrate a new or changed rule into the correct governance location without duplication, ambiguity, or contradictions. Rules live in `AGENTS.md`, `repo-governance/`, and any harness or skill instruction file.

## When to Use

Use it when introducing, changing, moving, or removing a rule for contributors, agents, skills, validation, or workflows.

Run it directly for a correction, a clarification, or a rule confined to one document. Substantial rule work, a new policy area or a change landing across several documents, is planned first; the [plans organization policy](../conventions/plans-organization-policy.md) owns that boundary.

## Automatic Triggers

The repository announces this workflow rather than relying on memory. The [rule change trigger policy](../development/rule-change-trigger-policy.md) owns the rule paths, the hooks that watch them, and when [Harness Alignment](harness-alignment.md) is announced alongside this workflow.

An announcement is not the work; carrying out the steps below is still yours.

## Prerequisites

State the rule in one sentence: its scope, trigger, and required behavior. Keep the decision that justifies it.

## Steps

1. Inventory applicable guidance before editing; see [inventory](rules-propagation/01-inventory.md).

2. Choose one canonical home for the rule; see [canonical home](rules-propagation/02-canonical-home.md).

3. Merge equivalent, overlapping, or inverse rules into that source; see [conflict resolution](rules-propagation/03-conflict-resolution.md).

4. Resolve contradictions before writing, and never settle one alone; same document.

5. Integrate the approved rule using direct, testable language. Link from a concise document to its detailed source instead of copying the same rule. When creating or editing Markdown, follow the [Markdown style policy](../conventions/markdown-style-policy.md).

## Verification

Confirm the rule has one canonical source and accurate references, and that no contradictory guidance remains. When the change touched many documents, run the [rules-quality-gate](rules-quality-gate.md) workflow to check the corpus rather than the diff. Run:

```sh
npm run format:check
npm run check:governance
```

## Recovery

If scope or precedence stays unclear, leave the rules unchanged and ask. For an overlong document, follow the [document word limit policy](../conventions/document-word-limit-policy.md); [progressive disclosure](../principles/progressive-disclosure.md) states when a document is split.
