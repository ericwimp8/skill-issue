#!/bin/sh

set -eu

repo_root=$(git rev-parse --show-toplevel)
git -C "$repo_root" config --local core.hooksPath .githooks

echo "Repository privacy hooks enabled."
