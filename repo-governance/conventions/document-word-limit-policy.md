---
tldr: "Sets the word limit every governed document lives under and how a document that reaches it is fixed."
when_to_use: "Use when a governed document nears or exceeds the word limit, or when deciding how to shorten one."
---

# Document Word Limit Policy

## Scope

`npm run check:governance` holds `AGENTS.md`, `CLAUDE.md`, every `repo-governance/` document, and every harness directory README to 500 words. Agent and command definitions are prompts, not indexes, and stay unmeasured. This policy governs every document that check measures, and it binds anyone editing one of them, not only a [rules quality gate](../workflows/rules-quality-gate.md) run.

## The Limit

Every governed document stays within 500 words. A document of 475 words or more sits in the headroom band: 25 or fewer words are left under the cap.

The point is to plan a split rather than force one: a document with no room left is where a later fix reaches for compression, and compression is how a clause gets deleted.

## Closing the Finding

The finding is MEDIUM, and only relocation closes it: text moved into a linked document, whole. Shortening the document back under the band does not, because that is the edit the finding exists to prevent. An instruction file that cannot be split still relocates, into a focused document it links to.

An index that reaches the limit is not too wordy; its directory has too many peers. Group them: create a subdirectory, move the related documents into it, give it a README, and register that child from the parent, which then carries one line instead of many. The [documentation index policy](../documentation-index-policy.md) owns what an index must never do to fit.

Every `rules-checker` prompt states this scope, this limit, and this remedy in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

Severity levels are the ones defined for the [plan quality gate](../workflows/plan-quality-gate/01-severity-and-modes.md). Relocation is one of the edits the [check and fix loop](../workflows/rules-quality-gate/03-check-fix-loop.md) permits a fixer, and the [fixer discipline](../workflows/rules-quality-gate/04-fixer-discipline.md) is what keeps a relocation from losing a clause on the way.

## Verification

```sh
npm run check:governance
```

[Workspace commands](../development/workspace-commands.md#repository-checks) lists this check beside the repository's others. The [finding taxonomy](../workflows/rules-quality-gate/02-finding-taxonomy.md) records how the gate measures headroom and reports a document that has reached the band without breaching the limit.
