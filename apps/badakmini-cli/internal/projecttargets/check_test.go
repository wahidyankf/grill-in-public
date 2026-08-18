package projecttargets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// completeProject is the target set the testing policy requires. Tests start
// from it and remove one piece, so each case names a single missing rule.
const completeProject = `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"},
    "test:quick": {"dependsOn": ["typecheck", "lint"], "command": "echo test"}
  }
}`

func TestCheckAcceptsAProjectWithEveryRequiredTarget(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), completeProject)

	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsAProjectWithoutQuickTests(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"}
  }
}`)

	findings := checkOrFail(t, root)
	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	if findings[0].Project != "example" || !strings.Contains(findings[0].Problem, "test:quick") {
		t.Fatalf("unexpected finding: %#v", findings[0])
	}
}

func TestCheckReportsAProjectWithoutLintOrTypecheck(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "test:quick": {"command": "echo test"}
  }
}`)

	findings := checkOrFail(t, root)
	// Adding each target and naming it in dependsOn are separate repairs, so both
	// are reported for both targets and one pass is enough to finish the project.
	if len(findings) != 4 {
		t.Fatalf("expected a missing-target and a dependsOn finding for each, got %#v", findings)
	}
}

func TestCheckReportsQuickTestsThatDoNotDependOnLintAndTypecheck(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"},
    "test:quick": {"dependsOn": ["typecheck"], "command": "echo test"}
  }
}`)

	findings := checkOrFail(t, root)
	if len(findings) != 1 || !strings.Contains(findings[0].Problem, "lint") {
		t.Fatalf("expected the missing lint dependency to be reported, got %#v", findings)
	}
}

func TestCheckIgnoresADependsOnEntryThatPointsAtAnotherProject(t *testing.T) {
	// An object entry runs a target in a different project. It does not lint or
	// typecheck this one, so accepting it would let the requirement be dodged.
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"},
    "test:quick": {
      "dependsOn": ["lint", {"projects": ["other"], "target": "typecheck"}],
      "command": "echo test"
    }
  }
}`)

	findings := checkOrFail(t, root)
	if len(findings) != 1 || !strings.Contains(findings[0].Problem, "typecheck") {
		t.Fatalf("expected the cross-project entry to be rejected, got %#v", findings)
	}
}

func TestCheckAcceptsExtraQuickTestDependencies(t *testing.T) {
	// A compiled project runs its tests against build output, so build is a
	// legitimate addition rather than a deviation from the policy.
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"},
    "build": {"command": "echo build"},
    "test:quick": {"dependsOn": ["build", "typecheck", "lint"], "command": "echo test"}
  }
}`)

	findings := checkOrFail(t, root)
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsQuickTestsThatOptOutOfTheCache(t *testing.T) {
	// Cacheability comes from the workspace target defaults, so a project only
	// loses it by saying so. Pre-push depends on that cache to stay quick.
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), `{
  "name": "example",
  "targets": {
    "lint": {"command": "echo lint"},
    "typecheck": {"command": "echo typecheck"},
    "test:quick": {"dependsOn": ["typecheck", "lint"], "cache": false, "command": "echo test"}
  }
}`)

	findings := checkOrFail(t, root)
	if len(findings) != 1 || !strings.Contains(findings[0].Problem, "cache") {
		t.Fatalf("expected the disabled cache to be reported, got %#v", findings)
	}
}

func TestCheckSkipsGeneratedAndVendoredDirectories(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), completeProject)
	writeFile(t, filepath.Join(root, "node_modules/vendored/project.json"), `{"name": "vendored"}`)
	writeFile(t, filepath.Join(root, "libs/example/dist/project.json"), `{"name": "copied"}`)

	findings := checkOrFail(t, root)
	if len(findings) != 0 {
		t.Fatalf("expected vendored and built copies to be skipped, got %#v", findings)
	}
}

func TestCheckReportsEveryProjectRatherThanTheFirst(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "apps/one/project.json"), `{"name": "one", "targets": {}}`)
	writeFile(t, filepath.Join(root, "apps/two/project.json"), `{"name": "two", "targets": {}}`)

	findings := checkOrFail(t, root)
	projects := map[string]bool{}
	for _, finding := range findings {
		projects[finding.Project] = true
	}
	if !projects["one"] || !projects["two"] {
		t.Fatalf("expected both projects to be reported, got %#v", findings)
	}
}

func TestCheckFailsOnAProjectFileItCannotParse(t *testing.T) {
	// Unreadable configuration is not a passing project. Reporting it as an
	// error stops the check from certifying a file it never inspected.
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "libs/example/project.json"), "{not json")

	if _, err := Check(root); err == nil {
		t.Fatal("expected a parse failure to be reported")
	}
}

func TestFindingMessageNamesThePathAndProblem(t *testing.T) {
	message := Finding{
		Project: "example",
		Path:    "libs/example/project.json",
		Problem: "defines no test:quick target",
	}.Message()

	if !strings.Contains(message, "libs/example/project.json") || !strings.Contains(message, "test:quick") {
		t.Fatalf("expected the path and problem, got %q", message)
	}
}

func checkOrFail(t *testing.T, root string) []Finding {
	t.Helper()
	findings, err := Check(root)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	return findings
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
