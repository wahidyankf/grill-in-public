# Codex Subagents

Codex discovers these agents automatically from this directory. It reads only `*.toml` here, so this README is ignored.

## Available Agents

- [`drill-reviewer.toml`](drill-reviewer.toml) — reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer.
- [`plan-checker.toml`](plan-checker.toml) — reviews a plan under `plans/` against the plans organization policy and reports findings by severity. Use it before executing a plan or after changing one; it never edits the plan.
- [`plan-fixer.toml`](plan-fixer.toml) — resolves plan-checker findings by editing plan documents for clarity, never changing a decision or the scope. Use it between checker runs inside the plan-quality-gate loop.
- [`repo-explorer.toml`](repo-explorer.toml) — read-only explorer that reports where code, documentation, and governance rules live. Use it to locate things or check which rule applies before making a change.
- [`rules-checker.toml`](rules-checker.toml) — reviews every rule-bearing file for contradictions, duplication, orphan references, and gaps between harnesses. Use it inside the rules-quality-gate loop; it never edits anything.
- [`rules-fixer.toml`](rules-fixer.toml) — resolves rules-checker findings by editing governance documents, instruction files, and harness prompts, never settling a contradiction. Use it between checker runs inside the rules-quality-gate loop.

Each role is mirrored in [`.claude/agents/`](../../.claude/agents/README.md) and [`.opencode/agents/`](../../.opencode/agents/README.md) and must stay at parity; see the [agent harness support policy](../../repo-governance/conventions/agent-harness-support.md).
