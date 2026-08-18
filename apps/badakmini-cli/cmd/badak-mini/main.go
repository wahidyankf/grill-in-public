// Badak Mini is a focused repository-governance command-line tool.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/wahidyankf/grill-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grill-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grill-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grill-in-public/apps/badakmini-cli/internal/rulechange"
)

// Each supported invocation is spelled out, so an unknown command fails with
// usage rather than falling through to a check the caller did not ask for.
const (
	instructionSizeCommand    = "harness instruction-size validate"
	markdownLinksCommand      = "harness markdown-links validate"
	ruleChangeValidateCommand = "harness rule-change validate"
	ruleChangeHookCommand     = "harness rule-change hook"
	capabilityParityCommand   = "harness capability-parity validate"
)

var supportedCommands = []string{
	instructionSizeCommand,
	markdownLinksCommand,
	ruleChangeValidateCommand,
	ruleChangeHookCommand,
	capabilityParityCommand,
}

const usage = `Usage:
  badak-mini harness instruction-size validate
  badak-mini harness markdown-links validate
  badak-mini harness rule-change validate
  badak-mini harness rule-change hook
  badak-mini harness capability-parity validate

Validate governance Markdown word limits, repository-local Markdown links, or
the subagents, skills, and commands every harness exposes, or announce the
propagate-rules workflow when a rule changes. The rule-change validate form
reads the staged paths; its hook form reads a pre-edit payload on stdin.
`

func main() {
	// Keep run independent from process globals so command parsing and exit
	// behavior can be tested without spawning a subprocess.
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, findRepositoryRoot))
}

