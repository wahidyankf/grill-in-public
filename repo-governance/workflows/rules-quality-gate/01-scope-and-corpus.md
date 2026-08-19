---
tldr: "Defines which files the rules quality gate reads and how it treats each kind."
when_to_use: "Use when running the gate or deciding whether a file is in its corpus."
---

# Scope and Corpus

A rule is any sentence that tells a contributor or an agent what they must, may, or must not do. The gate reads wherever such a sentence can live, which is more places than the word-limit check governs.

## In Scope, Judged Fully

- `repo-governance/**` — every policy, principle, workflow, and index.
- `AGENTS.md` and `CLAUDE.md` — the instruction files.
- `.agents/`, `.claude/`, `.codex/`, `.opencode/` — subagent definitions, skills, commands, plugin notes, and each directory's README.
- Every `SKILL.md`, wherever it lives.

## In Scope, Judged Narrowly

- `docs/` — read only for two failures: a requirement stated there that belongs in `repo-governance/`, and a reference to a rule, path, or workflow that no longer exists under that name. The gate does not judge tutorial quality, Diátaxis fit, or prose style. A governance gate that starts reviewing how-to guides stops being one.

## Out of Scope

- `plans/` — governed by the [plan quality gate](../plan-quality-gate.md), which reads plans against the plans policy.
- `specs/` — behavior, not rules.
- Source code and tests — a rule expressed as code is verified by running it.

## Harness Directories

The gate reads harness files as rule documents, because a prompt tells an agent what to do. It does not repeat the sweep: the [Harness Alignment](../harness-alignment.md) workflow owns the inventory, the per-item comparison, the command and path verification, and the parity check, and the gate runs it as a step rather than doing that work twice. A gap the gate notices while reading is still its finding to report; the systematic search for gaps is alignment's.

## Unclassified Surfaces

Open finding, gate run 3: the sections above classify neither `cv/`, `scripts/`, nor the root `README.md`, though each carries rule sentences by the definition this document opens with, and `AGENTS.md` routes agents into `cv/README.md`. Classifying them either way changes what `rules-checker` covers, so the owner decides it. Run 3 found a copied rule and a stale reference on two of those surfaces.

## Prompts Are Rules

A subagent definition is a rule document: it tells an agent how to behave, and a contradiction between a prompt and a policy is a real defect rather than a stylistic difference. The gate treats prompts as rule-bearing, and `rules-fixer` may edit them — with the parity check as the proof it left the harnesses equal. What that proof does not reach is recorded in [capability and config parity](../harness-alignment/03-capability-and-config-parity.md).

Every `rules-checker` prompt states the corpus above and this judgment of a prompt in the imperative, because a subagent prompt has to stand alone. Change them in the same edit, in all three harness copies.
