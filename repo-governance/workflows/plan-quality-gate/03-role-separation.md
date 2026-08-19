---
tldr: "Says why the plan gate splits checking from fixing between two subagents."
when_to_use: "Use when questioning why plan-checker and plan-fixer are separate agents."
---

# Role Separation

Separating the roles keeps the checker honest. A single agent that both finds and fixes has an incentive to find only what it can fix, and the findings it cannot fix are the ones that matter most.

The [check and fix loop](02-check-fix-loop.md) divides the work between the two roles, and states what each may and may not do.
