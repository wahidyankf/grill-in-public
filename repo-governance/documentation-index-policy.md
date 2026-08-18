---
tldr: "Defines recursive README indexes and concise discovery metadata."
when_to_use: "Use when adding, moving, renaming, or reviewing Markdown in docs/, repo-governance/, or a harness directory."
---

# Documentation Index Policy

## Scope

Every directory in `docs/`, `repo-governance/`, and each agent harness directory — `.agents/`, `.claude/`, `.codex/`, and `.opencode/` — at every depth, must contain a `README.md`. The README is the directory's concise entry point for both people and agents.

A harness directory holds tool configuration rather than prose, so its README indexes what the directory contains and what each entry does. The `tldr` and `when_to_use` requirement below does not apply there, because those files carry the frontmatter their tool defines.

Two exemptions apply. A skill directory needs none: its `SKILL.md` names the skill and when to use it, and `skills/README.md` registers it. Second, a harness registers some directories by filename, so an index placed there becomes a command or an agent. Add the README only where the tool ignores it or offers a flag that keeps it inert. Where neither holds, index that directory from its parent. The [agent harness support policy](conventions/agent-harness-support.md) records the verified behavior per directory; check it, and test the tool, before adding the file.

## Required Indexing

Each README must register its immediate Markdown documents, excluding itself, with a descriptive relative link. It must also register each immediate child directory through that directory's README. Child READMEs own their descendants; do not repeat the full recursive tree.

Every Markdown document under `docs/` or `repo-governance/`, except `repo-governance/README.md`, must begin with YAML frontmatter containing `tldr` and `when_to_use`. Keep both values short enough to help a reader decide whether to open the document.

`repo-governance/README.md` is the exception because it is the governance entry index. Every index entry there must state both the linked document's short description and when a reader should use it.

## When an Index Reaches the Word Limit

`npm run check:governance` holds `AGENTS.md`, `CLAUDE.md`, every `repo-governance/` document, and every harness directory README to 500 words. Agent and command definitions are prompts, not indexes, and stay unmeasured.

An index that reaches the limit is not too wordy; its directory has too many peers. Group them: create a subdirectory, move the related documents into it, give it its own README, and register that child from the parent, which then carries one line instead of many.

Never make an index fit by dropping entries or their descriptions. An incomplete index hides work, which is the failure this policy exists to prevent. Shorten an entry's wording only when it runs past one line, and split the directory once that is not enough.

## Maintenance

Create the README with a new directory. When adding, moving, renaming, or removing Markdown content or a child directory, update every affected parent README in the same change. Keep entries brief and describe when the reader should open the linked document.

This policy implements [progressive disclosure](principles/progressive-disclosure.md): indexes make focused documents discoverable without turning a parent README into a duplicate manual.
