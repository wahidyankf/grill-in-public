# Claude Code Subagents

Claude Code discovers these agents automatically from this directory. It ignores this README because the file carries no agent frontmatter.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer.
- [`repo-explorer.md`](repo-explorer.md) — read-only explorer that reports where code, documentation, and governance rules live. Use it to locate things or check which rule applies before making a change.

Each role is mirrored in [`.codex/agents/`](../../.codex/agents/README.md) and [`.opencode/agents/`](../../.opencode/agents/README.md) and must stay at parity; see the [agent harness support policy](../../repo-governance/conventions/agent-harness-support.md).
