---
tldr: "Fixes the meaning of agent, harness, instruction file, and subagent for this repository."
when_to_use: "Use when writing or reviewing any text about agents, harnesses, instruction files, or subagents."
---

# Agent Vocabulary

## Purpose

These words describe different things, and mixing them makes a rule read as if it governs the wrong object. Use them as defined here in governance, documentation, commit messages, and code identifiers.

## Terms

**Harness** — the tool that runs a model: it holds the loop, the tools, the permissions, and the workspace. Claude Code, Codex, and opencode are harnesses. The industry shorthand is `agent = model + harness`.

**Agent** — a model plus its harness, working as one system. Use it for the whole thing, not for either half.

**Agent instruction file** — a file a harness reads for repository rules, such as `AGENTS.md` or `CLAUDE.md`. It is never called a harness file: the harness reads it, and one instruction file can serve several harnesses.

**Harness directory** — the directory holding one harness's project configuration: `.claude/`, `.codex/`, `.opencode/`, or the shared `.agents/`.

**Subagent** — a named role a harness can spawn with its own instructions and permissions, defined under a harness directory's `agents/`. Codex also calls these custom agents.

**Drill grilling** — questioning the owner to rehearse something being learned. The owner answers, and the agent reviews. The repository name refers to the practice itself, the grind, not to either kind of grilling.

**Grilling-with-options** — questioning the owner to resolve an open decision about the work itself, in the structured form the [grilling-with-options policy](grilling-with-options-policy.md) requires. The agent asks and the owner decides, so the roles are the reverse of a drill.

## Rules

Name the object you mean. A limit on `AGENTS.md` is a limit on an instruction file, not on a harness. Support for opencode is support for a harness, not for a file.

Qualify grilling whenever both senses could fit the sentence. The two differ in who answers, and nothing in the word says which is meant, so a rule about deciding must say grilling-with-options and a rule about practice must say drill.

Keep code identifiers on the same vocabulary as the prose that documents them.

Badak Mini's `harness` command group is the exception, and it stays. It names the family of harness-related checks, such as `harness instruction-size validate`, not the files each check reads. Renaming it would break the npm scripts and the shared command grammar for no gain in clarity.
