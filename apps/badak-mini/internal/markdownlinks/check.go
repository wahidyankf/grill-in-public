// Package markdownlinks validates repository-local Markdown links.
package markdownlinks

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const markdownExtension = ".md"

var referenceDefinitionPattern = regexp.MustCompile(`^ {0,3}\[([^\]]+)\]:\s*(?:<([^>]+)>|(\S+))`)

// Finding describes an invalid repository-local link.
type Finding struct {
	Path        string
	Line        int
	Destination string
	Problem     string
}

type link struct {
	destination string
	line        int
}

// Check validates every Git-tracked Markdown file beneath root. Using Git's
// tracked tree excludes metadata, dependency installs, and generated output,
// while still detecting links left dangling by a committed file deletion.
// External URLs are deliberately ignored.
func Check(root string) ([]Finding, error) {
	trackedFiles, err := findTrackedFiles(root)
	if err != nil {
		return nil, err
	}
	markdownFiles := filterMarkdownFiles(trackedFiles)

	findings := make([]Finding, 0)
	for _, sourcePath := range markdownFiles {
		contents, err := os.ReadFile(filepath.Join(root, sourcePath))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", sourcePath, err)
		}

		for _, candidate := range extractLinks(string(contents)) {
			finding := checkLink(root, sourcePath, candidate, trackedFiles)
			if finding != nil {
				findings = append(findings, *finding)
			}
		}
	}

	sort.Slice(findings, func(left, right int) bool {
		if findings[left].Path != findings[right].Path {
			return findings[left].Path < findings[right].Path
		}
		if findings[left].Line != findings[right].Line {
			return findings[left].Line < findings[right].Line
		}
		return findings[left].Destination < findings[right].Destination
	})

	return findings, nil
}

func findTrackedFiles(root string) (map[string]struct{}, error) {
	output, err := exec.Command("git", "-C", root, "ls-files", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("list tracked repository files: %w", err)
	}
	paths := strings.Split(strings.TrimSuffix(string(output), "\x00"), "\x00")
	files := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		if path != "" {
			files[filepath.ToSlash(path)] = struct{}{}
		}
	}
	return files, nil
}

func filterMarkdownFiles(trackedFiles map[string]struct{}) []string {
	files := make([]string, 0)
	for path := range trackedFiles {
		if strings.EqualFold(filepath.Ext(path), markdownExtension) {
			files = append(files, path)
		}
	}
	sort.Strings(files)
	return files
}

func extractLinks(contents string) []link {
	lines := strings.Split(contents, "\n")
	references := make(map[string]string)
	insideFence := false
	for _, line := range lines {
		if isFence(line) {
			insideFence = !insideFence
			continue
		}
		if insideFence {
			continue
		}
		matches := referenceDefinitionPattern.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}
		destination := matches[2]
		if destination == "" {
			destination = matches[3]
		}
		references[normalizeReference(matches[1])] = destination
	}

	links := make([]link, 0)
	insideFence = false
	for index, line := range lines {
		if isFence(line) {
			insideFence = !insideFence
			continue
		}
		if insideFence || referenceDefinitionPattern.MatchString(line) {
			continue
		}
		links = append(links, extractLineLinks(line, index+1, references)...)
	}
	return links
}

func isFence(line string) bool {
	trimmed := strings.TrimLeft(line, " ")
	return strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~")
}

func extractLineLinks(line string, lineNumber int, references map[string]string) []link {
	links := make([]link, 0)
	insideCode := false
	for position := 0; position < len(line); position++ {
		if line[position] == '`' && !isEscaped(line, position) {
			insideCode = !insideCode
			continue
		}
		if insideCode || line[position] != '[' || isEscaped(line, position) {
			continue
		}

		labelEnd, ok := matchingBracket(line, position)
		if !ok {
			continue
		}
		label := line[position+1 : labelEnd]
		next := labelEnd + 1
		if next < len(line) && line[next] == '(' {
			destination, end, ok := inlineDestination(line, next)
			if ok {
				links = append(links, link{destination: destination, line: lineNumber})
				position = end
			}
			continue
		}

		reference := ""
		if next < len(line) && line[next] == '[' {
			referenceEnd, ok := matchingBracket(line, next)
			if !ok {
				continue
			}
			reference = line[next+1 : referenceEnd]
			if reference == "" {
				reference = label
			}
			position = referenceEnd
		} else {
			reference = label
		}
		if destination, exists := references[normalizeReference(reference)]; exists {
			links = append(links, link{destination: destination, line: lineNumber})
		}
	}
	return links
}

