---
tldr: "Divides gate work between rules-checker and rules-fixer and bounds the loop."
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

## rules-fixer

Edits any file in the corpus, including subagent prompts and skills. That authority is deliberate — a rule that drifted into a prompt is where drift does the most damage, and leaving it manual leaves the loop unconverged.

It may:

- replace a duplicated rule with a link to its canonical source
- correct an orphan reference to the current name or path
- add a missing index entry, or missing `tldr` and `when_to_use` frontmatter
- close a gap by adding the missing rule to the harness or document that lacks it

It may not:

- resolve a contradiction by choosing a side
- change what a rule requires, or narrow or widen its scope
- delete a rule to make a finding disappear
- edit a prompt in a way that changes what its role does, rather than what it says

After any edit to a harness directory, `npm run check:harness-parity` must pass before the loop continues. Parity is the only automated proof that the fixer left the harnesses equal.

## Loop Bounds

Two consecutive clean runs end the loop; seven cycles end it with findings open. A contradiction stops the loop immediately and goes to the owner.
