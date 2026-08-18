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

The [rules-quality-gate](rules-quality-gate.md) workflow runs this one as a step rather than restating it, so the five cases below have a single implementation. Running this workflow directly is still correct when only a harness changed.

## Prerequisites

Run `npm install` so the validation commands work.

## Steps

1. Inventory the instruction files, harness configs, subagents, and skills:

   ```sh
   rg --files -g 'AGENTS.md' -g 'CLAUDE.md' -g 'GEMINI.md' -g 'COPILOT.md' -g '.cursorrules' -g 'SKILL.md' -g '!node_modules'
   ls .claude/agents .codex/agents .opencode/agents .claude/skills .agents/skills .opencode/skills
   ```

2. Read `AGENTS.md` first, then each derivative. Classify every rule, command, path, and link in a derivative as equal, contradiction, duplication, orphan, or gap, per the [finding taxonomy](rules-quality-gate/02-finding-taxonomy.md). Leave what is equal; replace duplication with a link; correct or delete an orphan; add a gap only to the harness that needs it. Resolve a contradiction at the canonical source with the [Rules Propagation](rules-propagation.md) workflow, then correct the derivative.

3. Verify every command quoted in an instruction file exists in `package.json` or a `project.json` target, and that every referenced path exists.

4. Confirm each derivative names `AGENTS.md` as authoritative and links to it.

5. Compare the shared subagents, skills, and commands as the [harness capability parity policy](../conventions/harness-capability-parity-policy.md) requires. Each description must say what the agent does and when to use it.

6. Confirm each project config holds only documented settings, and that each directory is indexed as the [documentation index policy](../documentation-index-policy.md) requires, exemptions included.

7. Apply the edits to every affected instruction file, harness config, README, and subagent in the same change.

## Verification

```sh
npm run format:check
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
```

Every governed document must stay within the 500-word limit; split one rather than trim a required rule.

[Workspace commands](../development/workspace-commands.md#repository-checks) records the caveats of running these checks locally.

## Recovery

If a difference is substantive and its resolution unclear, leave both files unchanged and report the conflicting text, its effect, and a recommended resolution to the owner.
