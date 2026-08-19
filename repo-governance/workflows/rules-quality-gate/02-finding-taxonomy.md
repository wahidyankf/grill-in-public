---
tldr: "Lists the five alignment cases, the three mechanical checks, and the sweep each defect triggers."
when_to_use: "Use when classifying a rules-checker finding or reading a gate report."
---

# Finding Taxonomy

The five cases are defined here, and the [Harness Alignment](../harness-alignment.md) workflow classifies every difference it finds with the same five, so the gate and that workflow share one vocabulary rather than two. This document adds the three checks a reader can verify mechanically.

## The Five Cases

**Equal** — the text matches canonical guidance. Not a finding; recording it is how a clean run proves it looked.

**Contradiction** — two documents differ in requirement, scope, or verification. The most serious case, because whichever one a reader finds first wins, and which one that is depends on where they started. Always CRITICAL or HIGH, and never resolved by the fixer alone.

**Duplication** — one rule stated in two places in words that can drift. Resolve by keeping the canonical statement and replacing the copy with a link. MEDIUM, unless the copies already disagree, which makes it a contradiction.

**Orphan** — a reference to a path, command, workflow, or policy that no longer exists under that name. Renames are the usual cause. HIGH when an instruction file points at it, because an agent will follow it.

**Gap** — a rule one harness or document has and its peers need but lack. HIGH when it changes behavior, MEDIUM when it is operational detail.

## The Mechanical Checks

**Word limit** — the [document word limit policy](../../conventions/document-word-limit-policy.md) sets the cap and how a document that reaches it is fixed. `npm run check:governance` is the authority, and it reports only breaches, so the gate measures headroom itself and reports every governed document of 475 words or more.

**Index freshness** — every directory's README registers its immediate documents and child directories, per the [documentation index policy](../../documentation-index-policy.md). A missing entry hides work.

**Frontmatter** — every document under `docs/` and `repo-governance/`, except the governance entry index, carries `tldr` and `when_to_use`.

## Sweeping for a Shape

A defect found once is a shape to sweep for, not an instance to report. Having found one, the checker looks for the same shape everywhere it could repeat — the other gate, the other child documents, the other harnesses, the other subagent roles — and says what it swept and what it found, including nothing. Without that step the gate reports instances and the corpus converges one document at a time.

## Severity

The gate reuses the levels and modes defined for the [plan quality gate](../plan-quality-gate/01-severity-and-modes.md).

Every `rules-checker` prompt states these cases, floors, checks, and the sweep in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.
