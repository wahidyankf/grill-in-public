---
tldr: "Defines recursive README indexes and concise discovery metadata."
when_to_use: "Use when adding, moving, renaming, or reviewing Markdown in docs/, repo-governance/, cv/, scripts/, or a harness directory."
---

# Documentation Index Policy

## Scope

Every directory in `docs/`, `repo-governance/`, `cv/`, `scripts/`, and each agent harness directory — `.agents/`, `.claude/`, `.codex/`, and `.opencode/` — at every depth, must contain a `README.md`. The README is the directory's concise entry point for both people and agents.

A harness directory holds tool configuration rather than prose, so its README indexes what the directory contains and what each entry does. The `tldr` and `when_to_use` requirement below does not apply there, because those files carry the frontmatter their tool defines.

Two exemptions apply. A skill directory needs none: its `SKILL.md` names the skill and when to use it, and `skills/README.md` registers it. Second, a harness registers some directories by filename, so an index placed there becomes a command or an agent; add the README only where the tool ignores it or offers a flag that keeps it inert, and otherwise index that directory from its parent. The [agent harness support policy](conventions/agent-harness-support.md) records the verified behavior per directory; check it, and test the tool, first.

## Required Indexing

Each README must register its immediate Markdown documents, excluding itself, with a descriptive relative link. It must also register each immediate child directory through that directory's README. Child READMEs own their descendants; do not repeat the full recursive tree.

Every Markdown document under `docs/` or `repo-governance/`, except `repo-governance/README.md`, must begin with YAML frontmatter containing `tldr` and `when_to_use`. Keep both values short enough to help a reader decide whether to open the document.

`repo-governance/README.md` is the exception because it is the governance entry index.

Every `rules-checker` prompt states this indexing rule and this frontmatter rule, with its one exception, in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.

## When an Index Reaches the Word Limit

Never make an index fit by dropping entries or their descriptions. An incomplete index hides work, the failure this policy exists to prevent.

The [document word limit policy](conventions/document-word-limit-policy.md) states the limit, which documents it governs, and how a document that has reached it is fixed; an index is fixed the same way.

## Maintenance

Create the README with a new directory. When adding, moving, renaming, or removing Markdown content or a child directory, update every affected parent README in the same change. Keep each entry brief, and say when a reader should open what it links to.

This policy implements [progressive disclosure](principles/progressive-disclosure.md): indexes make focused documents discoverable without turning a parent README into a duplicate manual.
