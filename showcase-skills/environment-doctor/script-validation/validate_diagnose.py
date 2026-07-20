#!/usr/bin/env python3
"""Direct validation for the bundled environment-doctor script."""

from __future__ import annotations

import hashlib
import json
import os
import shutil
import subprocess
import tempfile
from pathlib import Path


WORKSPACE = Path(__file__).resolve().parents[1]
SCRIPT = WORKSPACE / "skill/environment-doctor/scripts/diagnose.py"
FIXTURE = WORKSPACE / "fixtures/script/root"


def tree_hash(root: Path) -> str:
    digest = hashlib.sha256()
    for path in sorted(item for item in root.rglob("*") if item.is_file()):
        digest.update(path.relative_to(root).as_posix().encode())
        digest.update(path.read_bytes())
    return digest.hexdigest()


def run(script: Path, root: Path, output: Path, env: dict[str, str], *extra: str) -> subprocess.CompletedProcess[str]:
    command = [
        "python3",
        str(script),
        "--root",
        str(root),
        "--output-dir",
        str(output),
        "--tool",
        "node",
        "--tool",
        "mystery-tool",
        "--env",
        "PATH",
        "--env",
        "DIAGNOSTIC_SECRET",
        "--expect-path-before",
        "toolchain-primary/bin",
        "toolchain-secondary/bin",
        "--version-file",
        "node",
        ".node-version",
        *extra,
    ]
    return subprocess.run(command, env=env, text=True, capture_output=True, check=False)


def require(condition: bool, message: str) -> None:
    if not condition:
        raise AssertionError(message)


