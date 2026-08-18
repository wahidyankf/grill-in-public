---
tldr: "Requires a granular task list for multi-step work, updated as the work happens."
when_to_use: "Use when planning, executing, or reviewing any change that takes more than one verifiable step."
---

# Task Tracking Policy

## Scope

This policy covers the task list kept while work is in progress: when one is required, how small its items must be, and when it must be updated. It applies to every supported harness; see the [agent harness support policy](agent-harness-support.md). Each harness names the feature differently, task list or todo list, and the requirement is the same for all of them.

## When a List Is Required

Keep a task list for any change that takes more than one verifiable step, including a governance change, a multi-file edit, and an investigation whose findings drive later edits. A single edit with a single check needs no list.

## Granularity

Write one item per outcome that can be checked on its own. An item is too coarse when judging it done requires accepting several separate claims at once, and a plan step that names two verbs usually hides two items.

```text
too coarse:  "Add the policy and wire it everywhere and run the checks"
granular:    "Write the policy document"
             "Link it from AGENTS.md and the category README"
             "Run the verification gates"
```

Prefer the smaller split when unsure. A list that is too fine costs a line of output; a list that is too coarse hides how much work remains.

## Keeping It in Sync

The list must describe the present, not a plan written once and abandoned:

- Mark an item in progress before its first action, not after it succeeds.
- Mark it completed only when its outcome holds and has been verified. A failing gate leaves the item in progress.
- Record work discovered mid-task as new items instead of widening an existing one, so the count reflects the true remaining scope.
- Update the list as each item resolves. Marking several items complete in one batch at the end reports a state that was never observed.

## Why

The owner reads the list to know what is done, what is left, and what went wrong, and cannot see the reasoning behind it. A stale or coarse list therefore misreports the work rather than merely describing it briefly. Granular items also make an interrupted session resumable, because the first unfinished item states exactly where to restart.

## Verification

No automated gate can read a harness task list, since it is session state rather than repository content. This policy is verified in review: the list is compared against the change and the commands actually run. Announcements that a rule change occurred are separate; see the [rule change trigger policy](../development/rule-change-trigger-policy.md).
