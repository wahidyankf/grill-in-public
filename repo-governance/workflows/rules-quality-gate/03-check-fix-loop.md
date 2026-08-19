---
tldr: "Divides gate work between rules-checker and rules-fixer, and links the loop's bounds."
when_to_use: "Use when running the gate or reviewing what each subagent may do."
---

# Check and Fix Loop

```text
harness-alignment --> rules-checker --> findings --> rules-fixer --> rules-checker
                           ^                                             |
                           +------------- up to 7 cycles ----------------+
                                    clean twice --> pass
```

## rules-checker

Reads the corpus and reports findings with a case, a severity, a `file:line` citation, and the canonical source the finding is measured against. It edits nothing.

It reports a contradiction as a pair: both texts, and what each would make a reader do. A contradiction reported as a single quotation is unactionable, because the reader cannot see what it conflicts with.

Every `rules-checker` prompt states this reporting rule and this pairing rule in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

Having found a defect, the checker sweeps the corpus for the same shape; see [sweeping for a shape](02-finding-taxonomy.md#sweeping-for-a-shape).

## rules-fixer

Edits any file in the corpus, including subagent prompts and skills. That authority is deliberate — a rule that drifted into a prompt is where drift does the most damage, and leaving it manual leaves the loop unconverged.

It may:

- replace a duplicated rule with a link to its canonical source
- correct an orphan reference to the current name or path
- add a missing README index entry, or missing `tldr` and `when_to_use` frontmatter
- close a gap by adding the missing rule to the harness or document that lacks it, and only there
- relocate text whole into a linked document, creating it if needed, without shortening, merging, or rewording what moves, leaving every requirement stated exactly once, and recording in the run's report every document created and every pass that touched several documents

It may not:

- resolve a contradiction by choosing a side, ever
- change what a rule requires, or narrow or widen its scope
- delete a rule, or weaken its verification, to make a finding disappear
- change what a subagent's role does, as opposed to how its instruction is worded

After any edit to a harness directory, `npm run check:harness-parity` must pass before the loop continues. Parity is the only automated proof that the fixer left the harnesses equal.

The `rules-fixer` prompt states these two lists verbatim and carries the same parity rule in the imperative, because a subagent prompt has to stand alone. They are one rule in two places by necessity: change them in the same edit, in all three harness copies, or the fixer and the workflow start authorizing different things.

The loop's bounds are stated with [recovery](07-recovery.md), which is where a run that reaches them goes next.
