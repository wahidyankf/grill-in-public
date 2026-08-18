// Package rulechange detects edits to the repository's rules so the
// rules-propagation workflow is announced instead of being remembered.
package rulechange

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// Workflow is the procedure every rule change must follow. The notice points
// at it rather than restating its steps, so the workflow stays the one source.
const Workflow = "repo-governance/workflows/rules-propagation.md"

// HarnessWorkflow proves the harnesses stayed equal after a change to one of
// them. A harness change is also a rule change, so both workflows can apply.
const HarnessWorkflow = "repo-governance/workflows/harness-alignment.md"

// harnessFiles and harnessDirectories are the paths a harness reads: the agent
// instruction files, each tool's configuration, and the directories holding its
// subagents, skills, and commands. Changing one of them can leave the harnesses
// unequal, which is what the align workflow checks.
var (
	harnessFiles = []string{
		"AGENTS.md",
		"CLAUDE.md",
		"opencode.json",
	}

	harnessDirectories = []string{
		".claude",
		".codex",
		".opencode",
		".agents",
	}
)

// ruleFiles and ruleDirectories add the paths that carry rules without being
// harness surfaces: the shared governance and the Git hooks that enforce it.
var (
	ruleFiles = harnessFiles

	ruleDirectories = append([]string{
		"repo-governance",
		".husky",
	}, harnessDirectories...)
)

// RulePaths returns the given repository-relative paths that carry rules, in
// sorted order and without duplicates, so a caller can report them stably.
func RulePaths(paths []string) []string {
	return selectPaths(paths, ruleFiles, ruleDirectories)
}

// HarnessPaths returns the given paths a harness reads. Every harness path is
// also a rule path; the narrower list exists because only these can leave the
// harnesses unequal.
func HarnessPaths(paths []string) []string {
	return selectPaths(paths, harnessFiles, harnessDirectories)
}

func selectPaths(paths []string, files []string, directories []string) []string {
	seen := make(map[string]struct{}, len(paths))
	var matches []string

	for _, path := range paths {
		normalized := normalize(path)
		if normalized == "" || !matchesAny(normalized, files, directories) {
			continue
		}
		if _, duplicate := seen[normalized]; duplicate {
			continue
		}
		seen[normalized] = struct{}{}
		matches = append(matches, normalized)
	}

	sort.Strings(matches)
	return matches
}

// StagedPaths lists the paths staged for the next commit. Staged content is
// what a commit will record, so it is the honest input for a pre-commit check.
func StagedPaths(root string) ([]string, error) {
	// -z keeps paths intact when a name contains a space or a newline, which
	// splitting on plain newlines would corrupt.
	output, err := exec.Command("git", "-C", root, "diff", "--cached", "--name-only", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("list staged paths: %w", err)
	}

	var paths []string
	for path := range strings.SplitSeq(string(output), "\x00") {
		if path != "" {
			paths = append(paths, path)
		}
	}
	return paths, nil
}

// HookEvent is the part of a harness pre-edit hook payload this tool needs.
// Harnesses name the target differently: Claude Code sends the file path, while
// Codex sends the whole patch, so both shapes are read.
type HookEvent struct {
	ToolInput struct {
		FilePath     string `json:"file_path"`
		NotebookPath string `json:"notebook_path"`
		Command      string `json:"command"`
	} `json:"tool_input"`
	Cwd string `json:"cwd"`
}

// patchHeaders introduce a file in an apply_patch payload. A patch names every
// file it touches on one of these lines, which is what makes the paths readable
// before the edit lands.
var patchHeaders = []string{
	"*** Add File: ",
	"*** Update File: ",
	"*** Delete File: ",
	"*** Move to: ",
}

// HookPaths reads a harness hook payload and returns the repository-relative
// paths it is about to edit. An unreadable payload yields no paths rather than
// an error, because a notice must never break the edit it comments on.
func HookPaths(payload []byte, root string) []string {
	var event HookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil
	}

	var paths []string
	for _, candidate := range []string{event.ToolInput.FilePath, event.ToolInput.NotebookPath} {
		if candidate == "" {
			continue
		}
		paths = append(paths, relativeTo(root, candidate))
	}
	for _, candidate := range patchPaths(event.ToolInput.Command) {
		paths = append(paths, relativeTo(root, candidate))
	}
	return paths
}

// patchPaths reads the files an apply_patch payload touches. A shell command
// arrives in the same field and simply contains no patch header, so it yields
// nothing rather than a false match.
func patchPaths(command string) []string {
	var paths []string
	for line := range strings.SplitSeq(command, "\n") {
		trimmed := strings.TrimSpace(line)
		for _, header := range patchHeaders {
			if !strings.HasPrefix(trimmed, header) {
				continue
			}
			if path := strings.TrimSpace(strings.TrimPrefix(trimmed, header)); path != "" {
				paths = append(paths, path)
			}
			break
		}
	}
	return paths
}

// Notice describes the detected rule change and names the workflows to follow.
// It names a workflow only when its paths changed, so a notice that always
// listed both would teach readers to ignore the second line.
func Notice(paths []string) string {
	notice := fmt.Sprintf(
		"Rule change detected in %s.\nRead and follow %s before completing this change.",
		strings.Join(paths, ", "),
		Workflow,
	)

	harnessPaths := HarnessPaths(paths)
	if len(harnessPaths) == 0 {
		return notice
	}

	return fmt.Sprintf(
		"%s\nHarness setup changed in %s.\nRun %s so every harness stays equal.",
		notice,
		strings.Join(harnessPaths, ", "),
		HarnessWorkflow,
	)
}

func matchesAny(path string, files []string, directories []string) bool {
	if slices.Contains(files, path) {
		return true
	}
	for _, directory := range directories {
		if path == directory || strings.HasPrefix(path, directory+"/") {
			return true
		}
	}
	return false
}

// normalize turns a path into the repository-relative slash form the rule lists
// use, so Windows separators and ./ prefixes still match.
func normalize(path string) string {
	cleaned := filepath.ToSlash(filepath.Clean(strings.TrimSpace(path)))
	if cleaned == "." {
		return ""
	}
	return strings.TrimPrefix(cleaned, "./")
}

// relativeTo converts an absolute hook path into a repository-relative one and
// leaves anything outside the repository untouched, where it will not match.
func relativeTo(root string, path string) string {
	if !filepath.IsAbs(path) {
		return path
	}
	relative, err := filepath.Rel(root, path)
	if err != nil {
		return path
	}
	return relative
}
