#!/bin/sh
set -eu

script_directory=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cli_root=$(dirname "$script_directory")
output_directory="$cli_root/dist"

version=${VERSION:-dev}
commit=${COMMIT:-unknown}
build_date=${BUILD_DATE:-unknown}

rm -rf "$output_directory"
mkdir -p "$output_directory"

build_target() {
  target_os=$1
  target_arch=$2
  output_name="skill-issue_${target_os}_${target_arch}"
  if [ "$target_os" = "windows" ]; then
    output_name="${output_name}.exe"
  fi

  echo "Building ${target_os}/${target_arch}"
  CGO_ENABLED=0 GOOS="$target_os" GOARCH="$target_arch" go build \
    -trimpath \
    -ldflags "-s -w -X main.version=${version} -X main.commit=${commit} -X main.buildDate=${build_date}" \
    -o "$output_directory/$output_name" \
    ./cmd/skill-issue
}

cd "$cli_root"
build_target darwin amd64
build_target darwin arm64
build_target windows amd64
build_target windows arm64
build_target linux amd64
build_target linux arm64

echo "Built artifacts in $output_directory"
