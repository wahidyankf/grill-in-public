---
tldr: "Defines the four finding severities and the four gate levels."
when_to_use: "Use when choosing a gate level or classifying a finding."
---

# Severity and Modes

One severity vocabulary across both gates means a report reads the same way whichever produced it.

## Severities

**CRITICAL** — the plan would do damage or cannot be executed at all. A secret in a document, a step that deletes something the plan never restores, a missing document, or a phase whose gate cannot pass.

**HIGH** — the plan would be executed wrongly. A checkbox with no acceptance criterion, a checkbox hiding several actions, a behavior cycle binding more than one scenario or none, a phase without a gate, or a command that does not exist in any `project.json` or `package.json`.

**MEDIUM** — the plan would be executed correctly but slowly or with guesswork. A missing file path where one is knowable, a gate lacking a concrete command, a missing Pause Safety note, or a file-impact tree that omits a path the checklist touches.

**LOW** — the plan is sound but rough. Wording, ordering, a missing diagram for a concern that would benefit, or an index entry that reads poorly.

## Levels

| Level | Acts on | Use when |
| --- | --- | --- |
| lax | CRITICAL | An early draft still being shaped |
| normal | CRITICAL, HIGH | A plan under active revision |
| strict | CRITICAL, HIGH, MEDIUM | Default: before execution or before pushing a plan |
| ocd | all four | A plan that will be executed by a future session |

Strict is the default because MEDIUM findings are precisely the ones that turn into wrong guesses at execution time, when nobody is watching.

## Classification Rule

When a finding could sit at two levels, ask what happens if it ships unfixed. If an executor would do the wrong thing, it is HIGH. If an executor would do the right thing after asking a question nobody will be there to answer, it is MEDIUM.

Every `plan-checker` prompt states these four severities and this rule in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

Every `rules-checker` prompt names the same four severities, for the same reason, so a change to the vocabulary reaches all six prompts. The case floors the rules gate adds on top of them belong to its [finding taxonomy](../rules-quality-gate/02-finding-taxonomy.md).
