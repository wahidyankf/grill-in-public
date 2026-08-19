---
tldr: "Divides the gate's work between plan-checker and plan-fixer, and links the loop's bounds."
when_to_use: "Use when running the gate or reviewing what each subagent may do."
---

# Check and Fix Loop

```text
plan-checker --> findings --> plan-fixer --> plan-checker --> clean twice --> pass
     ^                                            |
     +--------------- up to 7 cycles -------------+
```

## plan-checker

Reads every document in the plan folder and reports findings with a severity, a `file:line` citation, and the specific rule the finding violates. It edits nothing. A finding without a cited rule is an opinion, and the checker does not report opinions.

The checker verifies the plan against itself as well as against the rules: a command named in `delivery.md` must exist, every path the checklist names must appear in the file-impact tree, and scenarios must bind both ways — one in `delivery.md` matching the one in `prd.md` verbatim, and one in `prd.md` inlined by a RED step.

Every `plan-checker` prompt states this reporting rule and these internal-consistency checks in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

## plan-fixer

Edits plan documents to resolve findings. Its mandate is clarity, not authority:

It may:

- add a missing file path, verbatim command, acceptance criterion, or executor tag
- split a checkbox that hides several actions, or a behavior cycle binding several scenarios
- add a missing `### Phase N Gate`, gate item, or Pause Safety note
- inline a Gherkin scenario verbatim from `prd.md`
- correct a file-impact tree that omits a path the checklist touches
- fix a link, a heading, or a diagram that uses Mermaid instead of ASCII

It may not:

- change a decision the owner made, or the reasoning recorded for it
- widen or narrow the plan's scope
- delete a step, or weaken an acceptance criterion, to make a finding disappear
- invent a command, a path, or a metric you have not verified exists

It uses the shell to verify, never to build: to check that a command, path, or target it is about to write into the plan actually exists, and never to run the work the plan describes.

A finding it cannot fix within that mandate is left open and reported, not silently dropped.

The `plan-fixer` prompt states these two lists verbatim and carries the same shell rule in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies, or the fixer and the workflow start authorizing different things.

## Why Both

[Role separation](03-role-separation.md) says why the checker and the fixer are two agents rather than one.

The loop's bounds are in the [workflow](../plan-quality-gate.md#loop-bounds), beside its recovery guidance.
