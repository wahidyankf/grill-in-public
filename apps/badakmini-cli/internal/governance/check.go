// Package governance validates the repository's concise-governance policy.
package governance

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// MaxWords is the hard word limit for each governed Markdown file.
	MaxWords = 500

	agentsFile          = "AGENTS.md"
	claudeFile          = "CLAUDE.md"
	governanceDirectory = "repo-governance"
	readmeFile          = "README.md"
)

// harnessDirectories hold each harness's project configuration, including the
// shared .agents directory more than one harness reads. Their READMEs are
// indexes, so they share the concise-guidance limit, while the agent, skill,
// and command definitions beside them are prompts and stay unmeasured.
var harnessDirectories = []string{".agents", ".claude", ".codex", ".opencode"}

// instructionFiles are the root agent instruction files. They share one limit
// because each must stay equally concise for the harness that reads it.
var instructionFiles = []string{agentsFile, claudeFile}

// Finding describes one governance document that exceeds MaxWords.
type Finding struct {
	Path      string
	WordCount int
}

// Check validates every root instruction file and the recursive Markdown
// governance documents below root. It returns every over-limit document so
// contributors can fix all of them in one pass.
func Check(root string) ([]Finding, error) {
	if err := requireGovernanceStructure(root); err != nil {
		return nil, err
	}

	findings, err := checkInstructionFiles(root)
	if err != nil {
		return nil, err
	}

	governanceFindings, err := checkGovernanceDocuments(root)
	if err != nil {
		return nil, err
	}
	findings = append(findings, governanceFindings...)

	harnessFindings, err := checkHarnessReadmes(root)
	if err != nil {
		return nil, err
	}
	findings = append(findings, harnessFindings...)

	return findings, nil
}

// requireGovernanceStructure distinguishes a missing policy surface from an
// empty, valid one before any word-count work begins.
func requireGovernanceStructure(root string) error {
	for _, instructionFile := range instructionFiles {
		if err := requireFile(filepath.Join(root, instructionFile), instructionFile); err != nil {
			return err
		}
	}

	return requireDirectory(filepath.Join(root, governanceDirectory), governanceDirectory)
}

func checkInstructionFiles(root string) ([]Finding, error) {
	var findings []Finding
	for _, instructionFile := range instructionFiles {
		fileFindings, err := checkFile(root, instructionFile)
		if err != nil {
			return nil, err
		}
		findings = append(findings, fileFindings...)
	}

	return findings, nil
}

func checkGovernanceDocuments(root string) ([]Finding, error) {
	var findings []Finding
	governancePath := filepath.Join(root, governanceDirectory)
	err := filepath.WalkDir(governancePath, func(path string, entry os.DirEntry, walkErr error) error {
		return visitGovernanceDocument(root, path, entry, walkErr, &findings)
	})
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", governanceDirectory, err)
	}

	return findings, nil
}

// visitGovernanceDocument limits the recursive scan to authored Markdown.
func visitGovernanceDocument(
	root, path string,
	entry os.DirEntry,
	walkErr error,
	findings *[]Finding,
) error {
	if walkErr != nil {
		return walkErr
	}
	if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
		return nil
	}

	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return fmt.Errorf("determine the path for %s: %w", path, err)
	}

	fileFindings, err := checkFile(root, relativePath)
	if err != nil {
		return err
	}
	*findings = append(*findings, fileFindings...)

	return nil
}

// checkHarnessReadmes measures the README index in every harness directory. A
// harness that this repository does not configure is absent rather than empty,
// so a missing directory is skipped instead of failing the run.
func checkHarnessReadmes(root string) ([]Finding, error) {
	var findings []Finding

	for _, harnessDirectory := range harnessDirectories {
		harnessFindings, err := checkHarnessDirectory(root, harnessDirectory)
		if err != nil {
			return nil, err
		}
		findings = append(findings, harnessFindings...)
	}

	return findings, nil
}

func checkHarnessDirectory(root, harnessDirectory string) ([]Finding, error) {
	harnessPath := filepath.Join(root, harnessDirectory)
	if _, err := os.Stat(harnessPath); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, fmt.Errorf("inspect %s: %w", harnessDirectory, err)
	}

	var findings []Finding
	err := filepath.WalkDir(harnessPath, func(path string, entry os.DirEntry, walkErr error) error {
		return visitHarnessReadme(root, harnessPath, path, entry, walkErr, &findings)
	})
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", harnessDirectory, err)
	}

	return findings, nil
}

func visitHarnessReadme(
	root, harnessPath, path string,
	entry os.DirEntry,
	walkErr error,
	findings *[]Finding,
) error {
	if walkErr != nil {
		return walkErr
	}
	if entry.IsDir() {
		// Harness-local dependency and cache directories are not authored indexes.
		if path != harnessPath && isVendoredDirectory(entry.Name()) {
			return filepath.SkipDir
		}

		return nil
	}
	if entry.Name() != readmeFile {
		return nil
	}

	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return fmt.Errorf("determine the path for %s: %w", path, err)
	}

	fileFindings, err := checkFile(root, relativePath)
	if err != nil {
		return err
	}
	*findings = append(*findings, fileFindings...)

	return nil
}

// isVendoredDirectory reports whether a directory inside a harness holds
// installed or generated content rather than authored configuration. Hidden
// directories are treated the same way because tools put caches there.
func isVendoredDirectory(name string) bool {
	return name == "node_modules" || strings.HasPrefix(name, ".")
}

func requireFile(path, displayPath string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("required file not found: %s", displayPath)
		}
		return fmt.Errorf("inspect %s: %w", displayPath, err)
	}
	if !info.Mode().IsRegular() {
		// A directory with the expected name must not silently pass as a document.
		return fmt.Errorf("required file is not regular: %s", displayPath)
	}
	return nil
}

func requireDirectory(path, displayPath string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("required directory not found: %s", displayPath)
		}
		return fmt.Errorf("inspect %s: %w", displayPath, err)
	}
	if !info.IsDir() {
		// Likewise, a file cannot serve as the recursive governance namespace.
		return fmt.Errorf("required path is not a directory: %s", displayPath)
	}
	return nil
}

func checkFile(root, relativePath string) ([]Finding, error) {
	// #nosec G304 -- relativePath is either a fixed policy path or emitted by a walk rooted under root.
	contents, err := os.ReadFile(filepath.Join(root, relativePath))
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", relativePath, err)
	}

	// Fields mirrors the previous shell check's whitespace-based definition of a
	// word and treats Markdown syntax as part of the concise-document budget.
	wordCount := len(strings.Fields(string(contents)))
	if wordCount <= MaxWords {
		return nil, nil
	}

	return []Finding{{Path: filepath.ToSlash(relativePath), WordCount: wordCount}}, nil
}
