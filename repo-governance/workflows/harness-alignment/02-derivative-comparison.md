---
tldr: "Classifies every rule, command, path, and link in a derivative, then checks its commands, paths, and authority."
when_to_use: "Use when comparing an instruction file, harness config, or subagent against AGENTS.md."
---

# Derivative Comparison

## Classification

Read `AGENTS.md` first, then each derivative. Classify every rule, command, path, and link in a derivative as equal, contradiction, duplication, orphan, or gap, per the [finding taxonomy](../rules-quality-gate/02-finding-taxonomy.md). Leave what is equal; replace duplication with a link; correct or delete an orphan; add a gap only to the harness that needs it. Resolve a contradiction at the canonical source with the [Rules Propagation](../rules-propagation.md) workflow, then correct the derivative.

## Commands and Paths

Verify every command quoted in an instruction file exists in `package.json` or a `project.json` target, and that every referenced path exists.

## Authority

Confirm each derivative names `AGENTS.md` as authoritative and links to it.
