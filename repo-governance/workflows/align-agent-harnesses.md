---
tldr: "Verifies every agent harness file stays equal to and consistent with AGENTS.md."
when_to_use: "Use after changing AGENTS.md, governance, tooling, or any assistant-specific instruction file."
---

# Align Agent Harnesses

## Purpose

Keep every agent harness equal in effect, so any supported tool receives the same rules, commands, and structure. That covers the instruction files, `AGENTS.md` and each derivative such as `CLAUDE.md`, plus each tool's project config and shared subagents. The standing rules live in the [agent instruction alignment policy](../conventions/agent-instruction-alignment-policy.md) and the [agent harness support policy](../conventions/agent-harness-support.md); this workflow is the procedure that proves they hold.

## When to Use

Use it after changing `AGENTS.md`, a document under `repo-governance/`, a harness file, or repository tooling that a harness file describes, such as an npm script, an Nx target, or a Git hook. Use it also before a thematic commit that touches any of those.

## Prerequisites

Run `npm install` so the validation commands work.

## Steps

1. Inventory the harness files, configs, and subagents:

   ```sh
   rg --files -g 'AGENTS.md' -g 'CLAUDE.md' -g 'GEMINI.md' -g 'COPILOT.md' -g '.cursorrules' -g '!node_modules'
   ls .claude/agents .codex/agents .opencode/agents
   ```

2. Read `AGENTS.md` first, then each derivative. For every rule, command, path, and link in a derivative, decide which case applies:
   - **Equal** — it matches canonical guidance. Leave it.
   - **Contradiction** — it differs in requirement, scope, or verification. Resolve it at the canonical source with the [Propagate Rules](propagate-rules.md) workflow, then correct the derivative.
   - **Duplication** — it restates a canonical rule in words that can drift. Replace it with a link.
   - **Orphan** — it describes tooling, a path, or a policy that no longer exists. Delete or correct it.
   - **Gap** — it is genuine tool-specific operational detail missing from that harness. Add it there only.

3. Verify every command quoted in a harness file still exists in `package.json` or a `project.json` target, and that every referenced path exists.

4. Confirm each derivative names `AGENTS.md` as authoritative and links to it.

5. Compare the shared subagents across `.claude/agents/`, `.codex/agents/`, and `.opencode/agents/`. Each role must exist everywhere with the same name, purpose, and permission posture; wording may differ per format, but capability must not. Confirm each project config still holds only documented settings, and no rules.

6. Apply the edits to every affected harness file, config, and subagent in the same change.

## Verification

```sh
npm run format:check
npm run check:governance
npm run check:markdown-links
```

`AGENTS.md`, `CLAUDE.md`, and every `repo-governance/` document must stay within the 500-word limit; move detail into a focused governance document rather than trimming a required rule.

The link check reads Git-tracked files, so a newly created document is invisible to it. Run `git add -N <file>` for each new Markdown file before trusting a local link run.

## Recovery

If a difference is substantive and its correct resolution is unclear, leave both files unchanged, and report the conflicting text, its practical effect, and a recommended resolution to the repository owner.
