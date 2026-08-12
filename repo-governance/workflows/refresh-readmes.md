---
tldr: "Reviews affected human-facing READMEs for accuracy before a thematic commit."
when_to_use: "Use when a change affects repository, project, documentation, or governance discoverability."
---

# Refresh READMEs

## Purpose

Keep human-facing READMEs accurate as the repository changes. Review the root `README.md` and every existing README below `apps/`, `libs/`, `docs/`, and `repo-governance/` recursively; update the smallest affected set in the same thematic commit as the change it explains.

## When to Use

Run this workflow before committing a change that alters purpose, directory layout, public behavior, commands, setup, dependencies, documentation navigation, governance, or contributor expectations. Also run it when adding, moving, or deleting a README or its indexed Markdown content.

## Steps

1. Inventory the review surface and the source of truth. Inspect changed, staged, and untracked files alongside the current README set:

   ```sh
   git diff --name-status origin/main...HEAD
   git diff --name-status
   git diff --cached --name-status
   printf '%s\n' README.md
   rg --files apps libs docs repo-governance -g 'README.md' | sort
   ```

2. Map each changed behavior to its human entry point. The root README explains the workspace; app and library READMEs explain each project's purpose, use, and contracts; `docs/` READMEs navigate Diátaxis content; and `repo-governance/` READMEs explain shared rules. Create a project README when a new runnable or reusable project needs a human starting point.

3. Read each affected README as a newcomer. Confirm names, commands, versions, paths, links, prerequisites, outputs, and claims match the implementation. Link to canonical detailed documents instead of copying rules. Keep agent instructions in `AGENTS.md`-style files; repository governance applies to both audiences.

4. For `docs/` and `repo-governance/`, follow the [documentation index policy](../documentation-index-policy.md): every directory has a README, each README indexes its immediate Markdown files and child directories, and child READMEs own deeper descendants.

5. Update only stale or missing material. Do not rewrite accurate prose, invent behavior, or mix unrelated documentation cleanup into the commit.

## Verification

Run the relevant commands and inspect the rendered Markdown links:

```sh
npm run format:check
npm run check:markdown-links
npm run check:governance
```

Confirm every changed human-facing behavior is discoverable from the appropriate README and that no README promises behavior absent from the repository.

## Recovery

If code, policy, or intended audience is ambiguous, stop and ask the owner; do not make a README authoritative by guessing. Repair the source of truth first when it is wrong, then refresh the README and its affected indexes.
