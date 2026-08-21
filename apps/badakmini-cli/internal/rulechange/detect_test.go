package rulechange

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRulePathsSelectsRuleCarryingPaths(t *testing.T) {
	paths := RulePaths([]string{
		"AGENTS.md",
		"CLAUDE.md",
		"opencode.json",
		"repo-governance/conventions/agent-vocabulary.md",
		".claude/agents/drill-reviewer.md",
		".codex/config.toml",
		".opencode/README.md",
		".husky/pre-push",
		// Ordinary work carries no rules and must stay silent, or the notice
		// becomes noise that contributors learn to skip.
		"apps/badakmini-cli/cmd/badak-mini/main.go",
		"docs/how-to/run-nx-workspace.md",
		"README.md",
	})

	expected := []string{
		".claude/agents/drill-reviewer.md",
		".codex/config.toml",
		".husky/pre-push",
		".opencode/README.md",
		"AGENTS.md",
		"CLAUDE.md",
		"opencode.json",
		"repo-governance/conventions/agent-vocabulary.md",
	}
	if strings.Join(paths, ",") != strings.Join(expected, ",") {
		t.Fatalf("expected %v, got %v", expected, paths)
	}
}

func TestRulePathsNormalizesAndDeduplicates(t *testing.T) {
	// A hook and a Git listing can name the same file differently, and one
	// change deserves one line in the notice.
	paths := RulePaths([]string{"./AGENTS.md", "AGENTS.md", "repo-governance/../AGENTS.md"})

	if len(paths) != 1 || paths[0] != "AGENTS.md" {
		t.Fatalf("expected a single AGENTS.md entry, got %v", paths)
	}
}

func TestRulePathsIgnoresLookalikePrefixes(t *testing.T) {
	// A sibling whose name merely starts with a rule directory's name is not
	// inside it, so prefix matching must respect the separator.
	paths := RulePaths([]string{".claude-backup/settings.json", "repo-governance-notes/idea.md"})

	if len(paths) != 0 {
		t.Fatalf("expected no matches, got %v", paths)
	}
}

func TestHookPathsReadsEditPayload(t *testing.T) {
	root := t.TempDir()
	payload := []byte(`{"tool_name":"Edit","tool_input":{"file_path":"` +
		filepath.ToSlash(filepath.Join(root, "AGENTS.md")) + `"}}`)

	paths := RulePaths(HookPaths(payload, root))

	if len(paths) != 1 || paths[0] != "AGENTS.md" {
		t.Fatalf("expected AGENTS.md, got %v", paths)
	}
}

func TestHookPathsReadsAnApplyPatchPayload(t *testing.T) {
	// Codex reports the patch itself instead of a file path, so the paths have
	// to be read from the patch headers to detect the same change.
	payload := []byte(`{"tool_name":"apply_patch","tool_input":{"command":` +
		`"*** Begin Patch\n*** Update File: AGENTS.md\n-old\n+new\n` +
		`*** Add File: .claude/agents/planner.md\n+text\n*** End Patch\n"}}`)

	paths := RulePaths(HookPaths(payload, t.TempDir()))

	expected := []string{".claude/agents/planner.md", "AGENTS.md"}
	if strings.Join(paths, ",") != strings.Join(expected, ",") {
		t.Fatalf("expected %v, got %v", expected, paths)
	}
}

func TestHookPathsIgnoresAShellCommandPayload(t *testing.T) {
	// The same field carries shell commands, and ordinary work must stay silent.
	payload := []byte(`{"tool_name":"Bash","tool_input":{"command":"npm test"}}`)

	if paths := RulePaths(HookPaths(payload, t.TempDir())); len(paths) != 0 {
		t.Fatalf("expected no paths, got %v", paths)
	}
}

func TestHookPathsIgnoresUnreadablePayload(t *testing.T) {
	// A notice must never break the edit it comments on, so malformed input
	// yields nothing to report instead of an error.
	if paths := HookPaths([]byte("not json"), t.TempDir()); paths != nil {
		t.Fatalf("expected no paths, got %v", paths)
	}
}

func TestNoticeNamesTheWorkflowWithoutRestatingIt(t *testing.T) {
	notice := Notice([]string{"AGENTS.md"})

	if !strings.Contains(notice, "AGENTS.md") || !strings.Contains(notice, Workflow) {
		t.Fatalf("expected the changed path and the workflow, got %q", notice)
	}
}

func TestHarnessPathsSelectsTheHarnessSurfaces(t *testing.T) {
	paths := HarnessPaths([]string{
		"AGENTS.md",
		"CLAUDE.md",
		"opencode.json",
		".claude/skills/review/SKILL.md",
		".codex/agents/repo-explorer.toml",
		".opencode/commands/review.md",
		".agents/skills/review/SKILL.md",
		// Governance and Git hooks carry rules without being a harness
		// surface, so aligning the harnesses cannot be what they need.
		"repo-governance/conventions/agent-vocabulary.md",
		".husky/pre-push",
	})

	expected := []string{
		".agents/skills/review/SKILL.md",
		".claude/skills/review/SKILL.md",
		".codex/agents/repo-explorer.toml",
		".opencode/commands/review.md",
		"AGENTS.md",
		"CLAUDE.md",
		"opencode.json",
	}
	if strings.Join(paths, ",") != strings.Join(expected, ",") {
		t.Fatalf("expected %v, got %v", expected, paths)
	}
}

func TestNoticeAddsTheHarnessWorkflowForAHarnessPath(t *testing.T) {
	notice := Notice(RulePaths([]string{".claude/agents/drill-reviewer.md"}))

	if !strings.Contains(notice, Workflow) || !strings.Contains(notice, HarnessWorkflow) {
		t.Fatalf("expected both workflows, got %q", notice)
	}
}

func TestNoticeOmitsTheHarnessWorkflowForOtherRulePaths(t *testing.T) {
	// A governance edit needs the rules propagated, not the harnesses
	// compared, and naming a workflow that does not apply trains readers to
	// skim past the one that does.
	notice := Notice(RulePaths([]string{"repo-governance/README.md", ".husky/pre-push"}))

	if strings.Contains(notice, HarnessWorkflow) {
		t.Fatalf("expected no harness workflow, got %q", notice)
	}
}
