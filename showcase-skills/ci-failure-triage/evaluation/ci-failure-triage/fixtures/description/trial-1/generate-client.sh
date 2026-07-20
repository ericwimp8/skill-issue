#!/bin/sh
set -eu
version=$(sed -n 's/.*"schemaVersion":\([0-9][0-9]*\).*/\1/p' "$1")
if [ "$version" -gt 2 ]; then
  echo "generator: schema version $version unsupported; maximum is 2" >&2
  exit 64
fi
mkdir -p "$(dirname "$2")"
printf 'export const endpoint = "/status";\n' > "$2"
