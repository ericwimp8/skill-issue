#!/bin/sh
set -eu
major=$(printf '%s' "$TOOL_VERSION" | cut -d. -f1)
if [ "$major" != "2" ]; then
  echo "unsupported tool major: $major" >&2
  exit 42
fi
printf 'tool accepted\n'
