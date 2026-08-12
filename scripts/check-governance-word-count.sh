#!/usr/bin/env bash

# Enforce concise agent guidance. Documents over this limit should be split into
# focused files and loaded only when relevant (progressive disclosure).
set -euo pipefail

# `wc -w` defines a word as a whitespace-delimited token.
MAX_WORDS=500

# Root guidance always applies, while repo-governance contains the detailed,
# task-specific documents that extend it.
AGENTS_FILE="AGENTS.md"
GOVERNANCE_DIRECTORY="repo-governance"

# Fail fast with the document path and an actionable remediation when a single
# guidance file exceeds the repository's word budget.
check_file() {
  local file="$1"
  local word_count

  word_count=$(wc -w < "$file" | tr -d '[:space:]')

  if ((word_count > MAX_WORDS)); then
    printf 'ERROR: %s contains %s words; the limit is %s.\n' \
      "$file" "$word_count" "$MAX_WORDS" >&2
    printf 'Use progressive disclosure: split detailed guidance into focused files.\n' >&2
    return 1
  fi
}

# These paths are part of the repository's governance layout. A missing path is
# an error instead of a skipped check, so a rename cannot silently disable it.
if [[ ! -f "$AGENTS_FILE" ]]; then
  printf 'ERROR: Required file not found: %s\n' "$AGENTS_FILE" >&2
  exit 1
fi

if [[ ! -d "$GOVERNANCE_DIRECTORY" ]]; then
  printf 'ERROR: Required directory not found: %s\n' "$GOVERNANCE_DIRECTORY" >&2
  exit 1
fi

# Check root guidance first because it is the common entry point for agents.
check_file "$AGENTS_FILE"

# `-print0` and `read -d ''` preserve paths containing spaces while checking
# every Markdown governance document below this directory, including nested
# documents. Other file types do not contain agent guidance and are excluded.
while IFS= read -r -d '' file; do
  check_file "$file"
done < <(find "$GOVERNANCE_DIRECTORY" -type f -name '*.md' -print0)

printf 'Governance word counts are within the %s-word limit.\n' "$MAX_WORDS"
