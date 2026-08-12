# Badak Mini

Badak Mini is a deliberately small Go governance CLI. Its name means
“rhinoceros” in Indonesian, and its command grammar follows the relevant slice
of [rhino-cli](https://github.com/wahidyankf/ose-public/tree/main/apps/rhino-cli)
without porting Rhino's broader repository-management surface.

## Current Command

```sh
badak-mini harness instruction-size validate
badak-mini harness markdown-links validate
```

The command finds the Git repository root and ensures that `AGENTS.md` plus
every recursive Markdown file in `repo-governance/` contains at most 500 words.
It ignores non-Markdown files and reports each violation with a
progressive-disclosure remediation.

The Markdown-link command scans every Git-tracked repository Markdown file. It
validates local file targets and Markdown heading fragments, including
reference-style links. Targets must also be Git-tracked. It does not check
external URLs. The pre-push hook always runs this command so deleting or moving
a document cannot leave a dangling local link.

## Run and Verify

From the repository root:

```sh
npm run check:governance
npm run check:markdown-links
npx nx run badak-mini:build
npx nx run badak-mini:test:quick
```

Badak Mini is a standard-library-only Go module pinned to Go 1.26.1. It is an
intentional replacement for the former shell governance checker, not a general
Rhino CLI port.
