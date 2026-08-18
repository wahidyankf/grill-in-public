# Shared Agent Directory

This directory holds capabilities a harness reads from a shared location rather than from its own harness directory. Codex reads skills from here; see the [agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- [`skills/`](skills/README.md) — the shared skills available in this repository.

Nothing here is Codex-specific by nature. A harness that adopts the same shared path reads the same files, which is the point of keeping them outside `.codex/`.
