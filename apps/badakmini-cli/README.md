# Badak Mini

Badak Mini is a deliberately small Go governance CLI. Its name means “rhinoceros” in Indonesian, and its command grammar follows the relevant slice of [rhino-cli](https://github.com/wahidyankf/ose-public/tree/main/apps/rhino-cli) without porting Rhino's broader repository-management surface.

## Current Command

```sh
badak-mini harness instruction-size validate
badak-mini harness markdown-links validate
```

The command finds the Git repository root and ensures that the root agent instruction files, `AGENTS.md` and `CLAUDE.md`, and every recursive Markdown file in `repo-governance/` contain at most 500 words. A missing instruction file fails the check. Its `harness` command group names the family of harness-related checks, not the files they read. It ignores non-Markdown files and reports each violation with a progressive-disclosure remediation.

The Markdown-link command scans every Git-tracked repository Markdown file. It validates local file targets and Markdown heading fragments, including reference-style links. It does not check external URLs. The pre-push hook always runs this command so deleting or moving a document cannot leave a dangling local link.

## Run and Verify

From the repository root:

```sh
npm run check:governance
npm run check:markdown-links
npm run check:go-vulnerabilities
npx nx run badakmini-cli:build
npx nx run badakmini-cli:test:quick
```

Badak Mini's application code is standard-library-only and pinned to Go 1.26.5. Its development module also pins golangci-lint and govulncheck; run them through `go -C apps/badakmini-cli tool <name>`. It is an intentional replacement for the former shell governance checker, not a general Rhino CLI port.
