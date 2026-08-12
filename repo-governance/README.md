# Repository Governance

This directory contains shared repository governance for human contributors and
AI agents.

## Related Entry Points

- [Repository README](../README.md) — the human project overview. Use it for
  repository purpose, setup, and human-facing navigation.
- [Documentation Hub](../docs/README.md) — human Diátaxis documentation. Use it
  when the task is learning, reference, explanation, or a how-to procedure.
- [AGENTS.md](../AGENTS.md) — concise instructions for coding agents. Use it at
  the start of repository work before loading focused governance below.

## Directory Index

- [Development Governance](development/README.md) — policies for code, testing,
  hooks, Nx, and validation. Use it when changing executable code or tooling.
- [Documentation Index Policy](documentation-index-policy.md) — README and
  metadata requirements for repository documents. Use it when adding, moving,
  or maintaining Markdown under `docs/` or `repo-governance/`.
- [Governance Principles](principles/README.md) — foundations every policy and
  workflow must follow. Use them before resolving a governance conflict.
- [Repository Workflows](workflows/README.md) — repeatable repository
  procedures. Use the relevant workflow whenever a task has a defined process.

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
