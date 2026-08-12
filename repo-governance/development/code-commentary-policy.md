---
tldr: "Requires purposeful comments that teach intent, flow, and non-obvious decisions."
when_to_use: "Use when adding or reviewing executable source, tests, or repository scripts."
---

# Code Commentary Policy

## Purpose

This is a learning repository. Code should help a reader understand both what it
does and the reasoning behind it, especially during interview preparation.

## Required Commentary

Add purposeful comments to executable source, tests, and repository scripts.
Explain decisions that are not obvious from syntax alone, including:

- the intent and boundary of a module, function, or test;
- important control-flow stages and data transformations;
- invariants, security or correctness checks, and failure behavior;
- why a dependency is injected, mocked, cached, or intentionally avoided; and
- non-obvious shell, Git, parsing, regular-expression, or library behavior.

## Lint Enforcement

Each project owns the commentary checks in its `lint` target. TypeScript uses
ESLint and `eslint-plugin-jsdoc` to require complete-sentence summaries for
named executable declarations; Go uses golangci-lint with Revive to document
exported declarations. Both require a specific, explained linter suppression.

These checks establish a minimum shape, not comment quality. They do not decide
whether a private helper's reasoning is worth explaining or whether a summary
is true, useful, and current.

## What to Avoid

Do not narrate self-evident statements or duplicate a precise name. Prefer a
short comment immediately before the decision it explains. Keep comments true
when code changes; stale explanations are defects.

## Review

When adding or changing executable code, manually review the surrounding flow
for the comments a learner would need to reconstruct the reasoning. A passing
linter never replaces that review. Keep configuration, generated artifacts,
lockfiles, and plain data free of explanatory noise unless their format supports
and benefits from comments.
