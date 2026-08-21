package governance

import (
	"errors"
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

func TestCheckReportsOverlongHarnessReadmes(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	writeFile(t, filepath.Join(root, ".claude", readmeFile), words(MaxWords+1))
	// A nested index is governed too, because depth does not make a directory
	// listing longer to read.
	writeFile(t, filepath.Join(root, ".opencode", "agents", readmeFile), words(MaxWords+1))
	// An agent definition beside the index is a prompt, not an index, so its
	// length is the author's call.
	writeFile(t, filepath.Join(root, ".opencode", "agents", "drill-reviewer.md"), words(MaxWords+1))
	// Only the configured harnesses are scanned; .codex is absent here.
	// Installed dependencies and caches are vendored content, so their own
	// READMEs must not be reported as this repository's overlong guidance.
	writeFile(t, filepath.Join(root, ".opencode", "node_modules", "dep", readmeFile), words(MaxWords+1))
	writeFile(t, filepath.Join(root, ".claude", ".cache", readmeFile), words(MaxWords+1))

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("expected two harness findings, got %#v", findings)
	}
	if findings[0].Path != ".claude/README.md" {
		t.Fatalf("unexpected first harness finding: %#v", findings[0])
	}
	if findings[1].Path != ".opencode/agents/README.md" {
		t.Fatalf("unexpected nested harness finding: %#v", findings[1])
	}
}

func TestCheckReportsTheSharedHarnessDirectory(t *testing.T) {
	// .agents holds the skills Codex and opencode both read, so its indexes are
	// governed guidance like any other harness directory's.
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	writeFile(t, filepath.Join(root, ".agents", readmeFile), words(MaxWords+1))
	// A skill body is a prompt, so it stays unmeasured beside its index.
	writeFile(t, filepath.Join(root, ".agents", "skills", "grill-me", "SKILL.md"), words(MaxWords+1))

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Path != ".agents/README.md" {
		t.Fatalf("expected the shared harness index to be reported, got %#v", findings)
	}
}

func TestCheckAcceptsRepositoriesWithoutHarnessDirectories(t *testing.T) {
	// A repository may configure no harness at all; that is a valid state, not a
	// setup error, so the check must stay silent rather than fail.
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
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

func TestCheckRejectsDirectoryInPlaceOfInstructionFile(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	if err := os.Mkdir(filepath.Join(root, agentsFile), 0o750); err != nil {
		t.Fatalf("create directory in place of AGENTS.md: %v", err)
	}

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "required file is not regular: AGENTS.md") {
		t.Fatalf("expected a non-regular AGENTS.md error, got %v", err)
	}
}

func TestCheckRejectsFileInPlaceOfGovernanceDirectory(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	writeFile(t, filepath.Join(root, governanceDirectory), words(1))

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "required path is not a directory: repo-governance") {
		t.Fatalf("expected a non-directory governance path error, got %v", err)
	}
}

func TestCheckFileReportsReadFailure(t *testing.T) {
	root := t.TempDir()

	_, err := checkFile(root, "missing.md")
	if err == nil || !strings.Contains(err.Error(), "read missing.md") {
		t.Fatalf("expected a document read error, got %v", err)
	}
}

func TestHarnessDirectoryReportsInspectionFailure(t *testing.T) {
	root := t.TempDir()
	harnessPath := filepath.Join(root, ".claude")
	if err := os.Symlink(harnessPath, harnessPath); err != nil {
		t.Fatalf("create cyclic harness symlink: %v", err)
	}

	_, err := checkHarnessDirectory(root, ".claude")
	if err == nil || !strings.Contains(err.Error(), "inspect .claude") {
		t.Fatalf("expected a harness inspection error, got %v", err)
	}
}

func TestInstructionCheckReportsDocumentReadFailure(t *testing.T) {
	root := t.TempDir()
	if err := os.Symlink("missing.md", filepath.Join(root, agentsFile)); err != nil {
		t.Fatalf("create broken instruction symlink: %v", err)
	}

	_, err := checkInstructionFiles(root)
	if err == nil || !strings.Contains(err.Error(), "read AGENTS.md") {
		t.Fatalf("expected an instruction read error, got %v", err)
	}
}

