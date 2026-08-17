package governance

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckAcceptsDocumentsAtTheWordLimit(t *testing.T) {
	// The boundary is inclusive: a document is only overlong once it exceeds the
	// progressive-disclosure limit, not when it reaches it.
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(MaxWords))
	writeFile(t, filepath.Join(root, claudeFile), words(MaxWords))
	writeFile(t, filepath.Join(root, governanceDirectory, "policy.md"), words(MaxWords))

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsDocumentsOverTheWordLimit(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(MaxWords+1))
	writeFile(t, filepath.Join(root, claudeFile), words(MaxWords+1))
	writeFile(t, filepath.Join(root, governanceDirectory, "nested", "policy.md"), words(MaxWords+1))
	// Non-Markdown material may be long without being repository guidance.
	writeFile(t, filepath.Join(root, governanceDirectory, "nested", "ignored.txt"), words(MaxWords+1))

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 3 {
		t.Fatalf("expected three findings, got %#v", findings)
	}
	if findings[0].Path != agentsFile || findings[0].WordCount != MaxWords+1 {
		t.Fatalf("unexpected root finding: %#v", findings[0])
	}
	// Every instruction file shares the limit, so CLAUDE.md must be reported beside
	// AGENTS.md rather than being exempt from the concise-guidance budget.
	if findings[1].Path != claudeFile || findings[1].WordCount != MaxWords+1 {
		t.Fatalf("unexpected instruction file finding: %#v", findings[1])
	}
	if findings[2].Path != "repo-governance/nested/policy.md" {
		t.Fatalf("expected recursive Markdown finding, got %#v", findings[2])
	}
}

func TestCheckRequiresAgentsFile(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, claudeFile), words(1))

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "required file not found: AGENTS.md") {
		t.Fatalf("expected a missing AGENTS.md error, got %v", err)
	}
}

func TestCheckRequiresClaudeFile(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "required file not found: CLAUDE.md") {
		t.Fatalf("expected a missing CLAUDE.md error, got %v", err)
	}
}

func TestCheckRequiresGovernanceDirectory(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "required directory not found: repo-governance") {
		t.Fatalf("expected a missing governance directory error, got %v", err)
	}
}

func newRepositoryFixture(t *testing.T) string {
	t.Helper()
	// Fixtures create only the minimum valid repository shape so individual
	// tests can remove or alter one prerequisite deliberately.
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, governanceDirectory), 0o755); err != nil {
		t.Fatalf("create governance directory: %v", err)
	}
	return root
}

func writeFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create parent directory for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func words(count int) string {
	return strings.TrimSpace(strings.Repeat("word ", count))
}
