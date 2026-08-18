# Codex Subagents

Codex discovers these agents automatically from this directory. It reads only `*.toml` here, so this README is ignored.

## Available Agents

- [`drill-reviewer.toml`](drill-reviewer.toml) — reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer.
- [`repo-explorer.toml`](repo-explorer.toml) — read-only explorer that reports where code, documentation, and governance rules live. Use it to locate things or check which rule applies before making a change.

Each role is mirrored in [`.claude/agents/`](../../.claude/agents/README.md) and [`.opencode/agents/`](../../.opencode/agents/README.md) and must stay at parity; see the [agent harness support policy](../../repo-governance/conventions/agent-harness-support.md).
