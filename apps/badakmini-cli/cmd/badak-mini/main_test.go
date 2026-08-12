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
}

func TestRunRejectsUnsupportedCommands(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

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
