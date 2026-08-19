---
tldr: "Specifies the five documents every plan folder contains and what each owns."
when_to_use: "Use when scaffolding a plan folder or deciding which file a section belongs in."
---

# Five-Document Structure

Every plan uses five documents. Each owns one concern, so a reader looking for the technical approach never has to skim the rationale.

```text
plans/<stage>/<identifier>/
+-- README.md        overview, scope, links to the rest
+-- brd.md           why this matters
+-- prd.md           what it must do, in user stories and Gherkin
+-- tech-docs.md     how it is built
+-- delivery.md      the phased checklist that drives execution
+-- learnings.md     (transient) running log, drained before archival
+-- evidence/        (optional) command output and artifacts referenced from delivery.md
```

## What Each File Owns

**`README.md`** — context, scope with affected projects named explicitly, a summary of the approach, and links to the other four. It is the first file opened and the first file `plan-checker` reads for scope.

**`brd.md`** — the business rationale: why the work is worth doing, who it affects, what success means, business-level non-goals, and the risks. In a personal repository the "business" is the owner's own goals, so write real reasoning and label a judgment call as one. Never invent a metric to fill a heading.

**`prd.md`** — product requirements: user stories in `As a … I want … So that …` form and acceptance criteria in Gherkin, per the [specs policy](../../development/specs-policy.md). In-scope and out-of-scope features live here.

**`tech-docs.md`** — architecture, design decisions with their rationale, the annotated file-impact tree, dependencies, risks, and the rollback path. No checklist.

**`delivery.md`** — the phased, ticked checklist that execution reads and `plan-checker` verifies; see [delivery checklists](delivery-checklists.md).

**`learnings.md`** — the transient running log described in [knowledge capture](knowledge-capture.md).

## No Single-File Exception

Five documents is the only shape. A plan too small for five documents is a plan for work that did not need planning: do it, and track it in the harness task list instead.

Every `plan-checker` prompt names these documents and requires all five in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.
