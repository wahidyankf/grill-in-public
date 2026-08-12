# Repository Governance

This directory contains shared repository governance for human contributors and
AI agents. It supplements human-facing [`README.md`](../README.md) and
[`docs/`](../docs/), plus AI-agent guidance such as
[`AGENTS.md`](../AGENTS.md). Both audiences read only the governance document
relevant to their work.

## Foundations

[`principles/`](principles/README.md) contains the foundational rules for this
directory. All policies and workflows here must follow them; resolve any
contradiction before accepting lower-level guidance.

## How to Use This Directory

- Keep `AGENTS.md` concise and link it to shared governance that agents need.
- Put detailed shared rules, workflows, and specialized policies here.
- Link to a governance document from the appropriate human or agent entry point
  when it becomes required for a recurring task or a specific area.
- Read the smallest relevant set of documents; do not load unrelated guidance.

## Document Conventions

Use focused, descriptive filenames such as:

- `exercise-layout.md` for exercise organization and naming details.
- `testing-policy.md` for test strategy and quality expectations.
- `dependency-policy.md` for package review, pinning, and audit procedures.
- `nx-workspace-policy.md` for the permitted Nx workspace integration.
- `testing-policy.md` for quick and integration-test responsibilities.
- `commit-hook-policy.md` for required Git-hook behavior and bypass exceptions.
- `badakmini-cli-policy.md` for repository-local validation ownership and scope.
- `contribution-workflow.md` for branch, pull request, or release practices.
- `principles/` for foundational governance rules.

Each document should state its scope, give actionable rules, and link to any
source-of-truth files or commands. Avoid duplicating `AGENTS.md`; keep shared
rules there and move only extended context here.

## Maintaining the Guidance

Update the relevant document when a practice changes. If a detailed rule becomes
universal or essential for every agent task, summarize it in `AGENTS.md` and
retain the full shared rationale here.
