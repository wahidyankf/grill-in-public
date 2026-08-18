---
tldr: "Maps delivery checkboxes to harness tasks one to one and keeps both in step."
when_to_use: "Use when starting execution or resuming it after an interruption."
---

# Task List Sync

`delivery.md` is the durable record; the harness task list is what the owner watches while the work runs. They must agree, because two disagreeing records mean neither can be trusted.

## One Task per Checkbox

Materialize the current phase's checkboxes as tasks, one to one, in checklist order. Do not batch several checkboxes into one task: that hides progress, which is the failure the [task tracking policy](../../conventions/task-tracking-policy.md) exists to prevent.

Materialize one phase at a time. A task list holding every phase of a long plan reports a backlog rather than the work in hand.

## The Sync Ritual

For each item, in this order:

1. Mark the task in progress before the first action, not after it succeeds.
2. Do the work.
3. Verify it with the command the checkbox names.
4. Tick the checkbox in `delivery.md`, adding a one-line implementation note when the outcome differed from the plan.
5. Mark the task completed.

A failing verification leaves both the task and the checkbox open. An item that failed and was worked around is not complete; it is a new item plus a note.

## Discovered Work

Work found mid-phase becomes a new checkbox and a new task, not a wider existing one. Widening an item hides how much remains and makes the checklist claim progress it has not made.

## Truth on Disk

When the task list and `delivery.md` disagree — after a crash, or on resuming a session — the repository is the authority. Read what is actually on disk and in the Git log, rebuild the task list from that, and correct any checkbox that was ticked without its change existing.
