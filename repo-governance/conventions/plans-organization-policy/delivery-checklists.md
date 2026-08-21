---
tldr: "Requires one action per checkbox, execution-grade detail, and an executor tag."
when_to_use: "Use when writing or reviewing any checkbox in a plan's delivery.md."
---

# Delivery Checklists

`delivery.md` is executed by an agent reading it literally. Every checkbox must therefore be unambiguous on its own, without reconstructing intent from the rest of the plan.

## One Action per Checkbox

One checkbox equals one concrete, independently verifiable action. The test: can this item be verified done without completing anything else on the list? If not, split it. Multi-step work hidden behind one checkbox makes progress invisible and makes "done" arguable.

## Execution-Grade Clarity

Each checkbox carries every element that applies:

- **The file path**, exactly, when it is known. When a new file's location is implementation-dependent, give the parent directory, the naming pattern, and a sibling to imitate.
- **The command**, verbatim — `npx nx run badakmini-cli:test:quick`, not "run the tests".
- **An acceptance criterion** stating the observable outcome that proves it done. No bare "implement", "set up", or "configure".
- **One Gherkin scenario** per behavior cycle, inlined verbatim, as the [TDD policy](../../development/tdd-policy.md) requires.

Bad, because it names no file, no command, and no observable outcome:

```markdown
- [ ] Add caching
```

Good:

```markdown
- [ ] [AI] Edit `apps/badakmini-cli/internal/rulechange/detect.go`: preserve one rule-change path after normalization. Verify with `npx nx run badakmini-cli:test:quick` — the suite exits 0.
```

## Executor Tags

Every checkbox states who can execute it. The tags are `[AI]`, `[HUMAN]`, and `[AI+HUMAN]`, matching `ose-public` so a migrated plan needs no translation.

- **`[AI]`** — an agent can fully perform it. Write the tag out even here: an untagged checkbox is a defect, not an `[AI]` one.
- **`[HUMAN]`** — only the owner: physical actions, out-of-band approvals, anything needing real credentials an agent must not hold, and any step the owner is doing by hand to learn it.
- **`[AI+HUMAN]`** — an agent prepares, the owner approves or performs the irreversible step.

Tag toward `[AI]`. Git-mechanical steps such as committing, pushing, or moving a plan folder are `[AI]` unless a specific reason says otherwise. A `delivery.md` opens with a one-line legend naming the three tags, so a reader meets them before the first checkbox.

Every `plan-checker` prompt states one action per checkbox, execution-grade clarity, and the executor tags with their legend in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.
