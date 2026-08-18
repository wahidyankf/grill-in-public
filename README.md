# Grind in Public

Personal lifelong-learning workspace maintained by **Wahidyan Kresna Fridayoka**.

[LinkedIn](https://www.linkedin.com/in/wahidyan-kresna-fridayoka/)

This repository is where I practice, review, and document the craft of software engineering, in the open, for as long as I keep building.

This README and [`docs/`](docs/README.md) are for people. AI agents use [`AGENTS.md`](AGENTS.md) and related instruction files; shared repository governance in [`repo-governance/`](repo-governance/README.md) applies to both.

Three agent harnesses are supported: Claude Code, Codex, and opencode. Codex and opencode read `AGENTS.md` directly, Claude Code reads [`CLAUDE.md`](CLAUDE.md), and each tool keeps its configuration and shared `drill-reviewer` and `repo-explorer` subagents under `.claude/`, `.codex/`, and `.opencode/`. See the [agent harness support policy](repo-governance/conventions/agent-harness-support.md).

## Hands-On by Design

Unlike my `ose-*` repositories, this repository is intentionally not fully automation-first. I will complete parts of the work by hand—especially coding exercises and drills—to build muscle memory, reinforce the fundamentals, and practice recalling and explaining solutions without help.

Tools can support review and feedback, but they should not replace the hands-on practice at the heart of this repository.

## Focus Areas

- Data structures and algorithms
- Coding exercises and problem-solving patterns
- System design
- Computer science and software engineering fundamentals
- Languages, tools, and engineering craft
- Retrospective notes on what stuck and what did not

## Practice Workflow

1. Choose a topic or problem worth understanding.
2. Work through it by hand under deliberate constraints, without shortcuts.
3. Record the solution and reasoning.
4. Review trade-offs, mistakes, and possible improvements.
5. Revisit the exercise until the approach is clear and repeatable.

## Goal

Build strong fundamentals, explain solutions clearly, and keep compounding skill as a software engineer over years rather than cramming for a deadline.

## Nx Workspace

The repository uses Nx as a raw task runner for its npm workspaces:

- `apps/` holds runnable applications.
- `libs/` holds reusable packages consumed by applications.

The [Dummy App](apps/dummy-app/README.md) TypeScript CLI consumes [`@grind-in-public/dummy-lib`](libs/dummy-lib/README.md). Run it with `npm run run:dummy`; build all projects with `npm run build`; and run quick checks with `npm test`. See [the Nx workspace guide](docs/how-to/run-nx-workspace.md) for the full workflow. This workspace deliberately does not use framework or language-specific Nx plugins. TypeScript projects use project-local ESLint commentary checks alongside Biome; passing lint complements, but does not replace human review of explanatory comments.

[Badak Mini](apps/badakmini-cli/README.md) is a small Go CLI that checks governance-document word limits and repository-local Markdown links. Run `npm run check:governance` or `npm run check:markdown-links`.

## Documentation

Human-facing project documentation is organized with the [Diátaxis framework](docs/README.md). Repository and agent rules are maintained separately in [`repo-governance/`](repo-governance/README.md).
