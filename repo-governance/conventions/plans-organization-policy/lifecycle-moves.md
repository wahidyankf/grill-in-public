---
tldr: "Defines how a plan is promoted, started, completed, archived, and reopened."
when_to_use: "Use when moving a plan between plans/ stages."
---

# Lifecycle Moves

A plan moves in one direction until it is done, and each move is a committed change that updates both indexes.

```text
ideas/ --promote--> backlog/ --start--> in-progress/ --complete--> done/
                                            ^                        |
                                            +--------reopen----------+
```

## Promoting an Idea

A ripe two-pager becomes a plan when its open questions are answered. Create `plans/backlog/<identifier>/`, author the five documents through the [plan-planning](../../workflows/plan-planning.md) workflow, delete the idea file, and update both indexes. The idea's prior art and non-goals carry into `brd.md` rather than being rewritten from scratch.

## Starting Work

1. Move the folder from `backlog/<identifier>/` to `in-progress/<identifier>/`. No rename: neither stage carries a date.
2. Update `backlog/README.md` and `in-progress/README.md`.
3. Commit and push the move before executing any checklist item, so the repository states what is active before it changes.

A plan is never executed out of `backlog/`.

## Completing Work

1. Confirm every checklist item is ticked and every phase gate passed.
2. Run the Knowledge Capture phase; see [knowledge capture](knowledge-capture.md).
3. Rename the folder to `YYYY-MM-DD__<identifier>` using the completion date and move it to `done/`.
4. Update `in-progress/README.md` and `done/README.md`.
5. Commit the move with a message naming the plan.

## Checkbox Lockstep

Tick a checkbox only after the change it describes exists. Ticking ahead of the work turns the checklist into a wish and makes a resumed session trust a state that was never reached.

## Reopening

If a defect surfaces after archival, move the folder back to `in-progress/`, strip the date prefix, and add a dated note in `README.md` stating what broke. A reopened plan is honest history; a quietly edited `done/` plan is not.

No `plan-checker` rule restates any move above, so nothing here has a harness copy to keep in step. That is checked rather than overlooked: a rule that gains one gets the notice its siblings carry.
