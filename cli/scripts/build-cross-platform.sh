#!/bin/sh
set -eu

script_directory=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cli_root=$(dirname "$script_directory")
output_directory="$cli_root/dist"
release_directory="$output_directory/release"

version=${VERSION:-dev}
commit=${COMMIT:-unknown}
build_date=${BUILD_DATE:-unknown}

rm -rf "$output_directory"
mkdir -p "$release_directory"

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

package_target() {
  target_os=$1
  target_arch=$2
  binary_name="skill-issue_${target_os}_${target_arch}"
  packaged_name="skill-issue"
  package_directory="$output_directory/package_${target_os}_${target_arch}"

  if [ "$target_os" = "windows" ]; then
    binary_name="${binary_name}.exe"
    packaged_name="${packaged_name}.exe"
  fi

  mkdir -p "$package_directory"
  cp "$output_directory/$binary_name" "$package_directory/$packaged_name"

  if [ "$target_os" = "windows" ]; then
    (cd "$package_directory" && zip -X -q "$release_directory/${binary_name%.exe}.zip" "$packaged_name")
  else
    tar --uid 0 --gid 0 --uname root --gname root -czf "$release_directory/${binary_name}.tar.gz" -C "$package_directory" "$packaged_name"
  fi

  rm -rf "$package_directory"
}

package_target darwin amd64
package_target darwin arm64
package_target windows amd64
package_target windows arm64
package_target linux amd64
package_target linux arm64

if command -v sha256sum >/dev/null 2>&1; then
  (cd "$release_directory" && sha256sum ./*.tar.gz ./*.zip > checksums.txt)
else
  (cd "$release_directory" && shasum -a 256 ./*.tar.gz ./*.zip > checksums.txt)
fi

echo "Built binaries in $output_directory"
echo "Packaged release assets in $release_directory"
