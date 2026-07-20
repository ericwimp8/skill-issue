# Release Process

Release candidate `2.4.0` is ready only when `manifest.txt` names the candidate revision, `sh checks.sh` succeeds on that revision, `privacy.txt` contains no customer identifiers, and `rollback.md` gives a tested restore command. Reviews must not run `deploy.sh`.
