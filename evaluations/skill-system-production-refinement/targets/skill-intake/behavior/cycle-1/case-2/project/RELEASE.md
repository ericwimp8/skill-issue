# Release Process

## Authority

`scripts/release.sh` is the only supported release entry point. Release operators must run it from the repository root.

## Safety

The script performs a dry-run unless the operator passes `--publish`. A release skill must preserve this default and must never infer publish authorization.

## Completion

A release is ready when validation passes, the proposed version is shown, and the operator explicitly decides whether to publish.
