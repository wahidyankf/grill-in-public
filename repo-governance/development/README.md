---
tldr: "Indexes development policies for code, testing, hooks, Nx, and validation."
when_to_use: "Use when changing executable code, development tooling, tests, or quality gates."
---

# Development Governance

This directory contains rules for building, testing, validating, and maintaining the repository's executable code. Read only the policy that matches the work at hand:

- [Badak Mini](badakmini-cli-policy.md) for repository-local validation checks.
- [Code Commentary](code-commentary-policy.md) for learning-oriented comments.
- [Code Style](code-style-policy.md) for the language target, naming, indentation, and import style.
- [Commit Hooks](commit-hook-policy.md) for required Git-hook behavior.
- [Harness Pre-Edit Triggers](harness-pre-edit-triggers.md) for what each harness wires before an edit, and how far it is verified.
- [Nx Workspace](nx-workspace-policy.md) for raw-Nx boundaries and verification.
- [Rule Change Triggers](rule-change-trigger-policy.md) for how a rule change announces the workflows that must follow it.
- [Specs](specs-policy.md) for Gherkin acceptance criteria and the `specs/` tree.
- [TDD](tdd-policy.md) for red-green-refactor cycles bound to scenarios.
- [Testing](testing-policy.md) for quick and integration-test responsibilities.
- [Workspace Commands](workspace-commands.md) for the canonical command, check, and hook reference.

Foundational principles remain in [`../principles/`](../principles/README.md), and repeatable procedures remain in [`../workflows/`](../workflows/README.md).
