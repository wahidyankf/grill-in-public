# Repository Governance

This directory contains shared repository governance for human contributors and AI agents.

## Related Entry Points

- [Repository README](../README.md) — the human project overview. Use it for repository purpose, setup, and human-facing navigation.
- [Documentation Hub](../docs/README.md) — human Diátaxis documentation. Use it when the task is learning, reference, explanation, or a how-to procedure.
- [AGENTS.md](../AGENTS.md) — concise instructions for coding agents. Use it at the start of repository work before loading focused governance below.

## Directory Index

- [Governance Conventions](conventions/README.md) — shared standards for repository content and practices. Use them when creating or reviewing work covered by a convention.
- [Development Governance](development/README.md) — policies for code, testing, hooks, Nx, and validation. Use it when changing executable code or tooling.
- [Documentation Index Policy](documentation-index-policy.md) — README and metadata requirements for repository documents. Use it when adding, moving, or maintaining Markdown under `docs/` or `repo-governance/`.
- [Governance Principles](principles/README.md) — foundations every policy and workflow must follow. Use them before resolving a governance conflict.
- [Repository Workflows](workflows/README.md) — repeatable repository procedures. Use the relevant workflow whenever a task has a defined process.

## Delivery Directories

- [Plans](../plans/README.md) — the working record of change, staged from idea to archive. Use it when planning, executing, or reviewing a piece of work.
- [Specifications](../specs/README.md) — Gherkin behavior for the workspace's apps and libraries. Use it when adding or changing what the software does.

## Gate History

Each [rules-quality-gate](workflows/rules-quality-gate.md) run appends one line to `local-tmp/gate-history/rules-quality-gate.md`, which is untracked. The log grows without bound, and this index does not, so neither rule has to give way. An open finding is written into the affected document as well, where a reader of the rule will meet it.

## How to Use This Directory

- Keep `AGENTS.md` concise and link it to shared governance that agents need.
- Put detailed shared rules, workflows, and specialized policies here.
- Link to a governance document from the appropriate human or agent entry point when it becomes required for a recurring task or a specific area.
- Read the smallest relevant set of documents; do not load unrelated guidance.

## Document Conventions

Name each document for its subject, in lowercase with hyphens, following the [document naming policy](conventions/document-naming-policy.md). Group related documents into a subdirectory — `conventions/` for stable shared standards, `development/` for code, testing, Nx, hook, and validation policies, `principles/` for foundational rules, and `workflows/` for repeatable procedures.

Each document should state its scope, give actionable rules, and link to any source-of-truth files or commands. Avoid duplicating `AGENTS.md`; keep shared rules there and move only extended context here.

## Maintaining the Guidance

Update the relevant document when a practice changes. If a detailed rule becomes universal or essential for every agent task, summarize it in `AGENTS.md` and retain the full shared rationale here.