func run(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
	rootFinder func() (string, error),
) int {
	// Help is successful even outside a repository because it does not need to
	// inspect files; validation commands must establish the repository boundary.
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		if err := writef(stdout, usage); err != nil {
			return 1
		}
		return 0
	}

	if !matchesCommand(args) {
		// Exit status 2 distinguishes an invalid invocation from a failed check.
		if err := writef(stderr, usage); err != nil {
			return 1
		}
		return 2
	}

	root, err := rootFinder()
	if err != nil {
		if writeErr := writef(stderr, "ERROR: could not find the Git repository root: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}

	// Command matching stays explicit while the CLI has a handful of narrowly
	// scoped workflows. Add a dedicated branch when a command needs new output.
	switch strings.Join(args, " ") {
	case instructionSizeCommand:
		return validateInstructionSize(root, stdout, stderr)
	case ruleChangeValidateCommand:
		return announceStagedRuleChange(root, stdout, stderr)
	case ruleChangeHookCommand:
		return announceHookRuleChange(root, os.Stdin, stdout)
	case capabilityParityCommand:
		return validateCapabilityParity(root, stdout, stderr)
	}
	return validateMarkdownLinks(root, stdout, stderr)
}

// announceStagedRuleChange reports a staged rule change to a contributor. It
// succeeds either way: the workflow is the author's next step, not a gate that
// a hook can decide has been satisfied.
func announceStagedRuleChange(root string, stdout io.Writer, stderr io.Writer) int {
	staged, err := rulechange.StagedPaths(root)
	if err != nil {
		if writeErr := writef(stderr, "ERROR: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}

	paths := rulechange.RulePaths(staged)
	if len(paths) == 0 {
		return 0
	}

	if err := writef(stdout, "%s\n", rulechange.Notice(paths)); err != nil {
		return 1
	}
	return 0
}

// announceHookRuleChange answers a harness pre-edit hook. It returns the notice
// as additional context so the agent loads the workflow before it edits, and it
// stays silent for every other file so ordinary work is never interrupted.
func announceHookRuleChange(root string, stdin io.Reader, stdout io.Writer) int {
	payload, err := io.ReadAll(stdin)
	if err != nil {
		// A hook that cannot read its payload must not block the edit.
		return 0
	}

	paths := rulechange.RulePaths(rulechange.HookPaths(payload, root))
	if len(paths) == 0 {
		return 0
	}

	response := hookResponse{}
	response.HookSpecificOutput.HookEventName = "PreToolUse"
	response.HookSpecificOutput.AdditionalContext = rulechange.Notice(paths)
	response.SystemMessage = rulechange.Notice(paths)

	encoded, err := json.Marshal(response)
	if err != nil {
		return 0
	}
	if err := writef(stdout, "%s\n", encoded); err != nil {
		return 0
	}
	return 0
}

// hookResponse is the pre-edit hook reply shape: additional context informs the
// agent without denying the edit, and the system message tells the human.
type hookResponse struct {
	HookSpecificOutput struct {
		HookEventName     string `json:"hookEventName"`
		AdditionalContext string `json:"additionalContext"`
	} `json:"hookSpecificOutput"`
	SystemMessage string `json:"systemMessage"`
}

func validateInstructionSize(root string, stdout io.Writer, stderr io.Writer) int {
	findings, err := governance.Check(root)
	if err != nil {
		if writeErr := writef(stderr, "ERROR: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		if err := writef(stdout, "Governance word counts are within the %d-word limit.\n", governance.MaxWords); err != nil {
			return 1
		}
		return 0
	}
	for _, finding := range findings {
		// Report every over-limit document in one run so a contributor can repair
		// the complete guidance set before retrying the hook.
		if err := writef(stderr, "ERROR: %s contains %d words; the limit is %d.\n", finding.Path, finding.WordCount, governance.MaxWords); err != nil {
			return 1
		}
	}
	if err := writef(stderr, "Use progressive disclosure: split detailed guidance into focused files.\n"); err != nil {
		return 1
	}
	return 1
}

// validateCapabilityParity fails when one harness exposes a capability another
// supporting harness lacks, because an agent that exists for one tool and not
// the next makes the repository behave differently depending on who runs it.
func validateCapabilityParity(root string, stdout io.Writer, stderr io.Writer) int {
	findings, err := parity.Check(root)
	if err != nil {
		if writeErr := writef(stderr, "ERROR: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		if err := writef(stdout, "Every harness exposes the same subagents, skills, and commands.\n"); err != nil {
			return 1
		}
		for _, note := range parity.UnsupportedNotes() {
			// Naming the exemption keeps an absent harness from reading as an
			// oversight the next time someone compares the directories.
			if err := writef(stdout, "Exempt, %s.\n", note); err != nil {
				return 1
			}
		}
		return 0
	}
	for _, finding := range findings {
		if err := writef(stderr, "ERROR: %s\n", finding.Message()); err != nil {
			return 1
		}
	}
	if err := writef(stderr, "See repo-governance/conventions/harness-capability-parity-policy.md.\n"); err != nil {
		return 1
	}
	return 1
}

func validateMarkdownLinks(root string, stdout io.Writer, stderr io.Writer) int {
	findings, err := markdownlinks.Check(root)
	if err != nil {
		if writeErr := writef(stderr, "ERROR: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		if err := writef(stdout, "Repository-local Markdown links are valid.\n"); err != nil {
			return 1
		}
		return 0
	}
	for _, finding := range findings {
		// Source path and one-based line number make a hook failure actionable in
		// a terminal or CI log without requiring a separate report file.
		if err := writef(stderr, "ERROR: %s:%d: %q %s.\n", finding.Path, finding.Line, finding.Destination, finding.Problem); err != nil {
			return 1
		}
	}
	return 1
}

// writef propagates output failures so commands do not report success when a
// caller cannot receive their validation result.
func writef(writer io.Writer, format string, arguments ...any) error {
	_, err := fmt.Fprintf(writer, format, arguments...)
	return err
}

func matchesCommand(args []string) bool {
	return slices.Contains(supportedCommands, strings.Join(args, " "))
}

func findRepositoryRoot() (string, error) {
	// Git, rather than the current directory, defines the validation boundary:
	// callers may run Badak Mini from any nested path in this repository.
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", errors.New("run this command from inside a Git repository")
	}

	return strings.TrimSpace(string(output)), nil
}
