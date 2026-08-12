# Propagate Rules

## Purpose

Integrate new or changed repository rules into the correct governance location
without duplication, ambiguity, or contradictions. Rules include guidance in
`AGENTS.md`, `repo-governance/`, and any agent- or skill-specific instruction
files added to the repository.

## When to Use

Use this workflow when introducing, changing, moving, or removing a rule for
contributors, coding agents, agent skills, validation, or repository workflows.

## Prerequisites

State the proposed rule in one sentence, including its scope, trigger, and
required behavior. Preserve the source or decision that justifies it when one
exists.

## Steps

1. Inventory applicable guidance before editing. Start with `AGENTS.md`,
   `repo-governance/`, and its `principles/`; also search for agent and skill
   files, for example:

   ```sh
   rg --files -g 'AGENTS.md' -g 'SKILL.md' -g 'CLAUDE.md' -g 'GEMINI.md' \
     -g 'COPILOT.md' -g '!node_modules'
   ```

2. Choose one canonical home:
   - Put universal, short requirements in root `AGENTS.md`.
   - Put detailed or conditional policy in `repo-governance/`.
   - Put repeatable procedures in `repo-governance/workflows/`.
   - Put directory-specific rules in the nearest scoped `AGENTS.md`.
   - Put capability-specific guidance in the relevant `SKILL.md`.
   - Create and categorize a focused document or subdirectory in
     `repo-governance/` when no existing location is a suitable canonical home.
     Update the relevant README so people and agents can discover it; do not
     create empty categories. In `docs/` and `repo-governance/`, every affected
     directory must have a README that indexes its immediate Markdown documents
     and child directories; see the
     [documentation index policy](../documentation-index-policy.md).

3. Search the inventory for equivalent, overlapping, or inverse rules. Merge
   equivalent guidance into the canonical source, update references, and remove
   redundant wording only when its meaning is fully preserved.

4. Resolve contradictions before writing. Do not silently choose between rules
   that differ in requirement, scope, priority, or verification. Present the
   conflicting paths, the relevant text, practical effect, and a recommended
   resolution to the repository owner. Wait for a decision when the conflict is
   substantive.

5. Integrate the approved rule using direct, testable language. Link from a
   concise document to its detailed source instead of copying the same rule.

## Verification

Confirm the rule has one canonical source, references are accurate, and no
contradictory guidance remains in the applicable files. Run:

```sh
npm run format:check
npm run check:governance
```

## Recovery

If scope or precedence remains unclear, keep the existing rules unchanged and
ask for clarification. Split an overlong document into focused Markdown files
under `repo-governance/` to preserve progressive disclosure.
