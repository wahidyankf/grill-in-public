---
tldr: "Verifies every harness receives the same rules through its instruction file, config, and subagents."
when_to_use: "Use after changing AGENTS.md, governance, tooling, or any harness's instruction file, config, or subagents."
---

# Harness Alignment

## Purpose

Keep every harness equal in effect, so any supported tool receives the same rules: the instruction files, each tool's config, its READMEs, and its capabilities. The standing rules live in the [agent instruction alignment policy](../conventions/agent-instruction-alignment-policy.md) and the [agent harness support policy](../conventions/agent-harness-support.md); this workflow proves they hold.

## When to Use

Use it after changing `AGENTS.md`, a `repo-governance/` document, an instruction file, or tooling those files describe, such as an npm script or Git hook. Use it also before a thematic commit touching those.

## Automatic Triggers

A change to an instruction file, `opencode.json`, or a harness directory announces this workflow, at pre-commit and in a harness pre-edit hook. The [rule change trigger policy](../development/rule-change-trigger-policy.md) owns those paths and mechanisms. An announcement is not the work.

## Composition

The [rules-quality-gate](rules-quality-gate.md) workflow runs this one as a step rather than restating it, so the five cases it classifies have a single implementation. Running this workflow directly is still correct when only a harness changed.

## Prerequisites

Run `npm install` so the validation commands work.

## Steps

1. Inventory what every harness reads; see [inventory](harness-alignment/01-inventory.md).

2. Classify every difference between a derivative and `AGENTS.md`; see [derivative comparison](harness-alignment/02-derivative-comparison.md).

3. Verify the commands and paths that an instruction file quotes; same document.

4. Confirm each derivative's statement of authority; same document.

5. Compare the shared subagents, skills, and commands; see [capability and config parity](harness-alignment/03-capability-and-config-parity.md).

6. Confirm each project config and each directory index; same document.

7. Apply the edits to every affected instruction file, harness config, README, and subagent in the same change.

## Verification

```sh
npm run format:check
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
```

The [document word limit policy](../conventions/document-word-limit-policy.md) governs the limit every governed document lives under, and how a document that has reached it is fixed.

[Workspace commands](../development/workspace-commands.md#repository-checks) records the caveats of running these checks locally.

## Recovery

If a difference is substantive and its resolution unclear, leave both files unchanged and report the conflicting text, its effect, and a recommended resolution to the owner.
