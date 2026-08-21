# Grind in Public

Personal lifelong-learning workspace maintained by **Wahidyan Kresna Fridayoka**.

[LinkedIn](https://www.linkedin.com/in/wahidyan-kresna-fridayoka/)

This repository is where I practice, review, and document the craft of software engineering, in the open, for as long as I keep building.

This README and [`docs/`](docs/README.md) are for people. AI agents use [`AGENTS.md`](AGENTS.md) and related instruction files; shared repository governance in [`repo-governance/`](repo-governance/README.md) applies to both.

Three agent harnesses are supported: Claude Code, Codex, and opencode. Codex and opencode read `AGENTS.md` directly, Claude Code reads [`CLAUDE.md`](CLAUDE.md), and each tool keeps its configuration and the shared subagents under `.claude/`, `.codex/`, and `.opencode/`, with skills the tools share in `.agents/`. See the [agent harness support policy](repo-governance/conventions/agent-harness-support.md).

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

Build all projects with `npm run build` and run quick checks with `npm test`. See [the Nx workspace guide](docs/how-to/run-nx-workspace.md) for the full workflow. This workspace deliberately avoids technology-specific Nx plugins; the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md) states which kinds are excluded and the one exception.

[Badak Mini](apps/badakmini-cli/README.md) is a small Go CLI behind the repository checks: governance-document word limits, repository-local Markdown links, rule-change announcements, harness capability parity, and project targets. [Workspace commands](repo-governance/development/workspace-commands.md) lists the `npm run check:` command for each one.

## Plans and Specs

Which work is planned before it starts, and which — drills among them — is not, is set by the [plans organization policy](repo-governance/conventions/plans-organization-policy.md). A plan is five documents in [`plans/`](plans/README.md), moving from `ideas/` through `backlog/` and `in-progress/` to `done/`, and delivering directly to `main` at each phase gate. What the software should do is described as Gherkin in [`specs/`](specs/README.md), separate from the code that implements it.

## Documentation

Human-facing project documentation is organized with the [Diátaxis framework](docs/README.md). Repository and agent rules are maintained separately in [`repo-governance/`](repo-governance/README.md).
