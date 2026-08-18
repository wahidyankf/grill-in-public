# Claude Code Subagents

Claude Code discovers these agents automatically from this directory. It ignores this README because the file carries no agent frontmatter.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer.
- [`plan-checker.md`](plan-checker.md) — reviews a plan under `plans/` against the plans organization policy and reports findings by severity. Use it before executing a plan or after changing one; it never edits the plan.
- [`plan-fixer.md`](plan-fixer.md) — resolves plan-checker findings by editing plan documents for clarity, never changing a decision or the scope. Use it between checker runs inside the plan-quality-gate loop.
- [`repo-explorer.md`](repo-explorer.md) — read-only explorer that reports where code, tests, documentation, and governance rules live. Use it to locate things or check which rule applies before making a change.
- [`rules-checker.md`](rules-checker.md) — reviews every rule-bearing file for contradictions, duplication, orphan references, and gaps between harnesses. Use it inside the rules-quality-gate loop; it never edits anything.
- [`rules-fixer.md`](rules-fixer.md) — resolves rules-checker findings by editing governance documents, instruction files, and harness prompts, never settling a contradiction. Use it between checker runs inside the rules-quality-gate loop.

Each role is mirrored in [`.codex/agents/`](../../.codex/agents/README.md) and [`.opencode/agents/`](../../.opencode/agents/README.md) and must stay at parity; see the [agent harness support policy](../../repo-governance/conventions/agent-harness-support.md).
