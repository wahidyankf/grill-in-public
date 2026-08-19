---
tldr: "Requires every harness to expose the same subagents, skills, and commands, and records where each one lives."
when_to_use: "Use when adding, renaming, or removing a subagent, skill, or command for any harness."
---

# Harness Capability Parity Policy

## Scope

This policy covers the capabilities a harness loads from repository files: subagents, skills, and commands. The instruction files themselves belong to the [agent instruction alignment policy](agent-instruction-alignment-policy.md), and which harness reads which file belongs to the [agent harness support policy](agent-harness-support.md).

## Rule

A capability is defined once for the repository, not for one tool. Every harness that supports it must expose the same entries: the same count and the same names. The same role must keep the same purpose and the same permission posture; wording may be reworded to fit each format, capability must not diverge.

Add a capability to every supporting harness in the change that introduces it. Never delete an entry from the other harnesses to make the counts match, and never leave one harness ahead; both hide the gap the parity rule exists to surface.

## Capability Directories

Verified against each tool's documentation in August 2026. Test discovery yourself before relying on it, because these paths change between releases:

| Capability | Claude Code | Codex | opencode |
| --- | --- | --- | --- |
| Subagents | `.claude/agents/*.md` | `.codex/agents/*.toml` | `.opencode/agents/*.md` |
| Skills | `.claude/skills/<name>/SKILL.md` | `.agents/skills/<name>/SKILL.md` | `.opencode/skills/`, and also `.claude/skills/` and `.agents/skills/` |
| Commands | `.claude/commands/*.md` | Unsupported for a project | `.opencode/commands/*.md` |

A shared skill therefore needs one copy under `.claude/skills/` and one under `.agents/skills/`; opencode reads either, so it needs no third copy.

## Unsupported Capabilities

Where a harness cannot load a capability, record it in the table above rather than in a comment or in memory, and say so in the affected `README.md`. An unsupported harness is exempt from the count for that capability only. Claude Code has merged commands into skills, so a `.claude/skills/` entry also answers as `/name`.

## Verification

```sh
npm run check:harness-parity
```

Badak Mini compares the entries per capability, ignores each directory's `README.md` index, and skips a capability that no harness uses yet. It runs during a push that changes a harness directory, and the [Harness Alignment](../workflows/harness-alignment.md) workflow runs it too.

The command proves the counts and names match. It cannot read intent, so still confirm by hand that a mirrored entry keeps the same purpose and permission posture.
