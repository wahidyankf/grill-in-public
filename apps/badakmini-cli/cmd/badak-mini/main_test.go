package main

import (
	"bytes"
	"errors"
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
	if !strings.Contains(stdout.String(), "harness instruction-size validate") || !strings.Contains(stdout.String(), "harness markdown-links validate") {
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

type writeFailure struct{}

func (writeFailure) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}
