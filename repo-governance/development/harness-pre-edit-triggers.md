---
tldr: "Records the pre-edit trigger each harness wires and how far each one is verified."
when_to_use: "Use when wiring, changing, or deciding how far to trust a harness pre-edit hook."
---

# Harness Pre-Edit Triggers

A pre-edit trigger says the same thing earlier, while the change is still being written. Support differs, and the difference is recorded rather than assumed:

| Harness | Pre-edit trigger |
| --- | --- |
| Claude Code | `PreToolUse` on `Edit`, `Write`, and `NotebookEdit`, wired in `.claude/settings.json`; verified |
| opencode | `.opencode/plugin/rule-change-notice.js` on `tool.execute.before`; the plugin loads, firing is unverified here |
| Codex | `PreToolUse` on `apply_patch`, `Edit`, and `Write`, wired in `.codex/hooks.json`; it runs only after the owner trusts the project and approves the hook with `/hooks`, so firing is unverified here |

Each harness asks Badak Mini for the notice rather than keeping its own copy of the rule paths, so the three cannot drift apart: all of them call `harness rule-change hook`. The payloads differ and the command reads both, since Claude Code names the file while Codex sends a patch whose file headers it parses.
