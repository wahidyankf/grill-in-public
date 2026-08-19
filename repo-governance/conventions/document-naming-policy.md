---
tldr: "Fixes how Markdown documents and their child directories are named."
when_to_use: "Use when creating, renaming, or splitting any Markdown document under docs/, repo-governance/, plans/, or specs/."
---

# Document Naming Policy

## Scope

This policy governs the filenames of Markdown documents in `docs/`, `repo-governance/`, `plans/`, and `specs/`, and the names of the child directories a split document creates. Code identifiers are a separate concern and belong to the [code style policy](../development/code-style-policy.md). Plan folder names are stage-dependent and belong to the [plans organization policy](plans-organization-policy/plan-naming.md).

## Rules

**Lowercase and hyphens.** Use lowercase letters, digits, and hyphens. No spaces, underscores, or capitals. The filename is a path a link has to reproduce exactly, so a name that needs escaping or shifting is a name that gets mistyped.

**Name the subject, not the action.** A document is named for what it covers: `markdown-style-policy.md`, `commit-hook-policy.md`, `folder-structure.md`. A workflow is named the same way, as a domain-prefixed noun phrase — `rules-propagation.md`, `harness-alignment.md`, `plan-quality-gate.md` — not as an imperative such as `propagate-rules.md`. The noun form sorts into families as a directory grows, and the family prefix is what makes six workflows readable at a glance.

**Number children only when they are steps.** A document split into a directory of children names each child for what it covers. Add a `NN-` prefix when the children are performed in order; there the number is part of the rule and skipping ahead is an error. Every child directory under `repo-governance/workflows/` is numbered, without exception: a workflow is a procedure, so its children are its steps, and the number states where each one falls. A child that reads like reference material is still consulted at a point in the run, so it takes the number of that point rather than escaping the sequence. Do not number children that are consulted individually, such as the rules behind a convention: a number asserts a sequence that does not exist, and it forces a renumber every time a rule is added. The directory's README carries whatever order helps a reader.

**Match the directory to the document it splits.** A child directory takes its parent document's name without the extension, so `plans-organization-policy.md` splits into `plans-organization-policy/`. A reader who sees one can predict the other.

## Verification

No automated check reads a filename. This policy is verified in review, and by `npm run check:markdown-links`, which fails when a rename leaves a link pointing at the old name.
