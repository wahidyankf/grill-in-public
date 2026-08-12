package markdownlinks

import (
	"os"
	"os/exec"
	"path/filepath"
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
	if err := os.Symlink(filepath.Join(externalRoot, "external.md"), filepath.Join(root, "docs", "external.md")); err != nil {
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
	if findings[0].Destination != "docs/missing.md" || findings[0].Problem != "targets a file that does not exist" {
		t.Fatalf("unexpected missing target finding: %#v", findings[0])
	}
	if findings[1].Destination != "docs/guide.md#not-there" || findings[1].Problem != "targets a heading that does not exist" {
		t.Fatalf("unexpected missing anchor finding: %#v", findings[1])
	}
	if findings[2].Destination != "../outside.md" || findings[2].Problem != "points outside this repository" {
		t.Fatalf("unexpected outside finding: %#v", findings[2])
	}
	if findings[3].Destination != "docs/external.md" || findings[3].Problem != "resolves outside this repository" {
		t.Fatalf("unexpected symlink finding: %#v", findings[3])
	}
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

func newRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	// Markdown source selection calls git ls-files, so each fixture must be a
	// genuine repository rather than only a temporary directory.
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
	command := exec.Command("git", arguments...)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("stage test files: %s: %v", output, err)
	}
}

func writeMarkdown(t *testing.T, root string, relativePath string, contents string) {
	t.Helper()
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create parent directory for %s: %v", relativePath, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}
