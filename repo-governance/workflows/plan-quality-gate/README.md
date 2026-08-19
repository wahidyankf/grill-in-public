---
tldr: "Indexes the details of the plan-quality-gate workflow."
when_to_use: "Use when looking up a severity level, the loop rules, or the report format."
---

# Plan Quality Gate Details

Detail behind the [plan-quality-gate](../plan-quality-gate.md) workflow. Filenames are numbered; the [document naming policy](../../conventions/document-naming-policy.md) says why.

## Contents

- [Severity and Modes](01-severity-and-modes.md) — what each severity means and which level to run.
- [Check and Fix Loop](02-check-fix-loop.md) — how `plan-checker` and `plan-fixer` divide the work.
- [Role Separation](03-role-separation.md) — why checking and fixing are two subagents rather than one.
- [Findings Report](04-findings-report.md) — what the gate records when it finishes.
