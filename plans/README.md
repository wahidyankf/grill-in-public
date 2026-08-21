# Plans

Plans are this repository's working record of change: why work exists, what it depends on, and what evidence proves it finished. They are not documentation. `docs/` explains the repository to a reader and `repo-governance/` holds its rules; a plan describes one piece of work and retires when that work lands.

Plan documents are created only when the owner explicitly requests a plan. When no plan is requested, work proceeds without creating a plan folder; the [plans organization policy](../repo-governance/conventions/plans-organization-policy.md) owns that authorization boundary.

## Stages

- [`ideas/`](ideas/README.md) — two-pager briefs for problems not yet worth a plan.
- [`backlog/`](backlog/README.md) — full plans prepared but not started.
- [`in-progress/`](in-progress/README.md) — plans under active execution.
- [`done/`](done/README.md) — completed plans, kept as history.

## How a Plan Runs

Three workflows drive the lifecycle:

1. [plan-planning](../repo-governance/workflows/plan-planning.md) turns a prompt into a five-document plan.
2. [plan-quality-gate](../repo-governance/workflows/plan-quality-gate.md) validates it until no findings remain.
3. [plan-execution](../repo-governance/workflows/plan-execution.md) executes it phase by phase and archives it.

Delivery goes directly to `main`: a phase ends, its gate passes, the work is committed and pushed. There are no worktrees and no pull-request flow here.

For structure, naming, checklist rules, and archival, read the [plans organization policy](../repo-governance/conventions/plans-organization-policy.md).
