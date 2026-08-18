# Badak Mini

Badak Mini is a deliberately small Go governance CLI. Its name means “rhinoceros” in Indonesian, and its command grammar follows the relevant slice of [rhino-cli](https://github.com/wahidyankf/ose-public/tree/main/apps/rhino-cli) without porting Rhino's broader repository-management surface.

## Current Command

```sh
badak-mini harness instruction-size validate
badak-mini harness markdown-links validate
badak-mini harness rule-change validate
badak-mini harness rule-change hook
badak-mini harness capability-parity validate
badak-mini harness project-targets validate
```

The command finds the Git repository root and ensures that the root agent instruction files, `AGENTS.md` and `CLAUDE.md`, every recursive Markdown file in `repo-governance/`, and every recursive `README.md` in `.agents/`, `.claude/`, `.codex/`, and `.opencode/` contain at most 500 words. A missing instruction file fails the check, while an absent harness directory is skipped. Agent and command definitions are prompts rather than indexes, so they are not measured. Its `harness` command group names the family of harness-related checks, not the files they read. It ignores non-Markdown files and reports each violation with a progressive-disclosure remediation.

The Markdown-link command scans every Git-tracked repository Markdown file. It validates local file targets and Markdown heading fragments, including reference-style links. It does not check external URLs. The pre-push hook always runs this command so deleting or moving a document cannot leave a dangling local link.

The rule-change commands announce the [Rules Propagation](../../repo-governance/workflows/rules-propagation.md) workflow when a change touches a rule path, and the [Harness Alignment](../../repo-governance/workflows/harness-alignment.md) workflow as well when that path is one a harness reads. The `validate` form reads the staged paths and runs during pre-commit; the `hook` form reads a harness pre-edit payload on stdin, in either the file-path or the `apply_patch` shape, and answers with the notice as additional context. Both stay silent for ordinary work and always exit zero, so neither can block an edit or a commit. See the [rule change trigger policy](../../repo-governance/development/rule-change-trigger-policy.md).

The capability-parity command compares the subagents, skills, and commands each harness exposes and fails when one harness lacks an entry its peers have. It skips a capability no harness uses and exempts a harness that cannot load one. See the [harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).

The project-targets command reads every `project.json` outside vendored and built directories and fails when a project omits `test:quick`, `lint`, or `typecheck`, does not reach the last two from `test:quick` through `dependsOn`, or turns off the cache pre-push relies on. A `dependsOn` entry written as an object, or prefixed with `^`, runs a target elsewhere and does not satisfy the requirement. See the [testing policy](../../repo-governance/development/testing-policy.md).

## Run and Verify

From the repository root:

```sh
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
npm run check:project-targets
npm run check:go-vulnerabilities
npx nx run badakmini-cli:build
npx nx run badakmini-cli:test:quick
```

Badak Mini's application code is standard-library-only and pinned to Go 1.26.5. Its development module also pins golangci-lint and govulncheck; run them through `go -C apps/badakmini-cli tool <name>`. It is an intentional replacement for the former shell governance checker, not a general Rhino CLI port.
