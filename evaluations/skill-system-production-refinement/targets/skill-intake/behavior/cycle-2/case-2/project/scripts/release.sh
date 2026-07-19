#!/bin/sh

mode="dry-run"
if [ "$1" = "--publish" ]; then
  mode="publish"
fi

printf 'release mode: %s\n' "$mode"
