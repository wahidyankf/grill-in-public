---
tldr: "Requires an agent to resolve an open decision with structured options, not an open prose question."
when_to_use: "Use whenever an agent must resolve an open design or scope decision with the owner."
---

# Grilling-With-Options Policy

## Scope

This policy governs how an agent questions the owner to resolve an open decision, adapted from the `ose-public` convention of the same name. It does not govern practice drills, where the owner is the one being questioned; the [agent vocabulary](agent-vocabulary.md) separates the two senses. The [harness binding](grilling-harness-binding.md) maps the rules below onto each harness's tool.

## Rules

1. **Explore before asking.** Read the repository first. A question the code, a policy, or the Git history already answers must not be asked.
2. **Two to four substantive options,** mutually exclusive, together covering the realistic decision space. Fewer than two is a yes-or-no confirmation, not a decision; more than four means the space was never pruned.
3. **Each option states its own trade-off** in one sentence, specific to this decision. "Simpler" is not a trade-off.
4. **Exactly one option is marked Recommended,** in that option's own label rather than by its position in the list, and with a reason grounded in repository state or a stated constraint. The [harness binding](grilling-harness-binding.md) owns the form the mark takes, because a mark the owner has to infer is one they can read past. Recommending nothing withholds the judgment the question exists to supply; recommending two is the same evasion in disguise.
5. **One decision per question.** Batch only decisions where one answer constrains the other. Unrelated decisions are separate questions.
6. **Use the harness's native question tool** when the session is interactive, and the Markdown fallback only when it is not.
7. **A write-in answer counts as much as a listed option.** When it opens a new branch, grill that branch before proceeding.
8. **Every question carries two standing options** beyond its substantive ones: a blank-state write-in, and a "chat about this" path that drops the options and discusses the decision in prose. These are escape hatches, not branches, so they do not count toward the cap. A harness whose list holds only four entries cannot show four substantive options and both escapes; there, prune to three substantive options rather than drop an escape, because a question with no way out is worse than a question with one fewer answer.

## Validation

A drafted question is checked before it is asked; [validation](grilling-with-options-policy/validation.md) holds the checklist and what makes a question invalid.

## Verification

No check can read a question asked at runtime, so this policy is verified in review, like the [task tracking policy](task-tracking-policy.md). What is checkable is that every harness exposes the `grill-me` skill; the [harness capability parity policy](harness-capability-parity-policy.md) owns that.
