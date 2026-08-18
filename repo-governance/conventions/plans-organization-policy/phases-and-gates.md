---
tldr: "Requires every phase to end at a passing gate with a Pause Safety note."
when_to_use: "Use when structuring a delivery.md into phases or writing a phase gate."
---

# Phases and Gates

A phase is a natural pause. At every phase boundary the working tree is coherent: it builds, its tests pass, nothing is half-applied, and no known-red state carries forward. A reader must be able to stop after any phase and find the repository usable.

## The Gate

Every phase ends with a `### Phase N Gate` — a must-pass checklist naming exact commands and observable criteria. Phase N+1 does not begin while any item in phase N's gate fails. Gate items carry executor tags like any other checkbox; a `[HUMAN]` gate makes the boundary a hand-off.

## Pause Safety

Immediately after each gate, a short Pause Safety note states the coherent state the phase reached and the single command that re-verifies it. This makes the pause property auditable instead of assumed.

## Template

```markdown
## Phase N: <name>

- [ ] [AI] <action> — acceptance: <observable outcome>

### Phase N Gate

> Every check below passes before Phase N+1 begins. A failure is fixed inside Phase N.

- [ ] [AI] `<command>` — <acceptance>

> **Pause Safety**: <what is coherent now, what changed and what did not>. Safe to stop. Resume with `<command>`.
```

## Phase 0

Phase 0 is environment setup and baseline. It records a clean starting state — dependencies installed, gates green before any change — so a later failure can be attributed to the work rather than to the machine. Its gate is that recorded baseline, and it commits nothing but the plan itself.

## Ordering

Order phases so each builds on a green predecessor. Delivery to `main` happens at phase boundaries: the gate passes, then the phase's work is committed and pushed. A phase whose gate cannot be expressed as a command is a phase that has not been thought through.
