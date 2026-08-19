---
tldr: "Lists the five reasons a finding survived earlier cycles, one of which every finding carries."
when_to_use: "Use when writing a finding's survival line, or reading a run's set of them."
---

# Survival Taxonomy

Every finding carries one line saying why earlier cycles did not catch it: the file was never read on that axis, the text was written by a previous cycle's own fix, it drifted after an earlier fix, it needed a comparison nobody had made, or it sat at the edge of the corpus. Read as a set across a run, those lines say whether the corpus is decaying or the gate is.

Every `rules-checker` prompt states these five reasons in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

The [findings report](05-findings-report.md) records these lines with the rest of a run's outcome.
