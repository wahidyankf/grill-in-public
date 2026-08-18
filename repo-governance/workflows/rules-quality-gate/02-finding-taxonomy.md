---
tldr: "Lists the five alignment cases and three mechanical checks the gate reports."
when_to_use: "Use when classifying a rules-checker finding or reading a gate report."
---

# Finding Taxonomy

The gate mechanizes the five cases the [Harness Alignment](../harness-alignment.md) workflow defines, so both use one vocabulary, and adds the three checks a reader can verify mechanically.

## The Five Cases

**Equal** — the text matches canonical guidance. Not a finding; recording it is how a clean run proves it looked.

**Contradiction** — two documents differ in requirement, scope, or verification. The most serious case, because whichever one a reader finds first wins, and which one that is depends on where they started. Always CRITICAL or HIGH, and never resolved by the fixer alone.

**Duplication** — one rule stated in two places in words that can drift. Resolve by keeping the canonical statement and replacing the copy with a link. MEDIUM, unless the copies already disagree, which makes it a contradiction.

**Orphan** — a reference to a path, command, workflow, or policy that no longer exists under that name. Renames are the usual cause. HIGH when an instruction file points at it, because an agent will follow it.

**Gap** — a rule one harness or document has and its peers need but lack. HIGH when it changes behavior, MEDIUM when it is operational detail.

## The Mechanical Checks

**Word limit** — every governed document stays within 500 words. `npm run check:governance` is the authority; the gate reports a document approaching the limit so a split is planned rather than forced.

**Index freshness** — every directory's README registers its immediate documents and child directories, per the [documentation index policy](../../documentation-index-policy.md). A missing entry hides work.

**Frontmatter** — every document under `docs/` and `repo-governance/`, except the governance entry index, carries `tldr` and `when_to_use`.

## Severity

The gate reuses the levels and modes defined for the [plan quality gate](../plan-quality-gate/01-severity-and-modes.md). One severity vocabulary across both gates means a report reads the same way whichever produced it.
