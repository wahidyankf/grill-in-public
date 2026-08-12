# Repository Governance

This directory contains detailed repository guidance that supplements the
root-level [`AGENTS.md`](../AGENTS.md). It supports progressive disclosure:
agents and contributors should start with `AGENTS.md`, then read only the
governance document relevant to the work at hand.

## How to Use This Directory

- Keep `AGENTS.md` as the concise, always-read guide for repository-wide rules.
- Put longer explanations, workflows, and specialized policies here.
- Link to a governance document from `AGENTS.md` when it becomes required for a
  recurring task or a specific area of the repository.
- Read the smallest relevant set of documents; do not load unrelated guidance.

## Document Conventions

Use focused, descriptive filenames such as:

- `exercise-layout.md` for exercise organization and naming details.
- `testing-policy.md` for test strategy and quality expectations.
- `dependency-policy.md` for package review, pinning, and audit procedures.
- `contribution-workflow.md` for branch, pull request, or release practices.

Each document should state its scope, give actionable rules, and link to any
source-of-truth files or commands. Avoid duplicating `AGENTS.md`; keep shared
rules there and move only extended context here.

## Maintaining the Guidance

Update the relevant document when a practice changes. If a detailed rule becomes
universal or essential for every task, summarize it in `AGENTS.md` and retain
the full rationale here.
