---
tldr: "Specifies what a rules quality gate run records and where."
when_to_use: "Use when closing out a gate run or reading its history."
---

# Findings Report

A gate whose runs leave no trace cannot be audited, and an unauditable gate quietly stops running.

## Where It Goes

Governance has no per-run home the way a plan does, so a run appends one line to `local-tmp/gate-history/rules-quality-gate.md`. That path is untracked, which is what lets the log keep every line: a governed document is capped by the [document word limit policy](../../conventions/document-word-limit-policy.md) and an append-only history is not, so the two cannot share a file. [`repo-governance/README.md`](../../README.md) names the location and holds nothing else.

```markdown
## Gate History

- 2026-08-18 — strict — 2 cycles — pass
- 2026-09-02 — strict — 7 cycles — fail (1 contradiction open, AGENTS.md vs testing-policy.md)
```

Date, level, cycles, status, and for anything short of a pass, the open findings in one clause each. A settled run has none, so its clause names the fixes no cycle re-read instead; that is what the next run verifies first, and this line is what it reads before opening anything. Keep every line: a corpus that needs seven cycles twice in a row is telling you something a single status cannot.

## Statuses

The [run statuses](08-run-statuses.md) document defines the four a run can close with, and which of them obliges the next run to verify the last fixes.

## Open Findings

A finding the run acted on and left unresolved is written into the affected document itself, not only into the untracked history line. A finding recorded where the reader of the rule will never look has not been recorded.

A contradiction is the case of that rule which admits no exception. One that reaches the owner and is not resolved in the session is written into both affected documents. A reader of a rule must be able to see that it is disputed; a dispute recorded only in a log is a dispute the next reader will not find.

## Why Each Finding Survived

The [survival taxonomy](06-survival-taxonomy.md) owns the line each finding carries about why earlier cycles did not catch it.

## Unplanned Changes

Record every document the run created, and every pass that touched several documents. A plan may be created for that work only when the owner explicitly requests one; the [plans organization policy](../../conventions/plans-organization-policy.md) owns that authorization boundary. `rules-fixer` does its work without a plan, so this report is what makes it auditable.

## What Not to Record

Do not list findings the fixer resolved. Once the guidance is right, the history of it being wrong is noise that buries what is still open.
