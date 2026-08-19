---
tldr: "Three checks a fixer runs before an edit lands, drawn from the repairs this gate has broken."
when_to_use: "Use before applying any gate fix, in either quality gate."
---

# Fixer Discipline

These are not intentions a fixer holds; they are steps it performs. Each one names a failure the gate's own repairs have committed more than once, and each is far cheaper to run than to find later. Both the [rules gate](../rules-quality-gate.md) and the [plan gate](../plan-quality-gate.md) bind their fixers to them, the way both gates share [one severity vocabulary](../plan-quality-gate/01-severity-and-modes.md).

## No Clause Dies Without a Home

Before shortening a document, list every clause you are removing, and for each one name the document that still states its requirement. Search for it; do not recall it. A requirement with no surviving home has been deleted, whatever the edit was called.

This is the failure a word limit manufactures: the pressure to compress is highest on exactly the document a finding is already open against.

## Diff Before Claiming Equality

A sentence asserting that two texts match — "verbatim", "in full", "states these lists" — is a verification claim. So is a sentence saying what a subagent checks, what a command enforces, or that one document's rule matches another's. Write either only after verifying it in the same session: diff the two texts, read the prompt, run the command, or extract the other document's rule and compare it. Confirm each extract is non-empty first, because an empty comparison passes silently.

A false claim of either kind is worse than no claim. An equality claim tells the next editor that synchronization already holds, so they change one side and stop. A claim about behavior does worse: the reader stops checking the thing themselves, because the sentence told them the gate already does.

## Fix the Sibling in the Same Pass

This guidance is built in pairs: two quality gates, their child documents, six subagent roles across three harnesses, and each subagent prompt against the workflow it implements. A defect found in one member of a pair is a defect suspected in the other.

Test the peer in the same pass, then either fix it or record why the shape does not apply there. Repairing one half of a pair silently makes the pair unequal, which is a new finding created by the repair.

Every `rules-fixer` and `plan-fixer` prompt carries these three checks in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all six copies.
