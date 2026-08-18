# opencode Plugins

opencode loads every plugin in this directory. It reads JavaScript and TypeScript here, so this README is not loaded as a plugin.

## Contents

- `rule-change-notice.js` — announces the rule-change workflows before an edit to a rule file. It only logs, so a failure cannot affect the edit.

The pre-commit hook is the guaranteed trigger for every harness; this plugin only announces the same thing earlier. See the [rule change trigger policy](../../repo-governance/development/rule-change-trigger-policy.md).
