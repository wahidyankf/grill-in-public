---
tldr: "Treats every maintained line and tool surface as a liability that must earn its cost."
when_to_use: "Use before adding code, scripts, commands, tooling, automation, or configuration."
---

# Maintenance Value

Every maintained line and tool surface is a liability. Before adding code, scripts, commands, tooling, automation, or configuration, prove that a recurring concrete benefit exceeds its implementation, review, execution, documentation, debugging, upgrade, and removal costs. Also prove that an existing mechanism or a clear manual step cannot satisfy the need. If either case is unproven, do not add the surface.

Prefer deleting, simplifying, or reusing over creating. Do not automate a rare action merely because it can be automated, duplicate an executable rule as metadata, or create a repository command for a check whose false confidence or upkeep costs more than running the underlying standard tool directly.

Review the net value again when a surface grows or its original need disappears. Passing tests prove behavior, not that the behavior deserves permanent maintenance.
