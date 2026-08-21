package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunPrintsHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// A failing finder proves help returns before it tries to locate Git.
	exitCode := run([]string{"--help"}, &stdout, &stderr, func() (string, error) {
		return "", errors.New("root lookup should not run")
	})

	if exitCode != 0 {
		t.Fatalf("expected successful help exit, got %d", exitCode)
	}
	hasInstructionCommand := strings.Contains(stdout.String(), "harness instruction-size validate")
	hasMarkdownCommand := strings.Contains(stdout.String(), "harness markdown-links validate")
	if !hasInstructionCommand || !hasMarkdownCommand {
		t.Fatalf("expected command usage, got %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no error output, got %q", stderr.String())
	}
	// Help is how a contributor discovers the checks, so every supported command
	// must appear there rather than only in the dispatch switch.
	for _, command := range supportedCommands {
		if !strings.Contains(stdout.String(), command) {
			t.Fatalf("expected %q in the usage text, got %q", command, stdout.String())
		}
	}
}

func TestRunRejectsUnsupportedCommands(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// The same injected failure keeps this test focused on argument validation.
	exitCode := run([]string{"check-governance"}, &stdout, &stderr, func() (string, error) {
		return "", errors.New("root lookup should not run")
	})

	if exitCode != 2 {
		t.Fatalf("expected usage error exit, got %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "Usage:") {
		t.Fatalf("expected usage output, got %q", stderr.String())
	}
}

func TestRunFailsWhenHelpCannotBeWritten(t *testing.T) {
	var stderr bytes.Buffer

	// A failing writer proves a successful command never masks the fact that it
	// could not deliver its result to the caller.
	exitCode := run([]string{"--help"}, writeFailure{}, &stderr, func() (string, error) {
		return "", errors.New("root lookup should not run")
	})

	if exitCode != 1 {
		t.Fatalf("expected output failure exit, got %d", exitCode)
	}
}

func TestRunFailsWhenInvalidUsageCannotBeWritten(t *testing.T) {
	exitCode := run([]string{"unknown"}, io.Discard, writeFailure{}, func() (string, error) {
		return "", errors.New("root lookup should not run")
	})

	if exitCode != 1 {
		t.Fatalf("expected output failure exit, got %d", exitCode)
	}
}

func TestRunReportsRepositoryDiscoveryFailure(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run([]string{"harness", "markdown-links", "validate"}, &stdout, &stderr, func() (string, error) {
		return "", errors.New("not a repository")
	})

	if exitCode != 1 || !strings.Contains(stderr.String(), "could not find") {
		t.Fatalf("expected repository discovery failure, got exit %d and %q", exitCode, stderr.String())
	}
}

func TestRunDispatchesEveryValidationCommand(t *testing.T) {
	testCases := []struct {
		name       string
		command    string
		prepare    func(*testing.T) string
		expectText string
	}{
		{
			name:       "instruction size",
			command:    instructionSizeCommand,
			prepare:    newGovernanceRepository,
			expectText: "Governance word counts",
		},
		{
			name:       "Markdown links",
			command:    markdownLinksCommand,
			prepare:    newTrackedMarkdownRepository,
			expectText: "Markdown links are valid",
		},
		{
			name:       "staged rule change",
			command:    ruleChangeValidateCommand,
			prepare:    newStagedRuleRepository,
			expectText: "Rule change detected",
		},
		{
			name:       "capability parity",
			command:    capabilityParityCommand,
			prepare:    newEmptyRepository,
			expectText: "Every harness exposes",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			root := testCase.prepare(t)
			exitCode, stdout, stderr := runAtRoot(strings.Fields(testCase.command), root)
			if exitCode != 0 || !strings.Contains(stdout, testCase.expectText) {
				t.Fatalf("expected successful %s, got exit %d, stdout %q, stderr %q", testCase.name, exitCode, stdout, stderr)
			}
		})
	}
}

