---
tldr: "Fixes plan folder names per stage, including the done/ completion-date prefix."
when_to_use: "Use when naming a plan folder or moving one between stages."
---

# Plan Naming

A plan folder's name depends on its stage. Only `done/` carries a date.

## backlog/ and in-progress/

```text
plans/backlog/<identifier>/
plans/in-progress/<identifier>/
```

Neither stage carries a date prefix, so promoting a plan from `backlog/` to `in-progress/` is a pure move with no rename. A date on an unfinished plan records when someone typed, not when anything shipped, and it ages badly while the work sits.

## done/

```text
plans/done/YYYY-MM-DD__<identifier>/
```

The date is the completion date — the day the final commit landed — not the day the plan was written. A double underscore separates it from the identifier.

## Identifier Rules

- Kebab-case: lowercase letters, digits, and hyphens only.
- No spaces, underscores, or capitals inside the identifier.
- ISO 8601 for the date, and only in `done/`.
- Describe the change, not the ticket: `wahidyankf-www-migration`, not `plan-3`.

## Examples

```text
good:  plans/backlog/wahidyankf-www-migration/
good:  plans/in-progress/wahidyankf-www-migration/
good:  plans/done/2026-08-18__wahidyankf-www-migration/
bad:   plans/backlog/2026-08-18__wahidyankf-www-migration/   date before completion
bad:   plans/done/2026-08-18_wahidyankf_www_migration/       single underscore, underscores
```

## Ideas

An idea is a file, not a folder: `plans/ideas/<slug>.md`, kebab-case, no date. It gains a folder only when promoted to `backlog/`.
