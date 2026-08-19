---
tldr: "Indexes the details of the rules-quality-gate workflow."
when_to_use: "Use when looking up the gate's corpus, taxonomy, loop, or report format."
---

# Rules Quality Gate Details

Detail behind the [rules-quality-gate](../rules-quality-gate.md) workflow. Filenames are numbered; the [document naming policy](../../conventions/document-naming-policy.md) says why.

## Contents

- [Scope and Corpus](01-scope-and-corpus.md) — which files the gate reads, and how it treats each kind.
- [Finding Taxonomy](02-finding-taxonomy.md) — the five alignment cases, the three mechanical checks, and the sweep each defect triggers.
- [Check and Fix Loop](03-check-fix-loop.md) — how `rules-checker` and `rules-fixer` divide the work.
- [Fixer Discipline](04-fixer-discipline.md) — the three checks a fixer runs before an edit lands; open it before applying any fix.
- [Findings Report](05-findings-report.md) — what a run records when it finishes.
- [Survival Taxonomy](06-survival-taxonomy.md) — the five reasons a finding survived earlier cycles.
- [Recovery](07-recovery.md) — the loop's bounds, what to do when a contradiction is found or the loop will not converge, and why a run keeps finding things.