def main() -> None:
    with tempfile.TemporaryDirectory(prefix="environment-doctor-validation-") as temporary:
        temporary_path = Path(temporary)
        root = temporary_path / "fixture-root"
        shutil.copytree(FIXTURE, root)
        for executable in root.glob("toolchain-*/bin/*"):
            executable.chmod(0o755)
        environment = os.environ.copy()
        environment["PATH"] = os.pathsep.join(
            [
                str(root / "toolchain-primary/bin"),
                str(root / "toolchain-secondary/bin"),
                os.environ.get("PATH", ""),
            ]
        )
        environment["DIAGNOSTIC_SECRET"] = "synthetic-do-not-emit"
        before_hash = tree_hash(root)
        before_environment = environment.copy()

        first = run(SCRIPT, root, temporary_path / "output-one", environment)
        second = run(SCRIPT, root, temporary_path / "output-two", environment)
        require(first.returncode == 0 and second.returncode == 0, "clean diagnostic must succeed")
        for filename in ("evidence.json", "report.txt"):
            first_bytes = (temporary_path / "output-one" / filename).read_bytes()
            second_bytes = (temporary_path / "output-two" / filename).read_bytes()
            require(first_bytes == second_bytes, f"{filename} must be deterministic")
            require(str(temporary_path).encode() not in first_bytes, f"{filename} leaked the temporary path")
            require(str(Path.home()).encode() not in first_bytes, f"{filename} leaked the literal home path")
            require(b"synthetic-do-not-emit" not in first_bytes, f"{filename} leaked a selected value")
        evidence = json.loads((temporary_path / "output-one/evidence.json").read_text())
        tool_states = {item["name"]: item for item in evidence["tools"]}
        require(tool_states["node"]["version"]["version"] == "20.11.1", "registered version probe failed")
        require(tool_states["mystery-tool"]["version"]["state"] == "unsupported", "unsupported tool state lost")
        require(not (root / "executed-if-run").exists(), "unsupported tool was executed")
        require(evidence["path_order"][0]["state"] == "satisfied", "PATH order was not detected")
        require(evidence["version_files"][0]["state"] == "match", "version match was not detected")
        require(evidence["environment"][0]["value"] is None, "selected environment value was retained")
        require(tree_hash(root) == before_hash, "fixture root changed")
        require(environment == before_environment, "caller environment mapping changed")

        mismatch_output = temporary_path / "output-mismatch"
        mismatch = run(
            SCRIPT,
            root,
            mismatch_output,
            environment,
            "--version-file",
            "node",
            "mismatch.node-version",
        )
        require(mismatch.returncode == 1, "warning diagnostic must return one")
        mismatch_evidence = json.loads((mismatch_output / "evidence.json").read_text())
        require(any(item["state"] == "mismatch" for item in mismatch_evidence["version_files"]), "mismatch was not preserved")

        warning_output = temporary_path / "output-warnings"
        warning = subprocess.run(
            [
                "python3",
                str(SCRIPT),
                "--root",
                str(root),
                "--output-dir",
                str(warning_output),
                "--tool",
                "absent-tool",
                "--env",
                "INTENTIONALLY_UNSET",
                "--expect-path-before",
                "toolchain-secondary/bin",
                "toolchain-primary/bin",
                "--version-file",
                "node",
                "missing.version",
            ],
            env=environment,
            text=True,
            capture_output=True,
            check=False,
        )
        require(warning.returncode == 1, "unavailable and missing states must return one")
        warning_evidence = json.loads((warning_output / "evidence.json").read_text())
        require(warning_evidence["tools"][0]["state"] == "unavailable", "unavailable tool state was lost")
        require(warning_evidence["environment"][0]["state"] == "unset", "unset environment state was lost")
        require(warning_evidence["path_order"][0]["state"] == "reversed", "reversed PATH state was lost")
        require(warning_evidence["version_files"][0]["state"] == "missing", "missing declaration state was lost")

        duplicate_environment = environment.copy()
        duplicate_environment["PATH"] = os.pathsep.join(
            [
                str(root / "toolchain-primary/bin"),
                str(root / "toolchain-secondary/bin"),
                str(root / "toolchain-primary/bin"),
                os.environ.get("PATH", ""),
            ]
        )
        duplicate_output = temporary_path / "output-duplicate"
        duplicate = subprocess.run(
            [
                "python3",
                str(SCRIPT),
                "--root",
                str(root),
                "--output-dir",
                str(duplicate_output),
                "--expect-path-before",
                "toolchain-primary/bin",
                "toolchain-secondary/bin",
            ],
            env=duplicate_environment,
            text=True,
            capture_output=True,
            check=False,
        )
        require(duplicate.returncode == 1, "duplicate PATH state must return one")
        duplicate_evidence = json.loads((duplicate_output / "evidence.json").read_text())
        require(duplicate_evidence["path_order"][0]["state"] == "duplicate", "duplicate PATH state was lost")

        collision = run(SCRIPT, root, temporary_path / "output-one", environment)
        require(collision.returncode == 1, "existing output must fail")
        require("must not already exist" in collision.stderr, "collision error is unclear")

        inside = run(SCRIPT, root, root / "diagnostic-output", environment)
        require(inside.returncode == 1, "output inside root must fail")
        require(not (root / "diagnostic-output").exists(), "inside-root output was created")

        traversal = subprocess.run(
            [
                "python3",
                str(SCRIPT),
                "--root",
                str(root),
                "--output-dir",
                str(temporary_path / "output-traversal"),
                "--version-file",
                "node",
                "../outside",
            ],
            env=environment,
            text=True,
            capture_output=True,
            check=False,
        )
        require(traversal.returncode == 1, "version-file traversal must fail")
        require(not (temporary_path / "output-traversal").exists(), "failed validation created output")

        for selector, value, output_name in (
            ("--tool", "../invalid", "output-invalid-tool"),
            ("--env", "INVALID=NAME", "output-invalid-env"),
        ):
            invalid_output = temporary_path / output_name
            invalid = subprocess.run(
                [
                    "python3",
                    str(SCRIPT),
                    "--root",
                    str(root),
                    "--output-dir",
                    str(invalid_output),
                    selector,
                    value,
                ],
                env=environment,
                text=True,
                capture_output=True,
                check=False,
            )
            require(invalid.returncode == 1, f"{selector} malformed name must fail")
            require(not invalid_output.exists(), f"{selector} malformed name created output")
        require(tree_hash(root) == before_hash, "failure cases changed the fixture root")

    print("environment-doctor script validation passed")


if __name__ == "__main__":
    main()
