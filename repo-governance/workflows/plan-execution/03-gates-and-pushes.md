---
tldr: "Defines how a phase gate is passed and how the phase is delivered to main."
when_to_use: "Use at the end of every delivery phase."
---

# Gates and Pushes

A phase ends at its gate. The gate is the only thing that decides whether the next phase may start.

## Passing the Gate

Run every gate item as written, in order, and record the observable result. A gate item is not satisfied by a command that was run earlier in the phase: the gate exists to prove the phase's combined state, not each item separately.

If a gate item fails, fix it inside this phase. Do not start the next phase, and do not mark the gate passed with an exception noted. The [phase and gate rules](../../conventions/plans-organization-policy/phases-and-gates.md) hold the shape.

## Delivering to main

This repository delivers directly to `main`. Once the gate passes:

1. Confirm the working tree is coherent — it builds and its tests pass, which the gate has just shown.
2. Stage the phase's work, including the ticked checkboxes and any `learnings.md` entries.
3. Commit with a Conventional Commits message naming what the phase did, per the [commit hook policy](../../development/commit-hook-policy.md). Split unrelated work into separate commits.
4. Push to `origin main`. The [workspace commands](../../development/workspace-commands.md#hooks) reference lists what pre-push runs and when.

Phase 0 commits only the plan itself, since it changes nothing else.

## When a Push Fails

A rejected push means the remote moved. Pull, re-run the phase gate against the merged state, and push again. Do not force, and do not bypass the hooks: `--no-verify` needs the owner's explicit approval, which is the point of the rule.

## Why Every Phase

Pushing at each gate keeps the remote a series of coherent states rather than one large landing. It also means an interrupted plan leaves finished phases delivered rather than stranded in a local tree.
