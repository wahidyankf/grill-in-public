---
tldr: "Requires learnings to be logged during execution and routed to a durable home before archival."
when_to_use: "Use while executing a plan and when running its final Knowledge Capture phase."
---

# Knowledge Capture

A plan that teaches something and then archives it has wasted the lesson. This repository exists to learn, so capture is a gate rather than a courtesy.

## The Running Log

`learnings.md` is written during execution, in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to prevent the failure it targets. It is not reconstructed from memory afterwards, because reconstruction records what the author already believed rather than what actually happened.

Each entry is one short paragraph: what happened, and what a future reader should do differently.

## The Capture Phase

The final phase of every substantive plan, immediately before archival, is Knowledge Capture. It triages every entry in `learnings.md` and routes each to exactly one durable home:

- a rule that belongs in `repo-governance/`, integrated through the [rules-propagation](../../workflows/rules-propagation.md) workflow
- a document under `docs/` when the lesson helps a reader rather than binding an agent
- a subagent or skill instruction, when the lesson changes how a role behaves
- code or a test, when the lesson is executable — a regression test is the strongest form of capture
- a new two-pager in `plans/ideas/`, when the lesson is large enough to be its own work
- discarded, with a one-line reason, when it is not generalizable

Before routing, every surviving entry passes two checks: it contains no secret or sensitive detail, and it is relevant to this repository rather than to one incident.

## Archival Is Blocked

A plan does not move to `done/` until every `learnings.md` entry has reached a terminal state — routed or discarded with a reason — or the plan records the explicit escape `No generalizable learnings — <reason>`.

`learnings.md` is transient. It moves with the plan folder and may be deleted from `done/` later, so nothing durable may depend on it surviving.
