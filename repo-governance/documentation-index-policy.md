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

## Maintenance

Create the README with a new directory. When adding, moving, renaming, or
removing Markdown content or a child directory, update every affected parent
README in the same change. Keep entries brief and describe when the reader
should open the linked document.

This policy implements [progressive disclosure](principles/progressive-disclosure.md):
indexes make focused documents discoverable without turning a parent README into
a duplicate manual.
