---
tldr: "Runs rules-checker and rules-fixer over every rule-bearing file until no findings remain."
when_to_use: "Use when governance may have drifted, or before or after a substantial rule change."
---

# Rules Quality Gate

## Purpose

Find and resolve the ways written guidance decays: two documents that contradict each other, one rule stated in three places, a reference to something that no longer exists, and a harness that never received a rule the others did.

## When to Use

Run it on demand — after a large rule change, when a contradiction is suspected, or periodically to check drift. It is deliberately not mandatory: [rules-propagation](rules-propagation.md) integrates one rule correctly on its own, and gating every one-line edit behind a full corpus review would make small corrections expensive enough to skip.

## Prerequisites

A clean working tree, or a change complete enough to review as a whole. Choose a severity level; the gate reuses the levels defined for the [plan quality gate](plan-quality-gate/01-severity-and-modes.md), and strict is the default.

## Steps

1. Establish the corpus; see [scope and corpus](rules-quality-gate/01-scope-and-corpus.md). It spans `repo-governance/`, `AGENTS.md`, `CLAUDE.md`, every harness directory and `SKILL.md`, and `docs/` on the narrow terms that document sets.
2. Run the [Harness Alignment](harness-alignment.md) workflow as a step. It owns the harness inventory, the command and path verification, and the parity comparison, and this gate invokes it rather than restating it.
3. Run `rules-checker` over the corpus. It reports findings by severity against the [finding taxonomy](rules-quality-gate/02-finding-taxonomy.md), citing `file:line`.
4. Stop if no finding meets the chosen level. Otherwise run `rules-fixer` on the findings at or above it; see the [check and fix loop](rules-quality-gate/03-check-fix-loop.md).
5. Re-run `rules-checker`. Repeat until two consecutive runs are clean at that level, or seven cycles have passed.
6. Record the outcome; see the [findings report](rules-quality-gate/04-findings-report.md).

## Verification

```sh
npm run format:check
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
```

Every check passes, and two consecutive `rules-checker` runs report nothing at the chosen level. The parity check matters most after a run: `rules-fixer` may edit harness prompts, and parity is the only automated proof it left the harnesses equal.

## Recovery

A contradiction between two rules is not the fixer's to settle. When one is found, present both texts, their practical effect, and a recommended resolution to the owner, and wait — the same rule [rules-propagation](rules-propagation.md) states, for the same reason.

If the loop reaches seven cycles, stop. Governance that will not converge is usually structured wrong rather than worded wrong, and restructuring is a rule change, not a fix.
