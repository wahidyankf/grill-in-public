---
tldr: "Lists where applicable guidance is searched for before a rule is edited."
when_to_use: "Use at the start of a rules propagation run, before changing any guidance."
---

# Inventory

Inventory applicable guidance before editing. Start with `AGENTS.md`, `repo-governance/`, and its `principles/`; also search for instruction and skill files:

```sh
rg --files -g 'AGENTS.md' -g 'SKILL.md' -g 'CLAUDE.md' -g 'GEMINI.md' \
  -g 'COPILOT.md' -g '!node_modules'
```
