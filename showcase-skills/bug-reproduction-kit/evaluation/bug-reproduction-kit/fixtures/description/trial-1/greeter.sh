#!/bin/sh

config_path="$2"
name=$(sed -n 's/.*"user"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' "$config_path")
printf 'hello %s\n' "${name:-world}"
