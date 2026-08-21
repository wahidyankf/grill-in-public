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

// Reference definitions accept CommonMark's optional indentation and either an
// angle-bracketed destination or an unbracketed non-whitespace destination.
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
	// Git defines the repository's document set. That avoids traversing ignored
	// installs and build artifacts while still seeing staged deletions.
	trackedFiles, err := findTrackedFiles(root)
	if err != nil {
		return nil, err
	}
	markdownFiles := filterMarkdownFiles(trackedFiles)

	findings := make([]Finding, 0)
	for _, sourcePath := range markdownFiles {
		// #nosec G304 -- sourcePath comes from Git's tracked paths inside root.
		contents, err := os.ReadFile(filepath.Join(root, sourcePath))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", sourcePath, err)
		}

		for _, candidate := range extractLinks(string(contents)) {
			finding := checkLink(root, sourcePath, candidate)
			if finding != nil {
				findings = append(findings, *finding)
			}
		}
	}

	// Stable ordering makes hook output reproducible and lets a reader repair
	// findings from top to bottom without nondeterministic map iteration.
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
	// NUL delimiters preserve unusual but valid filenames that contain spaces or
	// newlines; splitting ordinary line output would corrupt those paths.
	// #nosec G204 -- the executable and operation are fixed; root is only a Git working-directory argument.
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
		// Extension comparison is case-insensitive so the validator's scope is not
		// accidentally changed by a filename's casing.
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
	// Resolve definitions first because CommonMark permits a reference definition
	// to appear after the link that uses it.
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
	// A second pass records actual links with their source line after the complete
	// definition table is available.
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
	// Links in examples should remain literal teaching text, not validation input.
	trimmed := strings.TrimLeft(line, " ")
	return strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~")
}

func extractLineLinks(line string, lineNumber int, references map[string]string) []link {
	links := make([]link, 0)
	insideCode := false
	for position := 0; position < len(line); position++ {
		// Inline code can legitimately contain Markdown-looking characters. Track
		// the delimiter state instead of trying to remove code before parsing.
		if line[position] == '`' && !isEscaped(line, position) {
			insideCode = !insideCode
			continue
		}
		if insideCode || line[position] != '[' || isEscaped(line, position) {
			continue
		}

		candidate, end, ok := parseLinkAt(line, position, lineNumber, references)
		if ok {
			links = append(links, candidate)
			position = end
		}
	}
	return links
}

func parseLinkAt(line string, position, lineNumber int, references map[string]string) (link, int, bool) {
	labelEnd, ok := matchingBracket(line, position)
	if !ok {
		return link{}, position, false
	}

	label := line[position+1 : labelEnd]
	next := labelEnd + 1
	if next < len(line) && line[next] == '(' {
		return parseInlineLink(line, next, lineNumber)
	}

	return parseReferenceLink(line, label, labelEnd, next, lineNumber, references)
}

func parseInlineLink(line string, openingParenthesis, lineNumber int) (link, int, bool) {
	destination, end, ok := inlineDestination(line, openingParenthesis)
	if !ok {
		return link{}, openingParenthesis, false
	}

	return link{destination: destination, line: lineNumber}, end, true
}

// parseReferenceLink handles full, collapsed, and shortcut references against
// the complete definition map assembled during extractLinks' first pass.
func parseReferenceLink(
	line, label string,
	labelEnd, next, lineNumber int,
	references map[string]string,
) (link, int, bool) {
	reference := label
	end := labelEnd
	if next < len(line) && line[next] == '[' {
		referenceEnd, ok := matchingBracket(line, next)
		if !ok {
			return link{}, labelEnd, false
		}
		reference = line[next+1 : referenceEnd]
		if reference == "" {
			reference = label
		}
		end = referenceEnd
	}

	destination, exists := references[normalizeReference(reference)]
	if !exists {
		return link{}, end, false
	}

	return link{destination: destination, line: lineNumber}, end, true
}

