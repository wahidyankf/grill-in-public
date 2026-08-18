// Package projecttargets verifies that every Nx project defines the targets the
// testing policy requires, so a project cannot quietly opt out of the checks
// pre-push runs on everyone else's behalf.
package projecttargets

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// projectFile is the Nx per-project configuration. Its presence is what makes a
// directory a project, so finding these files is the same as finding projects.
const projectFile = "project.json"

// quickTarget is the target pre-push runs for every affected project.
const quickTarget = "test:quick"

// requiredDependencies are the targets quick tests must run through dependsOn
// rather than reimplement, in the order the policy names them.
var requiredDependencies = []string{"typecheck", "lint"}

// skippedDirectories hold copies rather than sources: an installed dependency
// or a build output can carry a project file that describes nothing this
// repository maintains.
var skippedDirectories = map[string]bool{
	".git":         true,
	".nx":          true,
	"node_modules": true,
	"dist":         true,
}

// Finding is one project failing one requirement. A project can break several
// at once, so each is reported separately and repaired separately.
type Finding struct {
	Project string
	Path    string
	Problem string
}

// Message states the failure against the file a reader has to open to fix it.
func (finding Finding) Message() string {
	return fmt.Sprintf("%s: %s %s.", finding.Path, finding.Project, finding.Problem)
}

// configuration is the part of an Nx project file this check reads. Everything
// else is left untouched so an unrelated field cannot fail the parse.
type configuration struct {
	Name    string                 `json:"name"`
	Targets map[string]targetEntry `json:"targets"`
}

// targetEntry is one target. Cache is a pointer so an absent field stays
// distinct from an explicit false: cacheability is set in the workspace target
// defaults, and only an explicit false takes it away.
type targetEntry struct {
	DependsOn []json.RawMessage `json:"dependsOn"`
	Cache     *bool             `json:"cache"`
}

// Check inspects every project under root. It returns each failure rather than
// the first, so one run tells a contributor everything they have to repair.
func Check(root string) ([]Finding, error) {
	paths, err := findProjectFiles(root)
	if err != nil {
		return nil, err
	}

	findings := make([]Finding, 0)
	for _, path := range paths {
		projectFindings, err := checkProject(root, path)
		if err != nil {
			return nil, err
		}
		findings = append(findings, projectFindings...)
	}

	// Directory order is filesystem-dependent, so sorting keeps hook output the
	// same from one run to the next.
	sort.Slice(findings, func(first, second int) bool {
		if findings[first].Path != findings[second].Path {
			return findings[first].Path < findings[second].Path
		}
		return findings[first].Problem < findings[second].Problem
	})
	return findings, nil
}

// findProjectFiles walks the repository for project configuration, returning
// paths relative to root so findings read the way a contributor types them.
func findProjectFiles(root string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if skippedDirectories[entry.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.Name() != projectFile {
			return nil
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		paths = append(paths, filepath.ToSlash(relative))
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("search for %s files: %w", projectFile, err)
	}
	return paths, nil
}

// checkProject reads one project file and applies the testing policy to it.
func checkProject(root string, relativePath string) ([]Finding, error) {
	contents, err := os.ReadFile(filepath.Join(root, relativePath))
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", relativePath, err)
	}

	var project configuration
	if err := json.Unmarshal(contents, &project); err != nil {
		// A file that cannot be parsed has not been checked. Failing is honest;
		// skipping it would report a pass the check never established.
		return nil, fmt.Errorf("parse %s: %w", relativePath, err)
	}

	name := project.Name
	if name == "" {
		name = filepath.ToSlash(filepath.Dir(relativePath))
	}

	var findings []Finding
	report := func(problem string) {
		findings = append(findings, Finding{Project: name, Path: relativePath, Problem: problem})
	}

	quick, hasQuick := project.Targets[quickTarget]
	if !hasQuick {
		report(fmt.Sprintf("defines no %s target", quickTarget))
	}
	for _, required := range requiredDependencies {
		if _, exists := project.Targets[required]; !exists {
			report(fmt.Sprintf("defines no %s target", required))
		}
	}

	if !hasQuick {
		// Every remaining rule describes the quick-test target. Without one there
		// is nothing further to say, and saying it twice buries the real repair.
		return findings, nil
	}

	if quick.Cache != nil && !*quick.Cache {
		report(fmt.Sprintf("turns off the cache on %s, which pre-push relies on", quickTarget))
	}

	declared := selfDependencies(quick.DependsOn)
	for _, required := range requiredDependencies {
		if !declared[required] {
			report(fmt.Sprintf("does not run %s from %s through dependsOn", required, quickTarget))
		}
	}
	return findings, nil
}

// selfDependencies collects the dependsOn entries that name a target in this
// same project. An entry written as an object runs a target somewhere else, and
// a "^" prefix runs it in this project's dependencies, so neither satisfies a
// requirement about this project's own quick tests.
func selfDependencies(entries []json.RawMessage) map[string]bool {
	declared := map[string]bool{}
	for _, entry := range entries {
		var name string
		if err := json.Unmarshal(entry, &name); err != nil {
			continue
		}
		if strings.HasPrefix(name, "^") {
			continue
		}
		declared[name] = true
	}
	return declared
}
