---
tldr: "Lists the instruction files, harness configs, subagents, and skills an alignment run reads."
when_to_use: "Use at the start of a harness alignment run, to enumerate everything that has to be compared."
---

# Inventory

Inventory the instruction files, harness configs, subagents, and skills:

```sh
rg --files -g 'AGENTS.md' -g 'CLAUDE.md' -g 'GEMINI.md' -g 'COPILOT.md' -g '.cursorrules' -g 'SKILL.md' -g '!node_modules'
ls .claude/agents .codex/agents .opencode/agents .claude/skills .agents/skills .opencode/skills
```
