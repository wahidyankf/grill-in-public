---
tldr: "Covers the Knowledge Capture phase, archival to done/, and resuming a stopped plan."
when_to_use: "Use when a plan's delivery phases are complete, or when resuming an interrupted plan."
---

# Finalization

## Knowledge Capture

The last phase before archival triages `learnings.md`. Every entry is routed to exactly one durable home — a governance rule, a document, a subagent instruction, code, a test, or a new two-pager — or discarded with a one-line reason. Both safety checks run first: no secret leaves the plan, and the lesson generalizes beyond this one incident. The [knowledge capture rules](../../conventions/plans-organization-policy/knowledge-capture.md) hold the full routing.

Routing a rule into `repo-governance/` runs the [rules-propagation](../rules-propagation.md) workflow, and touching a harness path runs [harness-alignment](../harness-alignment.md) as well.

Archival is blocked until every entry is terminal, or the plan records `No generalizable learnings — <reason>`.

## Archival

1. Confirm every checkbox is ticked and every gate passed.
2. Rename the folder to `YYYY-MM-DD__<identifier>` with the completion date.
3. Move it to `plans/done/`.
4. Update `plans/in-progress/README.md` and `plans/done/README.md`.
5. Commit and push the move.

## Resuming

A plan resumed in a new session starts from the repository, not from memory:

1. Read `delivery.md` and find the first unticked item.
2. Check the Git log and the working tree against the ticked items above it. A ticked item whose change does not exist is untricked and re-run.
3. Rebuild the harness task list for the current phase only.
4. Re-run the previous phase's gate before continuing, so execution resumes from a state proven green rather than assumed green.

## Reopening

If a defect appears after archival, move the folder back to `in-progress/`, strip the date prefix, add a dated note in `README.md` explaining what broke, and execute the fix as a new phase. Editing a `done/` plan in place erases the history that makes the archive worth keeping.
