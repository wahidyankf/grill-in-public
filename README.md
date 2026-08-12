# SWE Grilling

Personal interview-preparation workspace maintained by **Wahidyan Kresna
Fridayoka**.

[LinkedIn](https://www.linkedin.com/in/wahidyan-kresna-fridayoka/)

This repository is where I practice, review, and document the skills needed for
software engineering interviews.

This README and [`docs/`](docs/README.md) are for people. AI agents use
[`AGENTS.md`](AGENTS.md) and related instruction files; shared repository
governance in [`repo-governance/`](repo-governance/README.md) applies to both.

## Hands-On by Design

Unlike my `ose-*` repositories, this repository is intentionally not fully
automation-first. I will complete parts of the work by hand—especially coding
exercises and drills—to build muscle memory, reinforce the fundamentals, and
practice recalling and communicating solutions under interview conditions.

Tools can support review and feedback, but they should not replace the hands-on
practice at the heart of this repository.

## Focus Areas

- Data structures and algorithms
- Coding exercises and problem-solving patterns
- System design
- Computer science and software engineering fundamentals
- Behavioral interview preparation
- Mock interviews and retrospective notes

## Practice Workflow

1. Choose a topic or interview question.
2. Work through it by hand under realistic interview constraints.
3. Record the solution and reasoning.
4. Review trade-offs, mistakes, and possible improvements.
5. Revisit the exercise until the approach is clear and repeatable.

## Goal

Build strong fundamentals, communicate solutions clearly, and become more
confident and consistent in software engineering interviews.

## Nx Workspace

The repository uses Nx as a raw task runner for its npm workspaces:

- `apps/` holds runnable applications.
- `libs/` holds reusable packages consumed by applications.

The starter TypeScript CLI consumes `@swe-grilling/dummy-lib`. Run it with
`npm run run:dummy`; build all projects with `npm run build`; and run its test
with `npm test`. See [the Nx workspace guide](docs/how-to/run-nx-workspace.md)
for the full workflow. This workspace deliberately does not use framework or
language-specific Nx plugins.

[Badak Mini](apps/badak-mini/README.md) is a small Go CLI that checks the
repository's governance-document word limits. Run it through
`npm run check:governance`.

## Documentation

Human-facing project documentation is organized with the
[Diátaxis framework](docs/README.md). Repository and agent rules are maintained
separately in [`repo-governance/`](repo-governance/README.md).
