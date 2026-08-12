// Badak Mini is a focused repository-governance command-line tool.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/wahidyankf/swe-grilling/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/swe-grilling/apps/badakmini-cli/internal/markdownlinks"
)

const usage = `Usage:
  badak-mini harness instruction-size validate
  badak-mini harness markdown-links validate

Validate governance Markdown word limits or repository-local Markdown links.
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

	// Command matching stays explicit while the CLI has two narrowly scoped
	// workflows. Add a dedicated branch when a future command needs new output.
	if matchesInstructionSizeCommand(args) {
		return validateInstructionSize(root, stdout, stderr)
	}
	return validateMarkdownLinks(root, stdout, stderr)
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

func matchesInstructionSizeCommand(args []string) bool {
	return strings.Join(args, " ") == "harness instruction-size validate"
}

func matchesCommand(args []string) bool {
	command := strings.Join(args, " ")
	return command == "harness instruction-size validate" || command == "harness markdown-links validate"
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
