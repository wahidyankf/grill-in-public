---
tldr: "Divides the gate's work between plan-checker and plan-fixer and bounds the loop."
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

The checker verifies the plan against itself as well as against the rules: a command named in `delivery.md` must exist, a path in the file-impact tree must be plausible, and a scenario in `delivery.md` must match the one in `prd.md` verbatim.

## plan-fixer

Edits plan documents to resolve findings. Its mandate is clarity, not authority:

- It may add a missing path, command, acceptance criterion, executor tag, gate, or Pause Safety note.
- It may split a coarse checkbox, or a behavior cycle binding several scenarios.
- It may not change a decision the owner made, alter scope, or remove a step to make a finding disappear.

A finding it cannot fix within that mandate is left open and reported, not silently dropped.

## Loop Bounds

Two consecutive clean runs end the loop. Seven cycles end it too, with the remaining findings reported. Cycle five raises a warning: a plan still finding new problems that late is usually structurally wrong rather than imprecise.

## Why Both

Separating the roles keeps the checker honest. A single agent that both finds and fixes has an incentive to find only what it can fix, and the findings it cannot fix are the ones that matter most.
