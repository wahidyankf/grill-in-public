---
tldr: "Specifies what the gate records in the plan when it finishes."
when_to_use: "Use when closing out a gate run or reading a plan's gate history."
---

# Findings Report

A gate run that leaves no trace cannot be audited, and an unaudited gate becomes a gate that quietly stops running.

## What to Record

Append to the plan's `README.md`, under a `## Quality Gate` heading:

```markdown
## Quality Gate

- 2026-08-18 — strict — 3 cycles — pass (0 findings on two consecutive runs)
- 2026-08-19 — strict — 7 cycles — partial (2 MEDIUM open, see below)
```

State the date, the level, the cycles run, and the status. Keep every run: the history shows whether a plan needed one pass or five, which is the clearest signal that its approach was unclear.

## Statuses

**pass** — two consecutive clean runs at the chosen level.

**partial** — the loop ended at seven cycles with findings open. List each open finding with its severity and location. A partial plan may still be executed if the owner accepts the findings, and that acceptance is written down.

**fail** — the checker could not run, or a CRITICAL finding remains. A failing plan is not executed.

## Open Findings

Open findings live in the plan, not in a separate report file. A finding recorded somewhere the executor will not look has not been recorded.

## What Not to Record

Do not record fixed findings individually. The history of what was wrong in a draft is noise once it is right, and it buries the open findings that actually need attention.
