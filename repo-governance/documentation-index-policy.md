---
tldr: "Defines recursive README indexes and concise discovery metadata."
when_to_use: "Use when adding, moving, renaming, or reviewing Markdown in docs/ or repo-governance/."
---

# Documentation Index Policy

## Scope

Every directory in `docs/` and `repo-governance/`, at every depth, must contain
a `README.md`. The README is the directory's concise entry point for both people
and agents.

## Required Indexing

Each README must register its immediate Markdown documents, excluding itself,
with a descriptive relative link. It must also register each immediate child
directory through that directory's README. Child READMEs own their descendants;
do not repeat the full recursive tree.

Every Markdown document in scope, except `repo-governance/README.md`, must begin
with YAML frontmatter containing `tldr` and `when_to_use`. Keep both values
short enough to help a reader decide whether to open the document.

`repo-governance/README.md` is the exception because it is the governance entry
index. Every link there must state both the linked document's short description
and when a reader should use it.

## Maintenance

Create the README with a new directory. When adding, moving, renaming, or
removing Markdown content or a child directory, update every affected parent
README in the same change. Keep entries brief and describe when the reader
should open the linked document.

This policy implements [progressive disclosure](principles/progressive-disclosure.md):
indexes make focused documents discoverable without turning a parent README into
a duplicate manual.