func TestAnnounceStagedRuleChangeStaysSilentWithoutStagedRules(t *testing.T) {
	root := newGitRepository(t)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := announceStagedRuleChange(root, &stdout, &stderr)
	if exitCode != 0 || stdout.Len() != 0 || stderr.Len() != 0 {
		t.Fatalf("expected a silent success, got exit %d, stdout %q, stderr %q", exitCode, stdout.String(), stderr.String())
	}
}

func TestAnnounceStagedRuleChangeReportsGitFailure(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := announceStagedRuleChange(t.TempDir(), &stdout, &stderr)

	if exitCode != 1 || !strings.Contains(stderr.String(), "list staged paths") {
		t.Fatalf("expected Git failure, got exit %d and %q", exitCode, stderr.String())
	}
}

func TestAnnounceStagedRuleChangeReportsWriteFailure(t *testing.T) {
	root := newStagedRuleRepository(t)
	exitCode := announceStagedRuleChange(root, writeFailure{}, io.Discard)
	if exitCode != 1 {
		t.Fatalf("expected output failure exit, got %d", exitCode)
	}
}

func TestAnnounceHookRuleChangeHandlesRelevantAndOrdinaryPayloads(t *testing.T) {
	root := t.TempDir()
	relevant := fmt.Sprintf(`{"tool_input":{"file_path":%q}}`, filepath.Join(root, "AGENTS.md"))
	var stdout bytes.Buffer

	exitCode := announceHookRuleChange(root, strings.NewReader(relevant), &stdout)
	if exitCode != 0 || !strings.Contains(stdout.String(), "PreToolUse") {
		t.Fatalf("expected hook context, got exit %d and %q", exitCode, stdout.String())
	}

	stdout.Reset()
	exitCode = announceHookRuleChange(root, strings.NewReader(`{"tool_input":{"file_path":"README.md"}}`), &stdout)
	if exitCode != 0 || stdout.Len() != 0 {
		t.Fatalf("expected ordinary edit silence, got exit %d and %q", exitCode, stdout.String())
	}
}

func TestAnnounceHookRuleChangeNeverBlocksOnIOFailure(t *testing.T) {
	if exitCode := announceHookRuleChange(t.TempDir(), readFailure{}, io.Discard); exitCode != 0 {
		t.Fatalf("expected input failure to stay nonblocking, got %d", exitCode)
	}

	payload := `{"tool_input":{"file_path":"AGENTS.md"}}`
	if exitCode := announceHookRuleChange(t.TempDir(), strings.NewReader(payload), writeFailure{}); exitCode != 0 {
		t.Fatalf("expected output failure to stay nonblocking, got %d", exitCode)
	}
}

func TestValidationCommandsReportFindings(t *testing.T) {
	t.Run("instruction size", func(t *testing.T) {
		root := newGovernanceRepository(t)
		writeTestFile(t, filepath.Join(root, "AGENTS.md"), strings.Repeat("word ", 501))
		assertValidationFailure(t, func(stdout, stderr io.Writer) int {
			return validateInstructionSize(root, stdout, stderr)
		}, "contains 501 words")
	})

	t.Run("Markdown links", func(t *testing.T) {
		root := newGitRepository(t)
		writeTestFile(t, filepath.Join(root, "README.md"), "[Missing](missing.md)\n")
		runGit(t, root, "add", "README.md")
		assertValidationFailure(t, func(stdout, stderr io.Writer) int {
			return validateMarkdownLinks(root, stdout, stderr)
		}, "does not exist")
	})

	t.Run("capability parity", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, ".claude/agents/review.md"), "review")
		assertValidationFailure(t, func(stdout, stderr io.Writer) int {
			return validateCapabilityParity(root, stdout, stderr)
		}, "parity")
	})
}

