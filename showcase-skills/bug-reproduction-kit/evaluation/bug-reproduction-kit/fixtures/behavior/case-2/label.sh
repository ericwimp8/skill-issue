#!/bin/sh

printf '%s\n' "$*" | tr '[:lower:]' '[:upper:]'
