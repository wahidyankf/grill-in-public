# Repository Governance

This directory contains shared repository governance for human contributors and
AI agents. It supplements human-facing [`README.md`](../README.md) and
[`docs/`](../docs/), plus AI-agent guidance such as
[`AGENTS.md`](../AGENTS.md). Both audiences read only the governance document
relevant to their work.

## Directory Index

- [Development Governance](development/README.md) for code, testing, hook, Nx,
  and validation policies.
- [Documentation Index Policy](documentation-index-policy.md) for keeping
  `docs/` and this directory recursively discoverable.
- [Governance Principles](principles/README.md) for the foundations every policy
  and workflow must follow.
- [Repository Workflows](workflows/README.md) for repeatable procedures.

## How to Use This Directory

- Keep `AGENTS.md` concise and link it to shared governance that agents need.
- Put detailed shared rules, workflows, and specialized policies here.
- Link to a governance document from the appropriate human or agent entry point
  when it becomes required for a recurring task or a specific area.
- Read the smallest relevant set of documents; do not load unrelated guidance.

## Document Conventions

Use focused, descriptive filenames such as:

- `exercise-layout.md` for exercise organization and naming details.
- `dependency-policy.md` for package review, pinning, and audit procedures.
- `development/` for code, testing, Nx, hook, and validation policies.
- `documentation-index-policy.md` for recursive README requirements.
- `contribution-workflow.md` for branch, pull request, or release practices.
- `principles/` for foundational governance rules.

Each document should state its scope, give actionable rules, and link to any
source-of-truth files or commands. Avoid duplicating `AGENTS.md`; keep shared
rules there and move only extended context here.

## Maintaining the Guidance

Update the relevant document when a practice changes. If a detailed rule becomes
universal or essential for every agent task, summarize it in `AGENTS.md` and
retain the full shared rationale here.
