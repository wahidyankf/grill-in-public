---
tldr: "Bounds the loop, and states what to do when a contradiction stops the gate or the loop will not converge."
when_to_use: "Use when a gate run finds a contradiction, or reaches the cycle count that ends the loop."
---

# Recovery

## Loop Bounds

Two consecutive clean runs end the loop; seven cycles end it with findings open. Cycle five raises a warning: a corpus still finding new problems that late is usually structurally wrong rather than imprecise. A contradiction stops the loop immediately and goes to the owner.

## A Contradiction

A contradiction between two rules is not the fixer's to settle. When one is found, present both texts, their practical effect, and a recommended resolution to the owner, and wait. [Conflict resolution](../rules-propagation/03-conflict-resolution.md) requires that wait for a substantive conflict; the gate waits on every contradiction it finds.

## A Loop That Will Not Converge

If the loop reaches seven cycles, stop. Governance that will not converge is usually structured wrong rather than worded wrong, and restructuring is a rule change, not a fix.

Stop earlier than the bound when the findings stop being about the corpus. A run that is finding defects in the gate's own scope, or in what a fixer is authorized to do, is asking a question no cycle is entitled to answer, and another cycle will find the same shape and still not be able to close it. Take it to the owner instead.

## Why a Run Keeps Finding Things

Two causes, both observed, both worth knowing before blaming the corpus.

A repair pass that touches many documents reliably produces findings in the next cycle, and they are in the pass's own new text rather than in what it repaired. Prefer the smallest edit that fully resolves each finding, and resist improving anything the findings did not name.

A prompt edited during a run does not reach the agents that run spawns: subagent definitions are read once, when the session starts. So a fix to a subagent prompt cannot be verified inside the run that makes it, and a checker may report against wording the repository no longer has. Record such a fix as unverified, and confirm it in the next run.