func TestCheckReportsGovernanceDocumentReadFailure(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	if err := os.Symlink("missing.md", filepath.Join(root, governanceDirectory, "broken.md")); err != nil {
		t.Fatalf("create broken governance symlink: %v", err)
	}

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "read repo-governance/broken.md") {
		t.Fatalf("expected a governance document read error, got %v", err)
	}
}

func TestCheckReportsHarnessReadmeReadFailure(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	if err := os.Mkdir(filepath.Join(root, ".agents"), 0o750); err != nil {
		t.Fatalf("create harness directory: %v", err)
	}
	if err := os.Symlink("missing.md", filepath.Join(root, ".agents", readmeFile)); err != nil {
		t.Fatalf("create broken harness README symlink: %v", err)
	}

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "read .agents/README.md") {
		t.Fatalf("expected a harness README read error, got %v", err)
	}
}

func TestCheckReportsHarnessInspectionFailure(t *testing.T) {
	root := newRepositoryFixture(t)
	writeFile(t, filepath.Join(root, agentsFile), words(1))
	writeFile(t, filepath.Join(root, claudeFile), words(1))
	harnessPath := filepath.Join(root, ".claude")
	if err := os.Symlink(harnessPath, harnessPath); err != nil {
		t.Fatalf("create cyclic harness symlink: %v", err)
	}

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "inspect .claude") {
		t.Fatalf("expected a harness inspection error, got %v", err)
	}
}

func TestRequiredPathsReportInspectionFailures(t *testing.T) {
	root := t.TempDir()
	cyclicPath := filepath.Join(root, "cyclic")
	if err := os.Symlink(cyclicPath, cyclicPath); err != nil {
		t.Fatalf("create cyclic path: %v", err)
	}

	if err := requireFile(cyclicPath, "cyclic-file"); err == nil || !strings.Contains(err.Error(), "inspect cyclic-file") {
		t.Fatalf("expected a file inspection error, got %v", err)
	}
	err := requireDirectory(cyclicPath, "cyclic-directory")
	if err == nil || !strings.Contains(err.Error(), "inspect cyclic-directory") {
		t.Fatalf("expected a directory inspection error, got %v", err)
	}
}

func TestVisitorsReturnFilesystemWalkErrors(t *testing.T) {
	walkErr := errors.New("walk failed")

	if err := visitGovernanceDocument("root", "path", nil, walkErr, nil); !errors.Is(err, walkErr) {
		t.Fatalf("expected governance visitor to return the walk error, got %v", err)
	}
	if err := visitHarnessReadme("root", "harness", "path", nil, walkErr, nil); !errors.Is(err, walkErr) {
		t.Fatalf("expected harness visitor to return the walk error, got %v", err)
	}
}

func TestVendoredDirectoryClassification(t *testing.T) {
	tests := []struct {
		name     string
		vendored bool
	}{
		{name: "node_modules", vendored: true},
		{name: ".cache", vendored: true},
		{name: "agents", vendored: false},
	}

	for _, test := range tests {
		if got := isVendoredDirectory(test.name); got != test.vendored {
			t.Fatalf("isVendoredDirectory(%q) = %t, want %t", test.name, got, test.vendored)
		}
	}
}

func newRepositoryFixture(t *testing.T) string {
	t.Helper()
	// Fixtures create only the minimum valid repository shape so individual
	// tests can remove or alter one prerequisite deliberately.
	root := t.TempDir()
	err := os.Mkdir(filepath.Join(root, governanceDirectory), 0o750)
	if err != nil {
		t.Fatalf("create governance directory: %v", err)
	}
	return root
}

func writeFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("create parent directory for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func words(count int) string {
	return strings.TrimSpace(strings.Repeat("word ", count))
}
