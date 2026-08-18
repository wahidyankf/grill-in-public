---
tldr: "Requires Gherkin acceptance criteria in specs/, separate from the code that implements them."
when_to_use: "Use when adding or changing behavior in an app or library, or when writing a plan's prd.md."
---

# Specs Policy

## Scope

This policy covers `specs/`, the repository's description of what its software should do. Specs state intent and observable behavior; `apps/` and `libs/` state how that behavior is delivered. Keeping them apart means a behavior discussion can start from one source of truth instead of from an implementation.

## Structure

Specs mirror the workspace: `specs/apps/<name>/` and `specs/libs/<name>/`. Each carries the folders its subject actually needs, and no others:

```text
specs/apps/<name>/
├── README.md          index for this subject
├── product/           the problem, its users, and scope
├── system-context/    the boundary with the outside world
├── containers/        the running parts
├── components/        the internals of a part
└── behavior/          Gherkin acceptance scenarios
```

`behavior/` is the only mandatory folder. Add another when the subject is complex enough to need it, not in advance.

## Gherkin

Acceptance criteria are written as Gherkin scenarios in `.feature` files under `behavior/`.

A scenario uses exactly one primary `Given`, one `When`, and one `Then`; further steps chain with `And` or `But`. A `Background` block and a `Scenario Outline` `Examples` table are exempt. Two `When` steps in one scenario describe two behaviors, so split them.

```gherkin
Feature: Greeting

  Scenario: The app greets the configured name
    Given the app is configured with the name "Wahidyan"
    When the app runs
    Then the output is "Hello, Wahidyan!"
```

## Binding to Tests

A scenario is not documentation: it binds to a test that fails when the behavior breaks, as the [TDD policy](tdd-policy.md) requires. A scenario with no test behind it is worse than no scenario, because it claims coverage that does not exist.

## When Specs Are Required

A plan that adds or changes behavior in `apps/` or `libs/` writes the scenarios into its `prd.md` and lands them in `specs/` as part of delivery. Existing projects gain specs when a plan next touches them, not through a retrofit pass. Drills are exempt: a drill is practice, not repository behavior.

## Verification

No automated check reads `specs/` yet. `plan-checker` verifies that a plan's Gherkin follows the cardinality rule above and that each scenario names the test that binds it.