func matchingBracket(value string, start int) (int, bool) {
	depth := 0
	for index := start; index < len(value); index++ {
		if isEscaped(value, index) {
			continue
		}
		switch value[index] {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return index, true
			}
		}
	}
	return 0, false
}

func inlineDestination(value string, openingParenthesis int) (string, int, bool) {
	position := openingParenthesis + 1
	for position < len(value) && (value[position] == ' ' || value[position] == '\t') {
		position++
	}
	if position >= len(value) {
		return "", 0, false
	}
	if value[position] == '<' {
		end := strings.IndexByte(value[position+1:], '>')
		if end < 0 {
			return "", 0, false
		}
		end += position + 1
		closing := strings.IndexByte(value[end+1:], ')')
		if closing < 0 {
			return "", 0, false
		}
		return value[position+1 : end], end + closing + 1, true
	}

	start := position
	depth := 0
	for ; position < len(value); position++ {
		if isEscaped(value, position) {
			continue
		}
		switch value[position] {
		case '(':
			depth++
		case ')':
			if depth == 0 {
				return strings.TrimSpace(value[start:position]), position, true
			}
			depth--
		case ' ', '\t':
			if depth == 0 {
				destinationEnd := position
				for position < len(value) && value[position] != ')' {
					position++
				}
				if position < len(value) {
					return value[start:destinationEnd], position, true
				}
			}
		}
	}
	return "", 0, false
}

func checkLink(root string, sourcePath string, candidate link, trackedFiles map[string]struct{}) *Finding {
	if isExternal(candidate.destination) {
		return nil
	}

	parsed, err := url.Parse(candidate.destination)
	if err != nil {
		return finding(sourcePath, candidate, "has an invalid URL")
	}
	path, err := url.PathUnescape(parsed.EscapedPath())
	if err != nil {
		return finding(sourcePath, candidate, "has an invalid escaped path")
	}
	fragment, err := url.PathUnescape(parsed.EscapedFragment())
	if err != nil {
		return finding(sourcePath, candidate, "has an invalid escaped fragment")
	}

	targetPath := filepath.Join(root, filepath.FromSlash(sourcePath))
	if path == "" {
		targetPath = filepath.Join(root, filepath.FromSlash(sourcePath))
	} else if strings.HasPrefix(path, "/") {
		targetPath = filepath.Join(root, filepath.FromSlash(strings.TrimPrefix(path, "/")))
	} else {
		targetPath = filepath.Join(filepath.Dir(targetPath), filepath.FromSlash(path))
	}
	targetPath = filepath.Clean(targetPath)
	if !isWithin(root, targetPath) {
		return finding(sourcePath, candidate, "points outside this repository")
	}

	info, err := os.Stat(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return finding(sourcePath, candidate, "targets a file that does not exist")
		}
		return finding(sourcePath, candidate, "cannot inspect its target")
	}
	targetRelativePath, err := filepath.Rel(root, targetPath)
	if err != nil {
		return finding(sourcePath, candidate, "cannot determine its target path")
	}
	if !info.IsDir() && !isTrackedFile(filepath.ToSlash(targetRelativePath), trackedFiles) {
		return finding(sourcePath, candidate, "targets a file that is not tracked by Git")
	}
	if info.IsDir() && !isTrackedDirectory(filepath.ToSlash(targetRelativePath), trackedFiles) {
		return finding(sourcePath, candidate, "targets a directory with no tracked files")
	}
	resolvedRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return finding(sourcePath, candidate, "cannot resolve the repository root")
	}
	resolvedTarget, err := filepath.EvalSymlinks(targetPath)
	if err != nil {
		return finding(sourcePath, candidate, "cannot resolve its target")
	}
	if !isWithin(resolvedRoot, resolvedTarget) {
		return finding(sourcePath, candidate, "resolves outside this repository")
	}
	if fragment == "" {
		return nil
	}
	if info.IsDir() {
		targetPath = filepath.Join(targetPath, "README.md")
		info, err = os.Stat(targetPath)
		if err != nil {
			return finding(sourcePath, candidate, "targets a directory without README.md for its fragment")
		}
		targetRelativePath, err = filepath.Rel(root, targetPath)
		if err != nil || !isTrackedFile(filepath.ToSlash(targetRelativePath), trackedFiles) {
			return finding(sourcePath, candidate, "targets a directory without a tracked README.md for its fragment")
		}
		resolvedTarget, err = filepath.EvalSymlinks(targetPath)
		if err != nil {
			return finding(sourcePath, candidate, "cannot resolve its target")
		}
		if !isWithin(resolvedRoot, resolvedTarget) {
			return finding(sourcePath, candidate, "resolves outside this repository")
		}
	}
	if strings.ToLower(filepath.Ext(targetPath)) != markdownExtension || !info.Mode().IsRegular() {
		return finding(sourcePath, candidate, "uses a fragment on a non-Markdown target")
	}
	contents, err := os.ReadFile(targetPath)
	if err != nil {
		return finding(sourcePath, candidate, "cannot read its fragment target")
	}
	if !hasAnchor(string(contents), fragment) {
		return finding(sourcePath, candidate, "targets a heading that does not exist")
	}
	return nil
}

