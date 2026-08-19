---
tldr: "Runs rules-checker and rules-fixer over every rule-bearing file until no findings remain."
when_to_use: "Use when governance may have drifted, or before or after a substantial rule change."
---

# Rules Quality Gate

## Purpose

Find and resolve the ways written guidance decays: two documents that contradict each other, one rule stated in three places, a reference to something that no longer exists, and a harness that never received a rule the others did.

## When to Use

Run it on demand — after a large rule change, when a contradiction is suspected, or periodically to check drift. It is deliberately not mandatory: [rules-propagation](rules-propagation.md) integrates one rule correctly on its own, and gating every one-line edit behind a full corpus review would make small corrections expensive enough to skip.

Running the gate needs no plan, but resolving a finding can. Substantial rule work, a new policy area or a change landing across several documents, is planned first; the [plans organization policy](../conventions/plans-organization-policy.md) owns that boundary and the `rules-fixer` exemption.

## Prerequisites

A clean working tree, or a change complete enough to review as a whole. Choose a severity level; the gate reuses the levels defined for the [plan quality gate](plan-quality-gate/01-severity-and-modes.md), and strict is the default.

## Steps

1. Establish the corpus; see [scope and corpus](rules-quality-gate/01-scope-and-corpus.md).
2. Run the [Harness Alignment](harness-alignment.md) workflow as a step. It owns the harness inventory, the command and path verification, and the parity comparison, and this gate invokes it rather than restating it.
3. Run `rules-checker` over the corpus. It reports findings by severity against the [finding taxonomy](rules-quality-gate/02-finding-taxonomy.md), citing `file:line`.
4. Run `rules-fixer` on the findings at or above the chosen level; see the [check and fix loop](rules-quality-gate/03-check-fix-loop.md) and the [fixer discipline](rules-quality-gate/04-fixer-discipline.md) it runs before each edit lands. A run that found nothing skips this step but does not end the gate: one clean run can mean a checker that stopped early rather than a corpus that is sound.
5. Re-run `rules-checker`. Repeat until two consecutive runs are clean at that level, or seven cycles have passed.
6. Record the outcome; see the [findings report](rules-quality-gate/05-findings-report.md).

## Verification

```sh
npm run format:check
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
```

Every check passes, and two consecutive `rules-checker` runs report nothing at the chosen level. The parity check matters most after a run: `rules-fixer` may edit harness prompts, and parity is the only automated proof it left the harnesses equal.

## Recovery

See [recovery](rules-quality-gate/07-recovery.md) for what to do when a contradiction is found, and when the loop reaches seven cycles.
