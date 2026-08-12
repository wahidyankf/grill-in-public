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
	governanceDirectory = "repo-governance"
)

// Finding describes one governance document that exceeds MaxWords.
type Finding struct {
	Path      string
	WordCount int
}

// Check validates AGENTS.md and recursive Markdown governance documents below
// root. It returns every over-limit document so contributors can fix all of
// them in one pass.
func Check(root string) ([]Finding, error) {
	// Fail fast for the two repository structures this policy promises to govern;
	// a missing structure is a setup error, not a zero-word success.
	if err := requireFile(filepath.Join(root, agentsFile), agentsFile); err != nil {
		return nil, err
	}
	if err := requireDirectory(filepath.Join(root, governanceDirectory), governanceDirectory); err != nil {
		return nil, err
	}

	findings, err := checkFile(root, agentsFile)
	if err != nil {
		return nil, err
	}

	governancePath := filepath.Join(root, governanceDirectory)
	// Walk recursively because progressive disclosure permits nested, focused
	// policies. Restricting the scan to Markdown leaves code and data untouched.
	err = filepath.WalkDir(governancePath, func(path string, entry os.DirEntry, walkErr error) error {
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
		findings = append(findings, fileFindings...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", governanceDirectory, err)
	}

	return findings, nil
}

func requireFile(path string, displayPath string) error {
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

func requireDirectory(path string, displayPath string) error {
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

func checkFile(root string, relativePath string) ([]Finding, error) {
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