func isTrackedFile(path string, trackedFiles map[string]struct{}) bool {
	_, exists := trackedFiles[path]
	return exists
}

func isTrackedDirectory(path string, trackedFiles map[string]struct{}) bool {
	if path == "." {
		return len(trackedFiles) > 0
	}
	prefix := strings.TrimSuffix(path, "/") + "/"
	for trackedPath := range trackedFiles {
		if strings.HasPrefix(trackedPath, prefix) {
			return true
		}
	}
	return false
}

func isExternal(destination string) bool {
	parsed, err := url.Parse(destination)
	return err == nil && (parsed.Scheme != "" || parsed.Host != "" || strings.HasPrefix(destination, "//"))
}

func isWithin(root string, target string) bool {
	relativePath, err := filepath.Rel(root, target)
	return err == nil && relativePath != ".." && !strings.HasPrefix(relativePath, ".."+string(filepath.Separator))
}

func hasAnchor(contents string, fragment string) bool {
	anchors := make(map[string]struct{})
	counts := make(map[string]int)
	insideFence := false
	lines := strings.Split(contents, "\n")
	for index, line := range lines {
		if isFence(line) {
			insideFence = !insideFence
			continue
		}
		if insideFence {
			continue
		}
		heading, isHeading := atxHeading(line)
		if !isHeading && index+1 < len(lines) && isSetextUnderline(lines[index+1]) {
			heading = strings.TrimSpace(line)
			isHeading = heading != ""
		}
		if !isHeading {
			continue
		}
		anchor := githubAnchor(heading)
		if anchor == "" {
			continue
		}
		count := counts[anchor]
		counts[anchor]++
		if count > 0 {
			anchor = fmt.Sprintf("%s-%d", anchor, count)
		}
		anchors[anchor] = struct{}{}
	}
	_, exists := anchors[githubAnchor(fragment)]
	return exists
}

func atxHeading(line string) (string, bool) {
	trimmed := strings.TrimLeft(line, " ")
	headingLevel := 0
	for headingLevel < len(trimmed) && trimmed[headingLevel] == '#' {
		headingLevel++
	}
	if headingLevel == 0 || headingLevel > 6 || len(trimmed) == headingLevel || trimmed[headingLevel] != ' ' {
		return "", false
	}
	return strings.TrimSpace(strings.TrimRight(trimmed[headingLevel+1:], "# ")), true
}

func isSetextUnderline(line string) bool {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) == 0 {
		return false
	}
	underline := trimmed[0]
	if underline != '=' && underline != '-' {
		return false
	}
	for index := 1; index < len(trimmed); index++ {
		if trimmed[index] != underline {
			return false
		}
	}
	return true
}

func githubAnchor(value string) string {
	var builder strings.Builder
	lastWasDash := false
	for _, character := range strings.ToLower(strings.TrimSpace(value)) {
		switch {
		case unicode.IsLetter(character), unicode.IsNumber(character), character == '-', character == '_':
			builder.WriteRune(character)
			lastWasDash = false
		case unicode.IsSpace(character):
			if builder.Len() > 0 && !lastWasDash {
				builder.WriteByte('-')
				lastWasDash = true
			}
		}
	}
	return strings.Trim(builder.String(), "-")
}

func normalizeReference(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}

func isEscaped(value string, index int) bool {
	backslashes := 0
	for index > 0 && value[index-1] == '\\' {
		backslashes++
		index--
	}
	return backslashes%2 == 1
}

func finding(sourcePath string, candidate link, problem string) *Finding {
	return &Finding{Path: filepath.ToSlash(sourcePath), Line: candidate.line, Destination: candidate.destination, Problem: problem}
}
