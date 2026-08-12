# Documentation

This directory contains human-facing documentation for SWE Grilling: interview
preparation material, exercise guidance, and project concepts. AI-agent
instructions belong in [`AGENTS.md`](../AGENTS.md) and related instruction
files. Shared rules and governance procedures apply to people and agents, and
belong in [`repo-governance/`](../repo-governance/).

## Diátaxis

Organize documentation by the reader's need, using the four Diátaxis forms:

| Form         | Reader need       | Use for                            | Location                                |
| ------------ | ----------------- | ---------------------------------- | --------------------------------------- |
| Tutorial     | Learning          | A guided first experience          | [`tutorials/`](tutorials/README.md)     |
| How-to guide | Completing a task | Goal-oriented directions           | [`how-to/`](how-to/README.md)           |
| Reference    | Looking up facts  | Accurate conventions and facts     | [`reference/`](reference/README.md)     |
| Explanation  | Understanding     | Rationale, trade-offs, and context | [`explanation/`](explanation/README.md) |

For example, a lesson that teaches a problem-solving pattern is a tutorial;
instructions for recording a drill are a how-to guide; an exercise directory
convention is reference; and the reasoning behind hands-on practice is
explanation.

Each section includes a README that defines its purpose. Add new documents to
the section that best answers the reader's question, and link them from this
index as the collection grows.

## Available Guides

- [Run the Nx Workspace](how-to/run-nx-workspace.md) explains how to build,
  test, and run the workspace's TypeScript demonstration app.

Learn more in the official [Diátaxis overview](https://diataxis.fr/) and its
[four-part map](https://diataxis.fr/map/).
