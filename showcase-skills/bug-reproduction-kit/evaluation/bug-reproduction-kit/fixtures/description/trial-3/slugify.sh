#!/bin/sh

printf '%s' "$*" | tr '[:upper:] ' '[:lower:]-'
printf '\n'