func matchingBracket(value string, start int) (int, bool) {
	// Labels may contain nested brackets. A depth counter finds the matching
	// closer without mistaking an escaped bracket for syntax.
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
	position := skipInlineWhitespace(value, openingParenthesis+1)
	if position >= len(value) {
		return "", 0, false
	}
	if value[position] == '<' {
		return angleBracketDestination(value, position)
	}

	return unbracketedDestination(value, position)
}

func skipInlineWhitespace(value string, position int) int {
	for position < len(value) && (value[position] == ' ' || value[position] == '\t') {
		position++
	}

	return position
}

func angleBracketDestination(value string, position int) (string, int, bool) {
	// Angle-bracket destinations may contain spaces; their closing angle bracket
	// is the delimiter rather than the first whitespace character.
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

func unbracketedDestination(value string, start int) (string, int, bool) {
	depth := 0
	for position := start; position < len(value); position++ {
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
				// A title follows an unbracketed destination after whitespace. Return
				// only the destination but skip through the closing parenthesis.
				closing := strings.IndexByte(value[position:], ')')
				if closing >= 0 {
					return value[start:position], position + closing, true
				}
			}
		}
	}
	return "", 0, false
}

func checkLink(root, sourcePath string, candidate link) *Finding {
	// Network validation would make pre-push slow and nondeterministic; this
	// checker owns only repository-local references.
	if isExternal(candidate.destination) {
		return nil
	}

	path, fragment, problem := decodeLocalDestination(candidate.destination)
	if problem != "" {
		return finding(sourcePath, candidate, problem)
	}

	targetPath := localTargetPath(root, sourcePath, path)
	// Reject lexical traversal before touching the filesystem; this also produces
	// a clearer diagnostic than a later failed stat.
	if !isWithin(root, targetPath) {
		return finding(sourcePath, candidate, "points outside this repository")
	}

	info, resolvedRoot, problem := inspectLocalTarget(root, targetPath)
	if problem != "" {
		return finding(sourcePath, candidate, problem)
	}
	if fragment == "" {
		// Existing local files are sufficient when the link does not promise a
		// particular document location.
		return nil
	}

	problem = checkFragmentTarget(targetPath, fragment, info, resolvedRoot)
	if problem != "" {
		return finding(sourcePath, candidate, problem)
	}

	return nil
}

func decodeLocalDestination(destination string) (string, string, string) {
	parsed, err := url.Parse(destination)
	if err != nil {
		return "", "", "has an invalid URL"
	}
	// Parse URL structure before mapping to disk so #fragment and percent escapes
	// cannot be confused with literal filename characters.
	path, err := url.PathUnescape(parsed.EscapedPath())
	if err != nil {
		return "", "", "has an invalid escaped path"
	}
	fragment, err := url.PathUnescape(parsed.EscapedFragment())
	if err != nil {
		return "", "", "has an invalid escaped fragment"
	}

	return path, fragment, ""
}

func localTargetPath(root, sourcePath, path string) string {
	// Markdown paths begin at the source document unless a leading slash makes
	// them repository-relative.
	sourceFile := filepath.Join(root, filepath.FromSlash(sourcePath))
	if path == "" {
		return filepath.Clean(sourceFile)
	}
	if repositoryPath, ok := strings.CutPrefix(path, "/"); ok {
		return filepath.Clean(filepath.Join(root, filepath.FromSlash(repositoryPath)))
	}

	return filepath.Clean(filepath.Join(filepath.Dir(sourceFile), filepath.FromSlash(path)))
}

