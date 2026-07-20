#!/bin/sh
test "$(cat manifest.txt)" = "version=7.3.0
revision=ready-730" || exit 1
echo 'revision=ready-730 verification=passed'
