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
  echo "privacy check failed: prospective commit contains a local denylist value" >&2
  exit 1
fi

echo "privacy check passed"