func inspectLocalTarget(root, targetPath string) (os.FileInfo, string, string) {
	info, err := os.Stat(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, "", "targets a file that does not exist"
		}

		return nil, "", "cannot inspect its target"
	}
	// Lexical containment is not enough: a symlink inside the repository could
	// resolve outside it, so validate both root and target after resolution.
	resolvedRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return nil, "", "cannot resolve the repository root"
	}
	resolvedTarget, err := filepath.EvalSymlinks(targetPath)
	if err != nil {
		return nil, "", "cannot resolve its target"
	}
	if !isWithin(resolvedRoot, resolvedTarget) {
		return nil, "", "resolves outside this repository"
	}

	return info, resolvedRoot, ""
}

func checkFragmentTarget(targetPath, fragment string, info os.FileInfo, resolvedRoot string) string {
	if info.IsDir() {
		var problem string
		targetPath, info, problem = directoryFragmentTarget(targetPath, resolvedRoot)
		if problem != "" {
			return problem
		}
	}
	if strings.ToLower(filepath.Ext(targetPath)) != markdownExtension || !info.Mode().IsRegular() {
		return "uses a fragment on a non-Markdown target"
	}
	// #nosec G304 -- targetPath passed lexical and resolved containment checks above.
	contents, err := os.ReadFile(targetPath)
	if err != nil {
		return "cannot read its fragment target"
	}
	if !hasAnchor(string(contents), fragment) {
		return "targets a heading that does not exist"
	}

	return ""
}

func directoryFragmentTarget(targetPath, resolvedRoot string) (string, os.FileInfo, string) {
	// Repository directory links render README.md, so fragments use that file.
	readmePath := filepath.Join(targetPath, "README.md")
	info, err := os.Stat(readmePath)
	if err != nil {
		return "", nil, "targets a directory without README.md for its fragment"
	}
	resolvedTarget, err := filepath.EvalSymlinks(readmePath)
	if err != nil {
		return "", nil, "cannot resolve its target"
	}
	if !isWithin(resolvedRoot, resolvedTarget) {
		return "", nil, "resolves outside this repository"
	}

	return readmePath, info, ""
}

func isExternal(destination string) bool {
	// Schemed and protocol-relative destinations are owned by external services,
	// including mailto links; only unqualified paths are repository-local.
	parsed, err := url.Parse(destination)
	return err == nil && (parsed.Scheme != "" || parsed.Host != "" || strings.HasPrefix(destination, "//"))
}

func isWithin(root, target string) bool {
	// filepath.Rel is platform-aware, unlike string-prefix checks that would
	// incorrectly accept siblings such as /repo-copy when root is /repo.
	relativePath, err := filepath.Rel(root, target)
	return err == nil && relativePath != ".." && !strings.HasPrefix(relativePath, ".."+string(filepath.Separator))
}

func hasAnchor(contents, fragment string) bool {
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
		// GitHub accepts both # ATX and underlined Setext headings, so mirror both
		// forms when constructing the target's anchor set.
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
		// GitHub disambiguates repeated headings with a zero-based suffix. Count the
		// base anchor before storing its first or subsequent generated form.
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
	// Markdown permits up to six leading # markers and requires whitespace after
	// them, which prevents ordinary text such as #hashtag from becoming a heading.
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
	// A Setext underline contains only one repeated marker: '=' for level one or
	// '-' for level two. Other punctuation remains ordinary paragraph text.
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
	// This intentionally models the subset of GitHub's slug behavior relevant to
	// repository headings: lowercase letters/numbers, preserved -/_, and spaces.
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
	// Reference labels are case-insensitive and collapse internal whitespace in
	// CommonMark, so definitions and uses compare on the same representation.
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}

func isEscaped(value string, index int) bool {
	// An odd run of backslashes escapes the next character; an even run represents
	// literal backslashes and leaves that character as Markdown syntax.
	backslashes := 0
	for index > 0 && value[index-1] == '\\' {
		backslashes++
		index--
	}
	return backslashes%2 == 1
}

func finding(sourcePath string, candidate link, problem string) *Finding {
	return &Finding{
		Path:        filepath.ToSlash(sourcePath),
		Line:        candidate.line,
		Destination: candidate.destination,
		Problem:     problem,
	}
}
