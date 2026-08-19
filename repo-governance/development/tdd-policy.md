---
tldr: "Requires red-green-refactor cycles bound one-to-one to Gherkin scenarios."
when_to_use: "Use when writing a delivery checklist or implementing any behavior change."
---

# TDD Policy

## Scope

This policy covers how behavior gets implemented: test first, one scenario at a time. It applies to every change in `apps/` and `libs/`, including Badak Mini, whose checks silently disable a repository gate when they are wrong.

## The Cycle

Each behavior is delivered as a RED → GREEN → REFACTOR cycle:

1. **RED** — write the test that expresses the scenario, and run it. It must fail, and it must fail for the stated reason. A test that passes on first run tested nothing.
2. **GREEN** — write the smallest change that makes it pass. Not the elegant version; the passing version.
3. **REFACTOR** — improve the code with the test still passing. This step is where the design happens, and skipping it is how a suite of passing tests accumulates an unmaintainable implementation.

## One Scenario per Cycle

Every behavior cycle targets exactly one Gherkin scenario from the [specs policy](specs-policy.md). Its RED step names the scenario and inlines the scenario verbatim as a fenced `gherkin` block, so the executor never has to open another file to know what to assert.

```text
- [ ] [AI] RED — **Gherkin (binds)** "The app greets the configured name". Add the failing
      test to `apps/dummy-app/src/index.test.ts`. Verify with
      `npx nx run dummy-app:test:quick` — the new test fails. Scenario:

~~~gherkin
Scenario: The app greets the configured name
  Given the app is configured with the name "Wahidyan"
  When the app runs
  Then the output is "Hello, Wahidyan!"
~~~
```

The scenario fence sits at the left margin rather than inside the list item, because a fence indented into a list is re-indented on every formatting pass and never settles.

Bundling several scenarios into one cycle hides which behavior a failure belongs to. Long checklists are the expected outcome and are not a reason to merge cycles.

Pure data or calculation tests that underpin several scenarios use `**Gherkin (underpins)**` and may name a list.

Every `plan-checker` prompt states this rule in the imperative, because a subagent prompt has to stand alone. Change it in the same edit, in all three harness copies.

## Regression Tests

Every bug fix begins with a test that reproduces the bug and fails. A fix without one asserts that the bug is gone without ever having shown that it was there.

## Verification

`test:quick` and `test:integration` run the resulting tests, per the [testing policy](testing-policy.md). `plan-checker` flags a RED step that inlines no scenario or names more than one, and a `prd.md` scenario that no RED step inlines, because its rule 4 runs in both directions. Behavior that reaches neither document escapes it, and whether a RED step truly failed first holds by review.
