---
tldr: "Lists what to read before authoring a plan or questioning the owner."
when_to_use: "Use at the start of plan-planning, before any question reaches the owner."
---

# Exploration

Exploration comes first because a plan written from assumptions encodes those assumptions as checklist items, where they are expensive to remove.

## Read the Repository

- `AGENTS.md` and the governance documents the change touches. A plan that contradicts a rule is a plan that fails its own gate.
- The projects named in the prompt: their `project.json` targets, their tests, and their READMEs.
- The Git history around the files in question. A change that was tried and reverted is worth knowing before proposing it again.
- Any prior plan in `plans/done/` covering the same area, which records what actually happened last time.

## Establish the Facts a Plan Needs

Before writing anything, be able to state:

- which projects and files the change touches, by path
- which commands verify each affected project, verbatim
- what currently passes, so a later failure can be attributed to the work
- what the change depends on that does not exist yet

## Name the Unknowns

List every question the repository could not answer. These become the grilling agenda in [02 Grilling](02-grilling.md). Distinguish two kinds: a fact that more reading would settle, which is still exploration, and a decision only the owner can make, which is a question.

## Record What Was Read

The plan's `tech-docs.md` states what was inspected and what it showed. A reader who disagrees with a design decision can then check the evidence rather than relitigate the reasoning.
