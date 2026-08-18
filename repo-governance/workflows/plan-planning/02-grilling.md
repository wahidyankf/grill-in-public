---
tldr: "Requires open plan decisions to be resolved by structured options before authoring."
when_to_use: "Use after exploration, when open decisions remain in plan-planning."
---

# Grilling

Every decision left open at authoring time becomes a guess written into a checklist, and a guess in a checklist is executed as though it were a decision.

## The Rules Apply Unchanged

The [grilling-with-options policy](../../conventions/grilling-with-options-policy.md) is normative here: two to four mutually exclusive options, each with its own specific trade-off, exactly one marked Recommended, one decision per question, plus the standing write-in and chat options. Use the harness's native question tool; see the [grilling harness binding](../../conventions/grilling-harness-binding.md).

## What to Grill in a Plan

Grill the decisions that change what gets built:

- scope boundaries, especially what is deliberately excluded
- the approach, when two designs would both work and differ in cost
- the delivery order, when one sequence leaves the repository broken between phases
- anything that would be expensive to reverse once the first phase lands

Do not grill what the repository already answers, what the owner has already decided in this session, or what has an obvious default. Those are questions that cost attention and return nothing.

## Recording the Outcome

Every answer is written into the plan: the decision, the option chosen, and the reason. `brd.md` carries the scope and rationale decisions; `tech-docs.md` carries the design decisions. A decision the owner cannot trace back to a question is one they never really made, and a decision recorded without its reason is one that gets reversed by the next reader who disagrees.

## Second Pass

If exploration or research changes the picture after the first round, grill again before authoring. Re-grilling costs a message; discovering the wrong decision mid-execution costs a phase.
