---
tldr: "Defines how a rule change announces the workflows that must follow it, and what each harness can trigger."
when_to_use: "Use when changing the rule paths, the pre-commit announcement, or a harness pre-edit hook."
---

# Rule Change Trigger Policy

## Scope

This policy covers the automation that announces the [Rules Propagation](../workflows/rules-propagation.md) and [Harness Alignment](../workflows/harness-alignment.md) workflows. Each workflow owns what to do once announced.

## Rule Paths

`badak-mini harness rule-change` is the single definition of a rule path: `AGENTS.md`, `CLAUDE.md`, `opencode.json`, and anything under `repo-governance/`, `.claude/`, `.codex/`, `.opencode/`, `.agents/`, or `.husky/`. Change that list in one place, and add a test with it.

A harness path is the narrower set the tools read: the instruction files, `opencode.json`, and the harness directories. Only these can leave one harness unequal to another, so only these announce the align workflow, on top of the propagate workflow every rule path announces. Announcing both every time would teach readers to skip the second line.

## Guaranteed Trigger

The pre-commit hook runs `npm run check:rule-change`, which names the applicable workflows when a staged path carries rules. It runs for every editor, harness, and human, so the trigger never depends on which tool made the change.

It reports and exits zero. A hook can tell that the workflow applies; it cannot tell whether it was followed, and a gate that cannot judge its own condition only teaches people to bypass it. The mechanical parts stay enforced elsewhere, by the word limits and the link check.

## Harness Pre-Edit Triggers

[Harness pre-edit triggers](harness-pre-edit-triggers.md) records what each harness wires before an edit, and how far each one is verified.

Do not treat a harness trigger as the guarantee. Each one is a convenience over the pre-commit hook, and each can be switched off outside this repository.

## Verification

```sh
npm run check:rule-change
echo '{"tool_input":{"file_path":"AGENTS.md"}}' | go -C apps/badakmini-cli run ./cmd/badak-mini harness rule-change hook
```

The first prints nothing unless a staged path carries rules. The second prints the hook response for a rule path, and nothing for any other file.
