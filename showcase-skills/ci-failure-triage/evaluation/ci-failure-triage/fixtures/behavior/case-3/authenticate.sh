#!/bin/sh
set -eu
if [ -z "${PREVIEW_TOKEN:-}" ]; then
  echo "authentication failed" >&2
  exit 23
fi
if [ "$PREVIEW_TOKEN" = "revoked-fixture-token" ]; then
  echo "authentication failed" >&2
  exit 23
fi
printf 'authenticated\n'
