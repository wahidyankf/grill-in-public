# opencode Skills

This directory is intentionally empty of shared skills. opencode loads `skills/*/SKILL.md` from `.opencode/`, `.claude/`, and `.agents/`, so a skill kept in the two canonical copies already reaches it; see the [harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md). A third copy here would be a third file to keep identical, and the parity check unions the three sources, so it could not prove they were.

## Available Skills

None. The shared skills live in [`.claude/skills/`](../../.claude/skills/README.md) and [`.agents/skills/`](../../.agents/skills/README.md).

Do not add a skill only here. Claude Code and Codex both load skills, so the parity policy requires all three harnesses to expose the same entries, and `npm run check:harness-parity` fails on a skill the others lack. Every skill goes in the two canonical directories.
