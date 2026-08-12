// Badak Mini is a focused repository-governance command-line tool.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/wahidyankf/swe-grilling/apps/badak-mini/internal/governance"
	"github.com/wahidyankf/swe-grilling/apps/badak-mini/internal/markdownlinks"
)

const usage = `Usage:
  badak-mini harness instruction-size validate
  badak-mini harness markdown-links validate

Validate governance Markdown word limits or repository-local Markdown links.
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, findRepositoryRoot))
}

func run(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
	rootFinder func() (string, error),
) int {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		fmt.Fprint(stdout, usage)
		return 0
	}

	if !matchesCommand(args) {
		fmt.Fprint(stderr, usage)
		return 2
	}

	root, err := rootFinder()
	if err != nil {
		fmt.Fprintf(stderr, "ERROR: could not find the Git repository root: %v\n", err)
		return 1
	}

	if matchesInstructionSizeCommand(args) {
		return validateInstructionSize(root, stdout, stderr)
	}
	return validateMarkdownLinks(root, stdout, stderr)
}

func validateInstructionSize(root string, stdout io.Writer, stderr io.Writer) int {
	findings, err := governance.Check(root)
	if err != nil {
		fmt.Fprintf(stderr, "ERROR: %v\n", err)
		return 1
	}
	if len(findings) == 0 {
		fmt.Fprintf(stdout, "Governance word counts are within the %d-word limit.\n", governance.MaxWords)
		return 0
	}
	for _, finding := range findings {
		fmt.Fprintf(stderr, "ERROR: %s contains %d words; the limit is %d.\n", finding.Path, finding.WordCount, governance.MaxWords)
	}
	fmt.Fprintln(stderr, "Use progressive disclosure: split detailed guidance into focused files.")
	return 1
}

func validateMarkdownLinks(root string, stdout io.Writer, stderr io.Writer) int {
	findings, err := markdownlinks.Check(root)
	if err != nil {
		fmt.Fprintf(stderr, "ERROR: %v\n", err)
		return 1
	}
	if len(findings) == 0 {
		fmt.Fprintln(stdout, "Repository-local Markdown links are valid.")
		return 0
	}
	for _, finding := range findings {
		fmt.Fprintf(stderr, "ERROR: %s:%d: %q %s.\n", finding.Path, finding.Line, finding.Destination, finding.Problem)
	}
	return 1
}

func matchesInstructionSizeCommand(args []string) bool {
	return strings.Join(args, " ") == "harness instruction-size validate"
}

func matchesCommand(args []string) bool {
	command := strings.Join(args, " ")
	return command == "harness instruction-size validate" || command == "harness markdown-links validate"
}

func findRepositoryRoot() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", errors.New("run this command from inside a Git repository")
	}

	return strings.TrimSpace(string(output)), nil
}
