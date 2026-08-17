---
tldr: "Verifies every agent harness file stays equal to and consistent with AGENTS.md."
when_to_use: "Use after changing AGENTS.md, governance, tooling, or any assistant-specific instruction file."
---

# Align Agent Harnesses

## Purpose

Keep every agent harness file — `AGENTS.md` and each assistant-specific derivative such as `CLAUDE.md` — equal in effect, so any agent reading any one of them receives the same rules, commands, and structure. The standing rules live in the [agent instruction alignment policy](../conventions/agent-instruction-alignment-policy.md); this workflow is the procedure that proves they hold.

## When to Use

Use it after changing `AGENTS.md`, a document under `repo-governance/`, a harness file, or repository tooling that a harness file describes, such as an npm script, an Nx target, or a Git hook. Use it also before a thematic commit that touches any of those.

## Prerequisites

Install dependencies with `npm install` so the formatting and validation commands run.

## Steps

1. Inventory the harness files:

   ```sh
   rg --files -g 'AGENTS.md' -g 'CLAUDE.md' -g 'GEMINI.md' -g 'COPILOT.md' -g '.cursorrules' -g '!node_modules'
   ```

2. Read `AGENTS.md` first, then each derivative. For every rule, command, path, and link in a derivative, decide which case applies:
   - **Equal** — it matches canonical guidance. Leave it.
   - **Contradiction** — it differs in requirement, scope, or verification. Resolve it at the canonical source with the [Propagate Rules](propagate-rules.md) workflow, then correct the derivative.
   - **Duplication** — it restates a canonical rule in words that can drift. Replace it with a link.
   - **Orphan** — it describes tooling, a path, or a policy that no longer exists. Delete or correct it.
   - **Gap** — it is genuine tool-specific operational detail missing from that harness. Add it there only.

3. Verify every command quoted in a harness file still exists in `package.json` or a `project.json` target, and that every referenced path exists.

4. Confirm each derivative names `AGENTS.md` as authoritative and links to it.

5. Apply the edits to all affected harness files in the same change.

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
