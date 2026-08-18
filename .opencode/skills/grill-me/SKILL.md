---
name: grill-me
description: Resolve open design decisions by interrogating the owner with structured multiple-choice questions, one decision at a time, until nothing is left ambiguous. Use when the owner says "grill me", asks to stress-test a plan or design, or when a task cannot proceed without a decision only the owner can make.
---

# Grill Me

Resolve every open decision before building, by asking rather than assuming.

## When to Activate

Activate when the owner says "grill me", asks for a plan or design to be stress-tested, or when work is blocked on a decision the repository does not already answer. Do not activate for a drill: there the owner answers questions to practice, and this skill is the reverse.

## Rules

The [grilling-with-options policy](../../../repo-governance/conventions/grilling-with-options-policy.md) is normative. Read it rather than working from memory. In short: explore first, offer two to four mutually exclusive options, give each a specific trade-off, mark exactly one Recommended, ask one decision per question, and always carry the write-in and chat options.

## Mechanism

opencode exposes the `question` tool, and this skill uses it whenever the session is interactive. Put the Recommended option first and keep each header at twelve characters or fewer. The [harness binding](../../../repo-governance/conventions/grilling-harness-binding.md) holds the shaping details and the Markdown fallback for a non-interactive session.

## After the Grilling

Summarize each decision and the reason it was chosen, then say what the answers changed about the plan. A decision the owner cannot trace back to a question is one they never really made.
