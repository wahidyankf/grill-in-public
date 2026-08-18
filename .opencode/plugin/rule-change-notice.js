/**
 * Announces the rule-change workflows before opencode edits a rule file.
 *
 * The Git pre-commit hook is the guaranteed trigger for every tool. This plugin
 * moves the same notice earlier in an opencode session, so the workflow is read
 * while the change is still being written rather than once it is staged.
 *
 * The notice comes from Badak Mini rather than from a copy of the rule paths
 * kept here, because a second list would drift from the one the other harnesses
 * and the pre-commit hook use.
 *
 * It only reports. Blocking an edit here would put a rule in a config file,
 * which the agent harness support policy forbids.
 */

import { execFile } from "node:child_process";

const EDIT_TOOLS = new Set(["edit", "write", "patch", "multiedit"]);
const NOTICE_TIMEOUT_MS = 20000;

/** Asks Badak Mini whether a path carries rules, and for the notice to show. */
function requestNotice(directory, filePath) {
  return new Promise((resolve) => {
    const child = execFile(
      "go",
      [
        "-C",
        `${directory}/apps/badakmini-cli`,
        "run",
        "./cmd/badak-mini",
        "harness",
        "rule-change",
        "hook",
      ],
      { cwd: directory, timeout: NOTICE_TIMEOUT_MS },
      (error, stdout) => {
        // A notice that cannot be produced must not break the session, so every
        // failure resolves to no notice instead of rejecting.
        if (error || !stdout) {
          resolve("");
          return;
        }
        try {
          resolve(JSON.parse(stdout).systemMessage ?? "");
        } catch {
          resolve("");
        }
      },
    );

    child.stdin.end(JSON.stringify({ tool_input: { file_path: filePath } }));
  });
}

export const RuleChangeNotice = async ({ client, directory }) => {
  return {
    "tool.execute.before": async (input, output) => {
      if (!EDIT_TOOLS.has(String(input?.tool ?? "").toLowerCase())) return;

      const filePath = output?.args?.filePath;
      if (typeof filePath !== "string" || filePath === "") return;

      try {
        const message = await requestNotice(directory, filePath);
        if (message === "") return;

        // Logging keeps the notice out of the edit's return value, so a failure
        // here can never corrupt the tool call it is commenting on.
        await client.app.log({
          body: { service: "rule-change", level: "warn", message },
        });
      } catch {
        // A notice that cannot be delivered must not break the session.
      }
    },
  };
};
