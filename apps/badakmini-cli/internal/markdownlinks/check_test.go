package markdownlinks

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckAcceptsValidLocalLinksAndIgnoresExternalLinks(t *testing.T) {
	root := newRepository(t)
	// Cover the supported local-link forms together and prove that network links
	// and fenced examples stay outside this deterministic checker.
	writeMarkdown(t, root, "README.md", `
[Guide](docs/guide.md)
[Guide section](docs/guide.md#getting-started)
[Guide with title](docs/guide.md "Optional title")
[Setext section](docs/guide.md#setext-heading)
[Local section](#introduction)
[Website](https://example.com/docs)
[Email](mailto:hello@example.com)

# Introduction

~~~markdown
[Ignored](also-missing.md)
~~~
`)
	writeMarkdown(t, root, "docs/guide.md", "# Getting Started\n\nSetext Heading\n--------------\n")
	stageAll(t, root)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsMissingTargetsAndAnchors(t *testing.T) {
	root := newRepository(t)
	externalRoot := t.TempDir()
	writeMarkdown(t, root, "README.md", `
[Missing](docs/missing.md)
[Missing anchor](docs/guide.md#not-there)
[Outside](../outside.md)
[Symlink outside](docs/external.md)
`)
	writeMarkdown(t, root, "docs/guide.md", "# Present\n")
	writeMarkdown(t, externalRoot, "external.md", "# External\n")
	// The symlink turns an apparently local path into an escape attempt, which
	// requires a distinct check after normal lexical path validation.
	externalFile := filepath.Join(externalRoot, "external.md")
	linkedFile := filepath.Join(root, "docs", "external.md")
	if err := os.Symlink(externalFile, linkedFile); err != nil {
		t.Fatalf("create external symlink: %v", err)
	}
	stageAll(t, root)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 4 {
		t.Fatalf("expected four findings, got %#v", findings)
	}
	assertFinding(t, findings[0], "docs/missing.md", "targets a file that does not exist")
	assertFinding(t, findings[1], "docs/guide.md#not-there", "targets a heading that does not exist")
	assertFinding(t, findings[2], "../outside.md", "points outside this repository")
	assertFinding(t, findings[3], "docs/external.md", "resolves outside this repository")
}

func TestCheckSupportsReferencesPercentEncodingAndDuplicateAnchors(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", `
[Guide][guide]
[Directory](docs/#second-heading-1)

[guide]: docs/guide%20file.md
`)
	// The final fragment uses GitHub's -1 suffix for the second repeated heading.
	writeMarkdown(t, root, "docs/README.md", "# First Heading\n# Second Heading\n# Second Heading\n")
	writeMarkdown(t, root, "docs/guide file.md", "# Guide\n")
	stageAll(t, root)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckIgnoresRepositoryMetadataAndBuildOutputs(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", "# Repository\n")
	writeMarkdown(t, root, ".git/ignored.md", "[Broken](missing.md)\n")
	writeMarkdown(t, root, "node_modules/package/ignored.md", "[Broken](missing.md)\n")
	writeMarkdown(t, root, "apps/example/dist/ignored.md", "[Broken](missing.md)\n")
	// Only README.md enters the Git index; the three broken ignored files prove
	// the source scan follows Git rather than recursively walking the filesystem.
	stage(t, root, "README.md")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected generated files to be ignored, got %#v", findings)
	}
}

func TestCheckAcceptsUntrackedTargetsBeforeTheyAreStaged(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", "[Draft](docs/draft.md)\n")
	writeMarkdown(t, root, "docs/draft.md", "# Draft\n")
	stage(t, root, "README.md")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected an untracked target to be valid, got %#v", findings)
	}
}

func TestCheckReportsTargetsDeletedFromTheGitTree(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", "[Guide](docs/guide.md)\n")
	writeMarkdown(t, root, "docs/guide.md", "# Guide\n")
	stageAll(t, root)
	// Stage a deletion while keeping the source link. This models the regression
	// that the pre-push check is meant to catch after a document move or removal.
	// #nosec G204 -- the test controls the executable, repository, and fixed file argument.
	command := exec.Command("git", "-C", root, "rm", "--quiet", "--force", "docs/guide.md")
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("delete staged target: %s: %v", output, err)
	}

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Problem != "targets a file that does not exist" {
		t.Fatalf("expected a deleted target finding, got %#v", findings)
	}
}

func TestCheckSupportsParserEdgeCases(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", strings.Join([]string{
		"[Angle](<docs/guide file.md>)",
		"[Nested [label]](docs/nested_(guide).md)",
		"[Collapsed][]",
		"[Shortcut]",
		"[Titled](docs/titled.md\t\"Optional title\")",
		"`[Inline code](missing.md)`",
		`\[Escaped](missing.md)`,
		"",
		"[Collapsed]: docs/collapsed.md",
		"[Shortcut]: docs/shortcut.md",
	}, "\n"))
	for _, path := range []string{
		"docs/guide file.md",
		"docs/nested_(guide).md",
		"docs/collapsed.md",
		"docs/shortcut.md",
		"docs/titled.md",
	} {
		writeMarkdown(t, root, path, "# Present\n")
	}
	stageAll(t, root)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected parser edge cases to resolve, got %#v", findings)
	}
}

func TestMalformedMarkdownLinksAreNotParsed(t *testing.T) {
	references := map[string]string{"known": "known.md"}
	tests := []struct {
		name     string
		line     string
		position int
	}{
		{name: "unclosed label", line: "[label", position: 0},
		{name: "unclosed inline destination", line: "[label](target", position: 0},
		{name: "unclosed full reference", line: "[label][known", position: 0},
		{name: "unknown shortcut reference", line: "[unknown]", position: 0},
	}

	for _, test := range tests {
		_, _, ok := parseLinkAt(test.line, test.position, 1, references)
		if ok {
			t.Fatalf("expected %s to be rejected", test.name)
		}
	}
}

func TestInlineDestinationBoundaries(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		opening     int
		destination string
		ok          bool
	}{
		{name: "only whitespace", value: "( \t", opening: 0, ok: false},
		{
			name: "angle destination", value: "(<path with spaces.md>)", opening: 0,
			destination: "path with spaces.md", ok: true,
		},
		{name: "angle missing closer", value: "(<path.md)", opening: 0, ok: false},
		{name: "angle missing parenthesis", value: "(<path.md>", opening: 0, ok: false},
		{name: "nested parentheses", value: "(path_(nested).md)", opening: 0, destination: "path_(nested).md", ok: true},
		{name: "escaped parenthesis", value: `(path\).md)`, opening: 0, destination: `path\).md`, ok: true},
		{name: "title without closer", value: `(path.md "title"`, opening: 0, ok: false},
	}

	for _, test := range tests {
		destination, _, ok := inlineDestination(test.value, test.opening)
		if ok != test.ok || destination != test.destination {
			t.Fatalf(
				"%s: got destination %q and ok=%t, want %q and ok=%t",
				test.name, destination, ok, test.destination, test.ok,
			)
		}
	}
}

func TestMatchingBracketHandlesEscapesAndMissingCloser(t *testing.T) {
	if end, ok := matchingBracket(`[outer \[literal\]]`, 0); !ok || end != len(`[outer \[literal\]]`)-1 {
		t.Fatalf("expected escaped brackets to remain inside the label, got end=%d ok=%t", end, ok)
	}
	if _, ok := matchingBracket("[missing", 0); ok {
		t.Fatal("expected an unmatched label to be rejected")
	}
	if isEscaped(`\\[label]`, 2) {
		t.Fatal("expected an even backslash run to leave the bracket active")
	}
}

func TestCheckReportsTargetTypeAndDirectoryFragmentFailures(t *testing.T) {
	root := newRepository(t)
	externalRoot := t.TempDir()
	writeMarkdown(t, root, "README.md", strings.Join([]string{
		"[Directory without index](empty/#missing)",
		"[Text fragment](notes.txt#missing)",
		"[Escaped URL](%)",
		"[Cyclic target](cyclic.md)",
		"[External directory index](linked/#outside)",
	}, "\n"))
	if err := os.Mkdir(filepath.Join(root, "empty"), 0o750); err != nil {
		t.Fatalf("create empty directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "notes.txt"), []byte("plain text"), 0o600); err != nil {
		t.Fatalf("write text target: %v", err)
	}
	cyclicPath := filepath.Join(root, "cyclic.md")
	if err := os.Symlink(cyclicPath, cyclicPath); err != nil {
		t.Fatalf("create cyclic target: %v", err)
	}
	if err := os.Mkdir(filepath.Join(root, "linked"), 0o750); err != nil {
		t.Fatalf("create linked directory: %v", err)
	}
	writeMarkdown(t, externalRoot, "README.md", "# Outside\n")
	externalReadme := filepath.Join(externalRoot, "README.md")
	linkedReadme := filepath.Join(root, "linked", "README.md")
	if err := os.Symlink(externalReadme, linkedReadme); err != nil {
		t.Fatalf("create external README symlink: %v", err)
	}
	stage(t, root, "README.md", "notes.txt")

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 5 {
		t.Fatalf("expected five findings, got %#v", findings)
	}
	assertFinding(t, findings[0], "empty/#missing", "targets a directory without README.md for its fragment")
	assertFinding(t, findings[1], "notes.txt#missing", "uses a fragment on a non-Markdown target")
	assertFinding(t, findings[2], "%", "has an invalid URL")
	assertFinding(t, findings[3], "cyclic.md", "cannot inspect its target")
	assertFinding(t, findings[4], "linked/#outside", "resolves outside this repository")
}

func TestCheckSupportsRepositoryRelativePaths(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "docs/source.md", "[Root guide](/guide.md)\n")
	writeMarkdown(t, root, "guide.md", "# Guide\n")
	stageAll(t, root)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected a valid repository-relative link, got %#v", findings)
	}
}

func TestCheckReportsSourceReadFailure(t *testing.T) {
	root := newRepository(t)
	writeMarkdown(t, root, "README.md", "# Tracked\n")
	stageAll(t, root)
	if err := os.Remove(filepath.Join(root, "README.md")); err != nil {
		t.Fatalf("remove tracked source: %v", err)
	}
	if err := os.Mkdir(filepath.Join(root, "README.md"), 0o750); err != nil {
		t.Fatalf("replace tracked source with directory: %v", err)
	}

	_, err := Check(root)
	if err == nil || !strings.Contains(err.Error(), "read README.md") {
		t.Fatalf("expected a tracked source read error, got %v", err)
	}
}

func TestCheckReportsGitDiscoveryFailure(t *testing.T) {
	_, err := Check(t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "list tracked repository files") {
		t.Fatalf("expected a Git discovery error, got %v", err)
	}
}

func newRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	// Markdown source selection calls git ls-files, so each fixture must be a
	// genuine repository rather than only a temporary directory.
	// #nosec G204 -- the test controls the executable and temporary repository path.
	command := exec.Command("git", "init", "--quiet", root)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("initialize test repository: %s: %v", output, err)
	}
	return root
}

func stageAll(t *testing.T, root string) {
	t.Helper()
	stage(t, root, ".")
}

func stage(t *testing.T, root string, paths ...string) {
	t.Helper()
	// The -- separator makes a test filename beginning with '-' a path, not a
	// Git option, mirroring the production command's filename safety.
	arguments := append([]string{"-C", root, "add", "--"}, paths...)
	// #nosec G204 -- callers provide only fixed Git test operations against a temporary repository.
	command := exec.Command("git", arguments...)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("stage test files: %s: %v", output, err)
	}
}

func writeMarkdown(t *testing.T, root, relativePath, contents string) {
	t.Helper()
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("create parent directory for %s: %v", relativePath, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}

func assertFinding(t *testing.T, finding Finding, destination, problem string) {
	t.Helper()
	if finding.Destination != destination || finding.Problem != problem {
		t.Fatalf("expected destination %q and problem %q, got %#v", destination, problem, finding)
	}
}
