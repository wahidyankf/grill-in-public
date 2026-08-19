---
tldr: "Maps the grilling rules onto each harness's question tool and the Markdown fallback."
when_to_use: "Use when asking a structured decision question, or when adding a harness that must ask one."
---

# Grilling Harness Binding

## Scope

The [grilling-with-options policy](grilling-with-options-policy.md) owns the rules. This document owns only how those rules reach the owner in each harness, so the shared `grill-me` skill can point at one table instead of naming a tool per copy and drifting apart.

## Details

Read the part the session needs rather than the whole binding:

- [Question Tools](grilling-harness-binding/question-tools.md) — the tool each harness exposes and its status here.
- [Markdown Fallback](grilling-harness-binding/markdown-fallback.md) — the inline rendering for a session that cannot ask.

## Shaping a Question

Keep the header short, at most twelve characters, since every harness renders it as a chip or label. Put the Recommended option first and suffix its label with `(Recommended)`; a reader who stops after one option should still see the recommendation. Give each option a one-to-five word label and a one-sentence trade-off. Ask at most four questions in one call, one decision each, and never request multiple selections for a decision that has one answer.

Read the row for your harness in the [question tools](grilling-harness-binding/question-tools.md) table before adding the options. Where it records that the client appends a free-text entry, add only the chat option explicitly; where the binding is unverified, add the write-in yourself and drop it if the client turns out to supply one, since a duplicated write-in is a smaller failure than a question that offers none. Rule 8 of the policy says what gives when the harness's list runs out of room.

## Adding a Harness

Record the new harness's tool in the [question tools](grilling-harness-binding/question-tools.md) table. Check which shared skill directories it reads before copying a `SKILL.md` into it; the [harness capability parity policy](harness-capability-parity-policy.md) owns how many copies a shared skill needs. State plainly whether the binding is verified or only documented; an unverified binding is useful, and a binding claimed as working when nobody ran it is not.
