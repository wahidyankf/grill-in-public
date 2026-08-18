---
tldr: "Maps the grilling rules onto each harness's question tool and the Markdown fallback."
when_to_use: "Use when asking a structured decision question, or when adding a harness that must ask one."
---

# Grilling Harness Binding

## Scope

The [grilling-with-options policy](grilling-with-options-policy.md) owns the rules. This document owns only how those rules reach the owner in each harness, so the three `grill-me` skills can name their tool without restating the rules and drifting apart.

## Tools

| Harness | Tool | Notes |
| --- | --- | --- |
| Claude Code | `AskUserQuestion` | Verified here. The harness appends a free-text entry, which serves as the Rule 8 write-in. |
| Codex | `request_user_input` | Needs `default_mode_request_user_input` in `.codex/config.toml`; Codex marks the feature as in development, so its effect is unverified here. |
| opencode | `question` | Documented by opencode; unverified here. |

A tool that returns a structured choice removes the parsing step, which is why Rule 6 prefers it over prose. Only a session with no such tool falls back to Markdown.

## Shaping a Question

Keep the header short, at most twelve characters, since every harness renders it as a chip or label. Put the Recommended option first and suffix its label with `(Recommended)`; a reader who stops after one option should still see the recommendation. Give each option a one-to-five word label and a one-sentence trade-off. Ask at most four questions in one call, one decision each, and never request multiple selections for a decision that has one answer.

Claude Code and opencode supply the free-text entry themselves, so add only the chat option explicitly. That leaves two or three substantive options plus chat inside a four-option list.

## Markdown Fallback

When the session cannot ask interactively, print the question inline and stop for an answer:

```text
**Where should the convention live?**

1. **Conventions directory (Recommended)** — matches the adjacent standards, but needs a README entry.
2. **Development directory** — sits with the executable policies, but this rule is not about code.
3. **Other — type your own answer**
4. **Let's chat about this**
```

The fallback is a rendering of the same question, not a weaker one: it still carries the trade-offs, the single recommendation, the write-in, and the chat option.

## Adding a Harness

Record the new harness's tool in the table above and give it a `grill-me` skill, as the [harness capability parity policy](harness-capability-parity-policy.md) requires. State plainly whether the binding is verified or only documented; an unverified binding is useful, and a binding claimed as working when nobody ran it is not.
