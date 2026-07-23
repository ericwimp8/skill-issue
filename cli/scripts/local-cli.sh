#!/bin/sh
set -eu

script_directory=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cli_root=$(dirname "$script_directory")
repository_root=$(dirname "$cli_root")
binary_root="$repository_root/.skill-issue/bin"
development_binary="$binary_root/development/skill-issue"
known_good_root="$binary_root/known-good"
known_good_current="$known_good_root/current"
plugin_dependency="dependencies/codex-skill-issue-plugin"

archive_root=
temporary_binary=

cleanup() {
  if [ -n "$temporary_binary" ]; then
    rm -f "$temporary_binary"
  fi
  if [ -n "$archive_root" ]; then
    rm -rf "$archive_root"
  fi
}

trap cleanup 0 1 2 15

build_binary() {
  source_root=$1
  output_path=$2
  version=$3
  commit=$4
  build_date=$5

  mkdir -p "$(dirname "$output_path")"
  temporary_binary="${output_path}.tmp.$$"
  (
    cd "$source_root"
    CGO_ENABLED=0 go build \
      -trimpath \
      -ldflags "-s -w -X main.version=${version} -X main.commit=${commit} -X main.buildDate=${build_date}" \
      -o "$temporary_binary" \
      ./cli/cmd/skill-issue
  )
  mv "$temporary_binary" "$output_path"
  temporary_binary=
}

build_development() {
  commit=$(git -C "$repository_root" rev-parse HEAD)
  short_commit=$(git -C "$repository_root" rev-parse --short=12 HEAD)
  build_date=$(date -u '+%Y-%m-%dT%H:%M:%SZ')
  state=clean
  if [ -n "$(git -C "$repository_root" status --porcelain -- cli bundle.go go.mod "$plugin_dependency" supporting-skills evaluations/skill-calling/built-ins ':(glob)evaluations/scenario-skill-refinement/*/skill/**')" ]; then
    state=dirty
  fi

  build_binary \
    "$repository_root" \
    "$development_binary" \
    "development-${short_commit}-${state}" \
    "$commit" \
    "$build_date"
  echo "Built development CLI at $development_binary"
}

build_known_good() {
  commit=$(git -C "$repository_root" rev-parse HEAD)
  short_commit=$(git -C "$repository_root" rev-parse --short=12 HEAD)
  output_path="$known_good_root/$commit/skill-issue"

  if [ ! -x "$output_path" ]; then
    archive_root=$(mktemp -d "${TMPDIR:-/tmp}/skill-issue-known-good.XXXXXX")
    git -C "$repository_root" archive HEAD | tar -xf - -C "$archive_root"
    dependency_commit=$(git -C "$repository_root" rev-parse "HEAD:$plugin_dependency")
    dependency_checkout="$repository_root/$plugin_dependency"
    if ! git -C "$dependency_checkout" cat-file -e "${dependency_commit}^{commit}" 2>/dev/null; then
      echo "Plugin dependency is unavailable. Run: git submodule update --init --recursive" >&2
      exit 1
    fi
    mkdir -p "$archive_root/$plugin_dependency"
    git -C "$dependency_checkout" archive "$dependency_commit" |
      tar -xf - -C "$archive_root/$plugin_dependency"
    build_binary \
      "$archive_root" \
      "$output_path" \
      "known-good-${short_commit}" \
      "$commit" \
      "$(date -u '+%Y-%m-%dT%H:%M:%SZ')"
    rm -rf "$archive_root"
    archive_root=
  fi

  mkdir -p "$known_good_root"
  printf '%s\n' "$commit" > "${known_good_current}.tmp.$$"
  mv "${known_good_current}.tmp.$$" "$known_good_current"
  echo "Selected known-good CLI at $output_path"
}

resolve_known_good() {
  if [ ! -f "$known_good_current" ]; then
    echo "Known-good CLI is not configured. Run: $0 build-known-good" >&2
    exit 1
  fi
  IFS= read -r selected_commit < "$known_good_current"
  selected_binary="$known_good_root/$selected_commit/skill-issue"
  if [ ! -x "$selected_binary" ]; then
    echo "Known-good CLI is missing. Run: $0 build-known-good" >&2
    exit 1
  fi
  printf '%s\n' "$selected_binary"
}

run_known_good() {
  selected_binary=$(resolve_known_good)
  exec "$selected_binary" "$@"
}

run_development() {
  if [ ! -x "$development_binary" ]; then
    echo "Development CLI is missing. Run: $0 build-development" >&2
    exit 1
  fi
  exec "$development_binary" "$@"
}

show_paths() {
  echo "development: $development_binary"
  if [ -f "$known_good_current" ]; then
    echo "known-good: $(resolve_known_good)"
  else
    echo "known-good: not configured"
  fi
}

command=${1:-}
case "$command" in
  build-development)
    build_development
    ;;
  build-known-good)
    build_known_good
    ;;
  development)
    shift
    run_development "$@"
    ;;
  known-good)
    shift
    run_known_good "$@"
    ;;
  paths)
    show_paths
    ;;
  *)
    run_known_good "$@"
    ;;
esac
