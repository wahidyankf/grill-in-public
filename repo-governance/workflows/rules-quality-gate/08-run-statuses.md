---
tldr: "Defines the four statuses a rules quality gate run can close with."
when_to_use: "Use when closing out a gate run, or reading what an earlier run's status meant."
---

# Run Statuses

A run closes with one of these, and the [findings report](05-findings-report.md) records it in the run's line.

**pass** — two consecutive clean runs at the chosen level.

**settled** — the loop ended with nothing open, but the last fixes were applied after the final check, so no cycle has read them. The next run starts by verifying those fixes, because a fix no checker has seen is a claim rather than a result.

**partial** — the loop ended with findings open, none of them a contradiction. Each open finding is listed with its case and location.

**fail** — a contradiction remains unresolved, or the checker could not read the corpus. A failing gate is not a reason to stop working; it is a reason not to claim the guidance is coherent.