func TestValidationCommandsReportInputAndOutputFailures(t *testing.T) {
	t.Run("instruction input", func(t *testing.T) {
		assertFailureExitCode(t, validateInstructionSize(t.TempDir(), io.Discard, io.Discard))
	})

	t.Run("instruction output", func(t *testing.T) {
		assertFailureExitCode(t, validateInstructionSize(newGovernanceRepository(t), writeFailure{}, io.Discard))
	})

	t.Run("capability input", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, ".claude/agents"), "not a directory")
		assertFailureExitCode(t, validateCapabilityParity(root, io.Discard, io.Discard))
	})

	t.Run("capability output", func(t *testing.T) {
		assertFailureExitCode(t, validateCapabilityParity(t.TempDir(), writeFailure{}, io.Discard))
	})

	t.Run("Markdown input", func(t *testing.T) {
		assertFailureExitCode(t, validateMarkdownLinks(t.TempDir(), io.Discard, io.Discard))
	})

	t.Run("Markdown output", func(t *testing.T) {
		assertFailureExitCode(t, validateMarkdownLinks(newTrackedMarkdownRepository(t), writeFailure{}, io.Discard))
	})
}

func TestFindRepositoryRootFindsThisCheckout(t *testing.T) {
	root, err := findRepositoryRoot()
	if err != nil {
		t.Fatalf("expected repository root, got %v", err)
	}
	if filepath.Base(root) != "grind-in-public" {
		t.Fatalf("expected grind-in-public root, got %q", root)
	}
}

func TestFindRepositoryRootRejectsANonRepository(t *testing.T) {
	t.Chdir(t.TempDir())
	if _, err := findRepositoryRoot(); err == nil {
		t.Fatal("expected repository discovery to fail")
	}
}

type writeFailure struct{}

func (writeFailure) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

type readFailure struct{}

func (readFailure) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func runAtRoot(args []string, root string) (int, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run(args, &stdout, &stderr, func() (string, error) { return root, nil })

	return exitCode, stdout.String(), stderr.String()
}

func assertValidationFailure(
	t *testing.T,
	validate func(io.Writer, io.Writer) int,
	expected string,
) {
	t.Helper()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := validate(&stdout, &stderr)
	if exitCode != 1 || !strings.Contains(stderr.String(), expected) {
		t.Fatalf(
			"expected failure containing %q, got exit %d, stdout %q, stderr %q",
			expected,
			exitCode,
			stdout.String(),
			stderr.String(),
		)
	}
}

func assertFailureExitCode(t *testing.T, actual int) {
	t.Helper()
	if actual != 1 {
		t.Fatalf("expected exit 1, got %d", actual)
	}
}

func newEmptyRepository(t *testing.T) string {
	t.Helper()

	return t.TempDir()
}

func newGovernanceRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeTestFile(t, filepath.Join(root, "AGENTS.md"), "# Agents\n")
	writeTestFile(t, filepath.Join(root, "CLAUDE.md"), "# Claude\n")
	writeTestFile(t, filepath.Join(root, "repo-governance/README.md"), "# Governance\n")

	return root
}

func newTrackedMarkdownRepository(t *testing.T) string {
	t.Helper()
	root := newGitRepository(t)
	writeTestFile(t, filepath.Join(root, "README.md"), "# Repository\n")
	runGit(t, root, "add", "README.md")

	return root
}

func newStagedRuleRepository(t *testing.T) string {
	t.Helper()
	root := newGitRepository(t)
	writeTestFile(t, filepath.Join(root, "AGENTS.md"), "# Agents\n")
	runGit(t, root, "add", "AGENTS.md")

	return root
}

func newGitRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	runGit(t, root, "init", "--quiet")

	return root
}

func runGit(t *testing.T, root string, arguments ...string) {
	t.Helper()
	commandArguments := append([]string{"-C", root}, arguments...)
	// #nosec G204 -- tests control the executable, temporary root, and every argument.
	command := exec.Command("git", commandArguments...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("run git %v: %s: %v", arguments, output, err)
	}
}

func writeTestFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("create directory for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
