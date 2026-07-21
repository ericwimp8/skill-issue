#!/bin/sh

set -eu

repo_root=$(git rev-parse --show-toplevel)
cd "$repo_root"

denylist="$repo_root/.privacy-denylist.local"
commit_message=${1:-}
if [ ! -s "$denylist" ]; then
  echo "privacy check failed: create the required ignored local denylist" >&2
  exit 1
fi

violations=0

raw_codex_files=$(git grep --cached -l -E '"type"[[:space:]]*:[[:space:]]*"(session_meta|turn_context)"|"role"[[:space:]]*:[[:space:]]*"developer"|"encrypted_content"[[:space:]]*:' -- '*.jsonl' 2>/dev/null || true)

if [ -n "$raw_codex_files" ]; then
  violations=1
  echo "privacy check failed: raw Codex session context is tracked:" >&2
  printf '%s\n' "$raw_codex_files" >&2
fi

check_value() {
  value=$1
  [ -n "$value" ] || return

  if git grep --cached -a -i -F -q -- "$value" 2>/dev/null; then
    violations=1
  fi
  if git ls-files -z | tr '\0' '\n' | grep -i -F -q -- "$value"; then
    violations=1
  fi
  if {
    git var GIT_AUTHOR_IDENT
    git var GIT_COMMITTER_IDENT
  } | grep -a -i -F -q -- "$value"; then
    violations=1
  fi
  if [ -n "$commit_message" ] && grep -a -i -F -q -- "$value" "$commit_message"; then
    violations=1
  fi
}

while IFS= read -r value || [ -n "$value" ]; do
  check_value "$value"
done < "$denylist"

if [ "$violations" -ne 0 ]; then
  echo "privacy check failed: prospective commit violates repository privacy" >&2
  exit 1
fi

echo "privacy check passed"
