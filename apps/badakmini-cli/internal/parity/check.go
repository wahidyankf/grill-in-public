// Package parity compares the capabilities each harness exposes, so a subagent,
// skill, or command added for one tool cannot silently stay missing elsewhere.
package parity

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// skillManifest marks a skill directory. A skill is a directory rather than a
// file, so its name comes from the directory holding this manifest.
const skillManifest = "SKILL.md"

// indexName is each directory's own index. It documents the directory instead of
// defining a capability, so counting it would report a difference that is not one.
const indexName = "README"

// source is one directory a harness loads a capability from, together with the
// file extension that defines an entry there. An empty extension means the
// entries are skill directories.
type source struct {
	directory string
	extension string
}

// surface is what one harness exposes for one capability. A harness can read
// several directories, so the names it offers are their union.
type surface struct {
	harness string
	sources []source
}

// capability is a repository-wide feature and the harnesses that can load it. A
// harness that cannot load it at all is left out rather than recorded as empty,
// because an unsupported tool is exempt from the count.
type capability struct {
	name      string
	unsupport string
	surfaces  []surface
}

// capabilities encodes the directories recorded in the harness capability parity
// policy. Change them together, or the check stops describing the repository.
var capabilities = []capability{
	{
		name: "subagent",
		surfaces: []surface{
			{harness: "Claude Code", sources: []source{{directory: ".claude/agents", extension: ".md"}}},
			{harness: "Codex", sources: []source{{directory: ".codex/agents", extension: ".toml"}}},
			{harness: "opencode", sources: []source{{directory: ".opencode/agents", extension: ".md"}}},
		},
	},
	{
		name: "skill",
		surfaces: []surface{
			{harness: "Claude Code", sources: []source{{directory: ".claude/skills"}}},
			{harness: "Codex", sources: []source{{directory: ".agents/skills"}}},
			// opencode reads its own directory and both shared ones, so a skill
			// mirrored for the other two harnesses already reaches it.
			{harness: "opencode", sources: []source{
				{directory: ".opencode/skills"},
				{directory: ".claude/skills"},
				{directory: ".agents/skills"},
			}},
		},
	},
	{
		name:      "command",
		unsupport: "Codex has no project command directory",
		surfaces: []surface{
			// Only the explicit command directories count here. Claude Code also
			// answers a skill as `/name`, but treating that as a command would
			// demand an opencode command for every shared skill.
			{harness: "Claude Code", sources: []source{{directory: ".claude/commands", extension: ".md"}}},
			{harness: "opencode", sources: []source{{directory: ".opencode/commands", extension: ".md"}}},
		},
	},
}

// Finding describes one harness missing entries its peers already expose.
type Finding struct {
	Capability string
	Harness    string
	Missing    []string
}

// Message states the difference in the terms the parity policy uses.
func (finding Finding) Message() string {
	return fmt.Sprintf(
		"%s parity: %s is missing %s %s, which another harness exposes.",
		finding.Capability,
		finding.Harness,
		strings.Join(finding.Missing, ", "),
		pluralize(finding.Capability, len(finding.Missing)),
	)
}

// Check compares every capability across the harnesses that support it. It
// returns each difference rather than the first, so one pass can fix them all.
func Check(root string) ([]Finding, error) {
	var findings []Finding

	for _, capability := range capabilities {
		exposed := make(map[string][]string, len(capability.surfaces))
		union := map[string]struct{}{}

		for _, surface := range capability.surfaces {
			names, err := surfaceNames(root, surface)
			if err != nil {
				return nil, err
			}
			exposed[surface.harness] = names
			for _, name := range names {
				union[name] = struct{}{}
			}
		}

		// A capability no harness uses yet is not a difference. Reporting it
		// would demand entries the repository has not decided to have.
		if len(union) == 0 {
			continue
		}

		for _, surface := range capability.surfaces {
			missing := missingFrom(union, exposed[surface.harness])
			if len(missing) == 0 {
				continue
			}
			findings = append(findings, Finding{
				Capability: capability.name,
				Harness:    surface.harness,
				Missing:    missing,
			})
		}
	}

	return findings, nil
}

// UnsupportedNotes lists the harnesses exempt from a capability, so a report can
// say why a tool is absent instead of leaving its absence to be guessed at.
func UnsupportedNotes() []string {
	var notes []string
	for _, capability := range capabilities {
		if capability.unsupport != "" {
			notes = append(notes, fmt.Sprintf("%s: %s", capability.name, capability.unsupport))
		}
	}
	return notes
}

func surfaceNames(root string, surface surface) ([]string, error) {
	var names []string
	for _, source := range surface.sources {
		sourceNames, err := sourceEntries(root, source)
		if err != nil {
			return nil, err
		}
		for _, name := range sourceNames {
			if !slices.Contains(names, name) {
				names = append(names, name)
			}
		}
	}
	sort.Strings(names)
	return names, nil
}

func sourceEntries(root string, source source) ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(root, source.directory))
	if err != nil {
		// A directory a harness never received is the same as an empty one for
		// this comparison, and the difference is reported against its peers.
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read %s: %w", source.directory, err)
	}

	var names []string
	for _, entry := range entries {
		name, ok := entryName(root, source, entry)
		if !ok {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}

func entryName(root string, source source, entry os.DirEntry) (string, bool) {
	if source.extension == "" {
		if !entry.IsDir() {
			return "", false
		}
		// Only a directory carrying the manifest is a skill; anything else is
		// supporting material.
		if _, err := os.Stat(filepath.Join(root, source.directory, entry.Name(), skillManifest)); err != nil {
			return "", false
		}
		return entry.Name(), true
	}

	if entry.IsDir() || filepath.Ext(entry.Name()) != source.extension {
		return "", false
	}
	name := strings.TrimSuffix(entry.Name(), source.extension)
	if strings.EqualFold(name, indexName) {
		return "", false
	}
	return name, true
}

func missingFrom(union map[string]struct{}, exposed []string) []string {
	var missing []string
	for name := range union {
		if !slices.Contains(exposed, name) {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	return missing
}

func pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}
