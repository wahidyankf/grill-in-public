---
tldr: "Renders a grilling question inline for a session that cannot ask interactively."
when_to_use: "Use when the harness exposes no question tool, or the session cannot stop to ask."
---

# Markdown Fallback

When the session cannot ask interactively, print the question inline and stop for an answer:

```text
**Where should the convention live?**

1. **Conventions directory (Recommended)** — matches the adjacent standards, but needs a README entry.
2. **Development directory** — sits with the executable policies, but this rule is not about code.
3. **Other — type your own answer**
4. **Let's chat about this**
```

The fallback is a rendering of the same question, not a weaker one: it still carries the trade-offs, the single recommendation, the write-in, and the chat option.
