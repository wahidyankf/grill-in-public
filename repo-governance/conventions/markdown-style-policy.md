---
tldr: "Defines a repository-wide source style for Markdown prose and diagrams."
when_to_use: "Use when creating, editing, reviewing, or formatting any Markdown file."
---

# Markdown Style Policy

## Scope

This policy applies to every committed `*.md` file in this repository, including human documentation, agent instructions, governance, project READMEs, CV material, and scripts documentation.

## Paragraphs

Write each prose paragraph as one continuous source line, separated from the next paragraph by one blank line. Do not manually hard-wrap flowing prose to a visual column width. Use structural Markdown, such as headings, lists, tables, blockquotes, and fenced code blocks, when the content is structurally distinct instead of splitting a paragraph for appearance.

## Diagrams and Schemas

For diagrams, schemas, flows, and similar visual models in Markdown, prefer ASCII art in a fenced `text` block. Use only ASCII characters such as `+`, `-`, `|`, and `>` with clear labels so the model remains legible in a terminal, NVIM, code review, and a rendered Markdown view. Prefer a Markdown table when a tabular schema communicates the relationship more directly.

Do not use Mermaid by default. Use it only when the task, user, or governing requirement explicitly calls for Mermaid; otherwise, choose terminal-readable ASCII art.

## Enforcement

Prettier is configured with `proseWrap: "never"` so `npm run format`, `npm run format:check`, and staged Markdown formatting preserve this style. Run `npm run format` after changing Markdown; use `npm run format:check` to verify it before committing.
