# Install Skill Issue

Skill Issue is distributed as a standalone command-line executable. It does
not require Go, Node.js, or another language runtime.

Download the current release from the
[Skill Issue releases page](https://github.com/ericwimp8/skill-issue/releases/latest).

## Prerequisites

Install and authenticate at least one supported agent harness before installing
the Skill Issue skills:

- OpenAI Codex
- Claude Code
- Cursor
- OpenCode
- Pi

## macOS

Check the machine architecture:

```sh
uname -m
```

Use the corresponding archive:

- `arm64`: `skill-issue_darwin_arm64.tar.gz`
- `x86_64`: `skill-issue_darwin_amd64.tar.gz`

For Apple Silicon, run:

```sh
mkdir -p "$HOME/.local/bin"
cd "$(mktemp -d)"

curl -LO https://github.com/ericwimp8/skill-issue/releases/latest/download/skill-issue_darwin_arm64.tar.gz
curl -LO https://github.com/ericwimp8/skill-issue/releases/latest/download/checksums.txt

grep skill-issue_darwin_arm64.tar.gz checksums.txt | shasum -a 256 -c -
tar -xzf skill-issue_darwin_arm64.tar.gz
chmod +x skill-issue
mv skill-issue "$HOME/.local/bin/skill-issue"
```

For an Intel Mac, replace both occurrences of
`skill-issue_darwin_arm64.tar.gz` with
`skill-issue_darwin_amd64.tar.gz`.

Add the installation directory to the shell path:

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.zshrc"
source "$HOME/.zshrc"
```

Confirm the installation:

```sh
skill-issue help
```

### macOS Security Prompt

macOS may initially block an executable downloaded from GitHub. After verifying
the published checksum:

1. Try running `skill-issue help`.
2. Open **System Settings → Privacy & Security**.
3. Find the message concerning `skill-issue`.
4. Select **Open Anyway**.
5. Run `skill-issue help` again.

## Linux

Check the machine architecture:

```sh
uname -m
```

Use the corresponding archive:

- `x86_64`: `skill-issue_linux_amd64.tar.gz`
- `aarch64` or `arm64`: `skill-issue_linux_arm64.tar.gz`

For x64 Linux, run:

```sh
mkdir -p "$HOME/.local/bin"
cd "$(mktemp -d)"

curl -LO https://github.com/ericwimp8/skill-issue/releases/latest/download/skill-issue_linux_amd64.tar.gz
curl -LO https://github.com/ericwimp8/skill-issue/releases/latest/download/checksums.txt

grep skill-issue_linux_amd64.tar.gz checksums.txt | sha256sum -c -
tar -xzf skill-issue_linux_amd64.tar.gz
chmod +x skill-issue
mv skill-issue "$HOME/.local/bin/skill-issue"
```

For ARM64 Linux, replace both occurrences of
`skill-issue_linux_amd64.tar.gz` with
`skill-issue_linux_arm64.tar.gz`.

Add the installation directory to the shell path if necessary:

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.profile"
export PATH="$HOME/.local/bin:$PATH"
```

Confirm the installation:

```sh
skill-issue help
```

## Windows

Download the appropriate archive:

- [Windows x64](https://github.com/ericwimp8/skill-issue/releases/latest/download/skill-issue_windows_amd64.zip)
- [Windows ARM64](https://github.com/ericwimp8/skill-issue/releases/latest/download/skill-issue_windows_arm64.zip)

Then:

1. Extract `skill-issue.exe` from the archive.
2. Create a directory such as
   `C:\Users\<your-name>\Apps\SkillIssue`.
3. Move `skill-issue.exe` into that directory.
4. Open **Edit environment variables for your account**.
5. Edit the user `Path` variable.
6. Add the Skill Issue directory.
7. Open a new PowerShell window.
8. Run `skill-issue help`.

Windows binaries are published, but native Windows runtime qualification is
still in progress.

## Install the Skills

For guided installation, run:

```sh
skill-issue install
```

The installer asks you to select:

1. The agent harness.
2. Project or user installation scope.
3. Whether to confirm the displayed installation.

The preview shows the selected harness, scope, native destination, and embedded
skills before anything is written.

### Project Scope

Project scope installs Skill Issue only in the current project:

```sh
cd /path/to/your/project
skill-issue install
```

Select **project** when prompted.

### User Scope

User scope installs the skills in the selected harness's user-level skill
directory, making them available across projects.

Select **user** when prompted.

### Argument-Driven Installation

The non-interactive form is:

```sh
skill-issue install \
  --workspace /path/to/project \
  --harness codex \
  --scope project
```

Supported harness identifiers are:

```text
claude-code
codex
cursor
opencode
pi
```

For example, install the Codex skills at user scope with:

```sh
skill-issue install \
  --workspace "$PWD" \
  --harness codex \
  --scope user
```

## Start Generating a Skill

After installation:

1. Open or restart the selected harness.
2. Open the project where you want to create a skill.
3. Explicitly request `skill-intake`.

For example:

```text
Use skill-intake. I want to create a skill that runs linting, applies automatic fixes, and resolves remaining compile-time errors.
```

You can also explicitly select `skill-intake` through the harness's native
skill-command interface when available. Skill intake is deliberately configured
for explicit invocation.

## Reinstall or Update

Download the latest executable and replace the existing `skill-issue` binary.
Then reinstall the embedded skills:

```sh
skill-issue install
```

Reinstallation replaces only the known Skill Issue skill directories.

## Uninstall

Use the same harness and scope selected during installation:

```sh
skill-issue uninstall \
  --workspace /path/to/project \
  --harness codex \
  --scope project
```

For user-scope removal, use `--scope user`.

See [`cli/README.md`](cli/README.md) for evaluation commands, harness
configuration, output formats, and recovery details.
