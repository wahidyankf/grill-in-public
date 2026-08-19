---
tldr: "Compares the shared subagents, skills, and commands, then each project config and directory index, and states what parity does not prove."
when_to_use: "Use after classifying derivative text, to check capability parity, harness configs, and READMEs."
---

# Capability and Config Parity

## Shared Capabilities

Compare the shared subagents, skills, and commands as the [harness capability parity policy](../../conventions/harness-capability-parity-policy.md) requires. Each description must take the two-part form the [agent harness support policy](../../conventions/agent-harness-support.md) requires.

## Configs and Indexes

Confirm each project config holds only documented settings, and that each directory is indexed as the [documentation index policy](../../documentation-index-policy.md) requires, exemptions included.

## What Parity Does Not Prove

That proof is narrower than it looks. The corpus has two mechanisms for keeping a copy in step with its source: the parity check, which shows the three harness copies equal each other but never that they match the policy they implement, and the sentence in a workflow requiring its prompt copy to change with it, which a person reviews. Neither reaches a sentence in one document that describes what another agent checks, so the [fixer discipline](../rules-quality-gate/04-fixer-discipline.md) requires reading the prompt before writing such a sentence.
