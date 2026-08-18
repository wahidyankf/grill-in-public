package parity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckAcceptsHarnessesThatExposeTheSameEntries(t *testing.T) {
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer", "repo-explorer")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsAHarnessMissingASubagent(t *testing.T) {
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")
	// Adding an agent for one tool and forgetting the others is the mistake
	// this check exists to catch.
	writeFile(t, filepath.Join(root, ".claude/agents/planner.md"), "planner")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("expected one finding per lagging harness, got %#v", findings)
	}
	for _, finding := range findings {
		if finding.Capability != "subagent" || len(finding.Missing) != 1 || finding.Missing[0] != "planner" {
			t.Fatalf("unexpected finding: %#v", finding)
		}
		if finding.Harness == "Claude Code" {
			t.Fatalf("the harness that has the agent must not be reported: %#v", finding)
		}
	}
}

func TestCheckIgnoresDirectoryIndexes(t *testing.T) {
	// Every harness directory carries a README index. Counting it would report
	// a difference for a file that defines no capability.
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")
	writeFile(t, filepath.Join(root, ".claude/agents/README.md"), "index")
	writeFile(t, filepath.Join(root, ".opencode/agents/README.md"), "index")
	writeFile(t, filepath.Join(root, ".codex/agents/README.md"), "index")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckSkipsCapabilitiesNoHarnessUses(t *testing.T) {
	// A repository with no skills has not decided to have any, so demanding
	// them everywhere would invent work rather than report a gap.
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	for _, finding := range findings {
		if finding.Capability == "skill" {
			t.Fatalf("expected no skill findings, got %#v", finding)
		}
	}
}

func TestCheckRequiresASkillInEveryHarnessDirectory(t *testing.T) {
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")
	writeFile(t, filepath.Join(root, ".claude/skills/review/SKILL.md"), "review")
	// A directory without the manifest is supporting material, not a skill.
	writeFile(t, filepath.Join(root, ".claude/skills/review/reference.md"), "notes")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("expected only Codex to be missing the skill, got %#v", findings)
	}
	if findings[0].Harness != "Codex" || findings[0].Missing[0] != "review" {
		t.Fatalf("unexpected finding: %#v", findings[0])
	}
}

func TestCheckReadsOpencodeSkillsFromTheSharedDirectories(t *testing.T) {
	// opencode loads .claude/skills and .agents/skills as well as its own, so a
	// skill mirrored for the other two harnesses already reaches it.
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")
	writeFile(t, filepath.Join(root, ".claude/skills/review/SKILL.md"), "review")
	writeFile(t, filepath.Join(root, ".agents/skills/review/SKILL.md"), "review")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckExemptsCodexFromCommands(t *testing.T) {
	// Codex has no project command directory, so it cannot be counted for one.
	root := t.TempDir()
	writeSubagents(t, root, "drill-reviewer")
	writeFile(t, filepath.Join(root, ".claude/commands/review.md"), "review")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Harness != "opencode" {
		t.Fatalf("expected only opencode to be reported, got %#v", findings)
	}
	if notes := UnsupportedNotes(); len(notes) == 0 || !strings.Contains(notes[0], "Codex") {
		t.Fatalf("expected the Codex exemption to be reported, got %v", notes)
	}
}

func TestFindingMessageNamesTheCapabilityAndHarness(t *testing.T) {
	message := Finding{Capability: "subagent", Harness: "Codex", Missing: []string{"planner"}}.Message()

	if !strings.Contains(message, "subagent") || !strings.Contains(message, "Codex") ||
		!strings.Contains(message, "planner") {
		t.Fatalf("expected the capability, harness, and entry, got %q", message)
	}
}

func writeSubagents(t *testing.T, root string, names ...string) {
	t.Helper()
	for _, name := range names {
		writeFile(t, filepath.Join(root, ".claude/agents", name+".md"), name)
		writeFile(t, filepath.Join(root, ".codex/agents", name+".toml"), name)
		writeFile(t, filepath.Join(root, ".opencode/agents", name+".md"), name)
	}
}

func writeFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create the directory for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
