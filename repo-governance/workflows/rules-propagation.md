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

The repository announces this workflow rather than relying on memory: pre-commit names it when a staged change touches a rule path, and each harness announces it while the edit is being written. A harness path announces [Harness Alignment](harness-alignment.md) as well. The [rule change trigger policy](../development/rule-change-trigger-policy.md) owns those paths and mechanisms.

An announcement is not the work; carrying out the steps below is still yours.

## Prerequisites

State the rule in one sentence: its scope, trigger, and required behavior. Keep the decision that justifies it.

## Steps

1. Inventory applicable guidance before editing. Start with `AGENTS.md`, `repo-governance/`, and its `principles/`; also search for instruction and skill files:

   ```sh
   rg --files -g 'AGENTS.md' -g 'SKILL.md' -g 'CLAUDE.md' -g 'GEMINI.md' \
     -g 'COPILOT.md' -g '!node_modules'
   ```

2. Choose one canonical home:
   - Put universal, short requirements in root `AGENTS.md`.
   - Put detailed or conditional policy in `repo-governance/`.
   - Put stable, cross-cutting standards in `repo-governance/conventions/`.
   - Put repeatable procedures in `repo-governance/workflows/`.
   - Put directory-specific rules in the nearest scoped `AGENTS.md`.
   - Put capability-specific guidance in the relevant `SKILL.md`.
   - Create and categorize a focused document or subdirectory in `repo-governance/` when no existing location suits, and do not create empty categories. The [documentation index policy](../documentation-index-policy.md) owns the README, index, and frontmatter requirements that follow.

3. Search the inventory for equivalent, overlapping, or inverse rules. Merge them into the canonical source, update references, and remove redundant wording only when its meaning is fully preserved.

4. Resolve contradictions before writing. Never silently choose between rules that differ in requirement, scope, or verification. Present the conflicting text, its practical effect, and a recommended resolution to the owner, and wait when the conflict is substantive.

5. Integrate the approved rule using direct, testable language. Link from a concise document to its detailed source instead of copying the same rule. When creating or editing Markdown, follow the [Markdown style policy](../conventions/markdown-style-policy.md).

## Verification

Confirm the rule has one canonical source and accurate references, and that no contradictory guidance remains. When the change touched many documents, run the [rules-quality-gate](rules-quality-gate.md) workflow to check the corpus rather than the diff. Run:

```sh
npm run format:check
npm run check:governance
```

## Recovery

If scope or precedence stays unclear, leave the rules unchanged and ask. Split an overlong document into focused files to preserve progressive disclosure.
