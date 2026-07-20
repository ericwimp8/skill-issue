# Cross-Platform CLI Distribution for Skill Issue

## Executive recommendation

Implement the thin Skill Issue CLI in **Go** and publish prebuilt, signed executables through **GitHub Releases**, using **GoReleaser** to produce the platform matrix, archives, checksums, and package-manager metadata.

This is the simplest route to the desired experience:

- users download one executable for their operating system and architecture;
- Node.js, Python, Dart, .NET, and Go are absent from the user's prerequisites;
- the same small codebase builds for macOS, Windows, and Linux;
- the initial release pipeline can remain compact; and
- mature release tooling can add Homebrew, WinGet, Scoop, Linux packages, checksums, and signatures without changing the CLI implementation.

Go is especially suitable because Skill Issue's proposed standalone component is a thin systems-integration tool: filesystem work, subprocess orchestration, archive/download handling, configuration, and report generation. Go produces an ordinary executable with `go build`, and its compiler supports macOS, Windows, and Linux across both x86-64 and Arm64 targets. A pure-Go program can usually build the full target matrix from one host by setting `GOOS` and `GOARCH`; the main caveat is avoiding `cgo` or accepting native-runner builds for dependencies that require it. The official Go target list includes `darwin`, `linux`, and `windows` on both `amd64` and `arm64`. [Go build tutorial](https://go.dev/doc/tutorial/compile-install) [Go supported targets](https://go.dev/doc/install/source)

The best conditional alternative is **Deno compile** if a TypeScript implementation already exists or TypeScript development speed materially outweighs binary size and embedded-runtime overhead. Deno produces standalone executables and explicitly supports cross-compiling from any host to Windows x64, macOS x64/Arm64, and Linux x64/Arm64. [Deno compile documentation](https://docs.deno.com/runtime/reference/cli/compile/)

## Product constraints

The decision is driven by these requirements:

1. macOS users should be able to download and run Skill Issue without installing a language runtime;
2. Windows should offer an equivalently simple executable or package-manager install;
3. Linux support should add little implementation burden;
4. the CLI should remain easy to release, sign, update, and troubleshoot;
5. the implementation should optimize for a thin orchestration tool rather than application-framework capability; and
6. the first release should avoid installer engineering that adds little value over an archive plus executable.

## Approach comparison

| Approach                     | User-facing output                                            | Cross-platform build story                                                                                                | Release and security friction                                                                       | Maintainability for Skill Issue                                                                         | Fit                                                  |
| ---------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ---------------------------------------------------- |
| **Go native CLI**            | Small ordinary executable per OS/architecture                 | Excellent for pure Go through `GOOS`/`GOARCH`; native runners remain useful for signing and final testing                 | Straightforward with GoReleaser; normal macOS and Windows signing still applies                     | Low operational complexity; explicit, stable standard library and tooling                               | **Best overall**                                     |
| **Deno compiled TypeScript** | One executable embedding the stripped `denort` runtime        | Explicit cross-compilation for Windows x64, macOS x64/Arm64, Linux x64/Arm64                                              | Standard platform signing; larger runtime-bearing artifacts                                         | Excellent if the codebase is already TypeScript; runtime/tool compatibility must be tested              | **Best TypeScript alternative**                      |
| **Bun compiled TypeScript**  | One executable containing the Bun runtime, code, and packages | Explicit targets for Windows, macOS, glibc Linux, and musl Linux; x64 baseline/modern variants add a choice               | Standard platform signing; embedded runtime and CPU-target variants enlarge release/testing surface | Fast JS/TS development, but more runtime-specific behavior to own                                       | Good when Bun APIs or npm compatibility are decisive |
| **Rust native CLI**          | Small native executable per target                            | Broad target support, but linkers, native dependencies, and Windows/macOS target details make cross-building less uniform | `dist`/cargo-dist provides strong generated CI, archives, and installer scripts                     | Excellent runtime properties; higher implementation and compile-time complexity for a thin orchestrator | Strong, but more engineering than needed             |
| **Dart AOT executable**      | Self-contained native executable with a small Dart runtime    | macOS/Windows/Linux executables; cross-compilation currently targets Linux only                                           | Standard platform signing; macOS and Windows require platform-specific build jobs                   | Attractive if the maintainers are substantially more productive in Dart                                 | Reasonable team-expertise alternative                |
| **.NET Native AOT**          | Self-contained native single file                             | Targets Windows, Linux, and macOS, but AOT publishing is runtime-identifier-specific                                      | Standard signing plus AOT/trimming compatibility checks                                             | Rich ecosystem, but trimming, reflection, and AOT restrictions are unnecessary risk here                | Lower fit                                            |
| **Node.js SEA**              | Node binary with the application injected                     | Intended to remove user-side Node, but Node labels the feature active development; current CI coverage excludes macOS x64 | Binary injection and re-signing are part of the workflow                                            | Retains JS ecosystem but adds a specialized release process                                             | Lower fit for a new tool                             |
| **PyInstaller**              | One file or directory containing Python and dependencies      | Must build separately on each target OS; one-file mode extracts to a temporary directory at runtime                       | Standard signing plus antivirus/bootloader troubleshooting surface                                  | Useful only if substantial Python code already exists                                                   | Lower fit                                            |

### Why Go wins

Go provides the best combined result rather than the absolute best result in only one dimension:

- **Build output:** an ordinary executable rather than a script plus a separately installed runtime. Go's official build workflow directly produces an executable. [Go build tutorial](https://go.dev/doc/tutorial/compile-install)
- **Target coverage:** the compiler supports the required operating-system and architecture combinations. [Go supported targets](https://go.dev/doc/install/source)
- **Release automation:** GoReleaser generates a `GOOS`/`GOARCH` matrix and can create archives, direct binary uploads, and SHA-256 checksum files. [GoReleaser Go builds](https://goreleaser.com/customization/builds/builders/go/) [GoReleaser archives](https://goreleaser.com/customization/package/archives/) [GoReleaser checksums](https://goreleaser.com/customization/package/checksum/)
- **Installer evolution:** the same release configuration can later publish Homebrew casks, WinGet manifests, Scoop manifests, and `.deb`/`.rpm`/`.apk` packages. [GoReleaser Homebrew](https://goreleaser.com/customization/publish/homebrew_casks/) [GoReleaser WinGet](https://goreleaser.com/customization/publish/winget/) [GoReleaser Scoop](https://goreleaser.com/customization/scoop/) [GoReleaser nFPM](https://goreleaser.com/customization/package/nfpm/)
- **Operational simplicity:** the proposed CLI does not need a browser UI, embedded server runtime, reflection-heavy plugin system, or other feature that would justify shipping a larger managed runtime.

The key implementation constraint is to keep the CLI pure Go where practical. Calling external coding-agent harnesses through subprocesses is compatible with that constraint. Native libraries introduced through `cgo` would weaken the one-host cross-build advantage and should be accepted only when they solve a concrete product need.

## Recommended release artifact shape

Start with six primary artifacts and a checksum file:

```text
skill-issue_<version>_darwin_arm64.tar.gz
skill-issue_<version>_darwin_amd64.tar.gz
skill-issue_<version>_windows_amd64.zip
skill-issue_<version>_windows_arm64.zip
skill-issue_<version>_linux_amd64.tar.gz
skill-issue_<version>_linux_arm64.tar.gz
skill-issue_<version>_checksums.txt
```

Each archive should contain:

```text
skill-issue[.exe]
README.md
LICENSE
```

Use `tar.gz` on macOS and Linux so executable permissions survive extraction naturally, and `.zip` on Windows. GoReleaser supports both archived binaries and direct binary uploads; an archive is preferable because it carries the license/readme and gives releases consistent filenames. [GoReleaser archives](https://goreleaser.com/customization/package/archives/)

Do not create a macOS `.app`, `.dmg`, or graphical installer for the MVP. They imply a graphical application that the product does not have. A signed and notarized CLI inside a small archive is the native-feeling shape for this product. GoReleaser can create a universal macOS binary, but separate Arm64 and x86-64 archives keep each download smaller and allow an installer or website to select the correct one automatically. [GoReleaser universal macOS binaries](https://goreleaser.com/customization/builds/universalbinaries/)

GitHub Releases is an appropriate first host because releases attach versioned binary files and release notes to Git tags, and GitHub supports stable `releases/latest/download/<asset>` URLs for manually uploaded assets. [GitHub releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases) [GitHub latest-release links](https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases)

## Installation experience

Offer installation in this order:

1. **Direct download:** the website detects OS/architecture and links to the matching GitHub Release archive. Users extract it and run `skill-issue` or `skill-issue.exe`.
2. **macOS package-manager path:** `brew install <tap>/skill-issue`. Homebrew casks are specifically the package definition for precompiled upstream binaries, and GoReleaser can update a tap automatically. [Homebrew formula/cask model](https://docs.brew.sh/Formula-Cookbook) [GoReleaser Homebrew](https://goreleaser.com/customization/publish/homebrew_casks/)
3. **Windows package-manager path:** `winget install <publisher>.SkillIssue`. WinGet supports ZIP and portable packages and provides install and upgrade commands on modern Windows. [WinGet formats and commands](https://learn.microsoft.com/en-us/windows/package-manager/winget/) [WinGet install](https://learn.microsoft.com/en-us/windows/package-manager/winget/install)
4. **Linux convenience path:** direct archive first; add a POSIX shell installer and `.deb`/`.rpm` packages only when usage justifies repository/package maintenance.

A one-line shell installer and PowerShell installer can improve the first-run experience by detecting the platform, downloading the correct archive, verifying its SHA-256 checksum, and installing into a user-writable directory already on `PATH`. They should remain convenience layers over the same release assets, not separate product implementations.

## Signing, notarization, and security friction

### macOS

A raw unsigned CLI downloaded from a browser can trigger Gatekeeper friction. Apple recommends signing outside-App-Store software with a Developer ID certificate and submitting it for notarization; `notarytool` uploads custom build artifacts and `stapler` attaches the result. Apple explicitly supports ZIP, PKG, and DMG submissions. [Apple Developer ID](https://developer.apple.com/developer-id/) [Apple notarization](https://developer.apple.com/documentation/security/notarizing-macos-software-before-distribution)

For a smooth public download:

1. build the macOS binaries;
2. sign each executable with a Developer ID Application certificate, hardened runtime, and secure timestamp;
3. archive the signed executable;
4. submit the archive with `xcrun notarytool`;
5. staple where the artifact format supports it and verify with `spctl`/`codesign`; and
6. test a fresh browser download on supported macOS versions.

Developer ID distribution requires Apple Developer Program membership, currently 99 USD per membership year. [Apple program enrollment](https://developer.apple.com/programs/enroll/)

Signing and notarization should be treated as part of the MVP's public distribution work rather than postponed indefinitely. Asking users to remove the quarantine attribute creates a visibly non-native and trust-eroding first run.

### Windows

An unsigned `.exe` can trigger Microsoft Defender SmartScreen's “Windows protected your PC” flow. Microsoft states that unsigned files start reputation from zero on every version, while consistently signing releases allows publisher reputation to accumulate. Even a newly signed binary can initially show a warning. Microsoft currently recommends Artifact Signing for non-Store distribution; it integrates with CI and requires identity validation. [Microsoft SmartScreen reputation](https://learn.microsoft.com/en-us/windows/apps/package-and-deploy/smartscreen-reputation)

For early private testing, unsigned Windows artifacts may be acceptable with clear checksums. Before broad public promotion, sign every `.exe` with one consistent publisher identity. WinGet improves install and update ergonomics but does not remove the need to consider the trust of the underlying artifact.

### Linux and supply-chain metadata

Linux has no equivalent single mandatory notarization gate for a portable CLI. Publish SHA-256 checksums for all release artifacts and sign either each artifact or the checksum manifest. GoReleaser supports executable/archive signing and notes that signing the checksum file is generally sufficient for verification workflows. [GoReleaser signing](https://goreleaser.com/customization/sign/sign/)

GitHub artifact attestations can be added later to make CI provenance independently verifiable, but checksums and platform signatures solve the immediate user experience first.

## CI and cross-compilation implications

Use tag-driven GitHub Actions with this shape:

1. run unit/integration tests on `ubuntu-latest`, `macos-latest`, and `windows-latest`;
2. build the unsigned pure-Go matrix with GoReleaser;
3. sign and notarize macOS artifacts in a macOS job with protected secrets;
4. sign Windows executables in a Windows or Artifact Signing job;
5. assemble archives and generate checksums only after signing so hashes describe final bytes;
6. smoke-test final artifacts on their native operating systems; and
7. publish one draft GitHub Release atomically after all artifacts pass.

GitHub Actions supports an OS matrix including macOS, Windows, and Ubuntu. Standard hosted runners currently cover x64 Windows/Ubuntu, x64 Intel macOS, Arm64 macOS, and public-preview Arm64 Windows/Ubuntu labels, so native smoke coverage can expand without self-hosted machines. [GitHub Actions matrix syntax](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax) [GitHub-hosted runner reference](https://docs.github.com/en/actions/reference/runners/github-hosted-runners)

Pure Go means compilation itself can be centralized, but signing and native smoke tests still justify platform jobs. This hybrid keeps the build definition simple without confusing “cross-compiled” with “validated on the target OS.”

## Update strategy

Use **package-manager-owned updates first**:

- Homebrew users run `brew upgrade skill-issue`;
- WinGet users run `winget upgrade <publisher>.SkillIssue`;
- direct-download users receive a lightweight version check and a link or exact update command.

For the MVP, implement `skill-issue version` and optionally a non-blocking “new version available” check with a documented opt-out. Avoid an automatic in-place self-updater initially because it must correctly handle signatures, checksum verification, permissions, locked Windows executables, rollback, package-manager installations, and interrupted writes.

If direct-download usage becomes material, add an explicit `skill-issue self-update` later. It should download the platform archive from a versioned release, verify the published checksum/signature, replace atomically, and refuse to overwrite package-manager-managed installations. Never update silently.

## Conditional alternatives

### Choose Deno compile when TypeScript is already the product language

Deno is the strongest alternative if Skill Issue accumulates substantial TypeScript before the CLI implementation begins. `deno compile` embeds the program in a stripped `denort` runtime, includes additional files/directories when requested, and cross-compiles to all supported targets regardless of host platform. [Deno compile documentation](https://docs.deno.com/runtime/reference/cli/compile/)

Trade-offs against Go:

- larger artifacts because the runtime is embedded;
- more testing around npm dependencies, dynamic loading, subprocess behavior, and embedded assets;
- a runtime-bearing executable rather than a small native systems CLI; and
- less benefit from the language choice if most CLI behavior is filesystem and subprocess orchestration.

Choose Deno when reusing TypeScript business logic or the team's TypeScript velocity clearly offsets those costs.

### Choose Bun compile when Bun compatibility is specifically valuable

Bun's `--compile` produces a single executable containing the application, packages, and Bun runtime. It supports macOS and Windows x64/Arm64 plus glibc and musl Linux targets. Its x64 targets distinguish baseline and modern CPU variants, which adds a compatibility choice to the artifact matrix. [Bun standalone executables](https://bun.sh/docs/bundler/executables)

Choose Bun over Deno only when Bun or Node API compatibility has been proven against Skill Issue's real dependencies. Otherwise Deno's documented compile surface and target matrix are a clearer TypeScript distribution contract.

### Choose Rust when the CLI becomes a substantial native systems product

Rust provides excellent native binaries, strong correctness properties, and broad target support. `dist` can generate GitHub release CI, per-platform builds, archives, shell/PowerShell installers, and package publishing. [Rust Cargo target builds](https://doc.rust-lang.org/cargo/commands/cargo-build.html) [dist release workflow](https://axodotdev.github.io/cargo-dist/) [dist installation outputs](https://axodotdev.github.io/cargo-dist/book/install.html)

Choose Rust if Skill Issue later requires a long-lived daemon, security-sensitive parsers, heavy concurrency, or a deep Rust library ecosystem. For the currently proposed thin orchestrator, the additional language and linker complexity does not buy enough user-facing value over Go.

### Choose Dart when team expertise dominates release simplicity

`dart compile exe` creates self-contained executables for Windows, macOS, and Linux and supports platform code signing. Dart cross-compilation currently targets Linux only, so macOS and Windows releases require their own build jobs. [Dart compile documentation](https://dart.dev/tools/dart-compile)

This is viable if maintainers are already substantially more effective in Dart and want to share pure-Dart libraries. It is less simple operationally than pure Go for the full release matrix and offers little direct advantage to a non-Flutter CLI user.

## Lower-fit approaches

- **Node.js SEA:** Node can now build a single executable, but the official feature remains “active development,” requires platform signing, and its regular CI coverage currently excludes macOS x64. This is an avoidable release risk for a new cross-platform tool. [Node.js SEA documentation](https://nodejs.org/api/single-executable-applications.html)
- **PyInstaller:** it successfully removes the user-side Python prerequisite, but it is not a cross-compiler; every target OS must build separately. One-file mode extracts an embedded runtime to a temporary directory on startup and can fail where temporary storage is mounted `noexec`. [PyInstaller operating model](https://www.pyinstaller.org/en/stable/operating-mode.html) [PyInstaller manual](https://pyinstaller.org/en/stable/index.html)
- **.NET Native AOT:** it creates a self-contained native single file without a user-installed .NET runtime, but trimming, reflection, dynamic loading, and library compatibility limitations add constraints that Skill Issue does not need. [Microsoft Native AOT](https://learn.microsoft.com/en-us/dotnet/core/deploying/native-aot/)

## Practical first-release scope

1. Implement a pure-Go CLI with embedded static assets and no `cgo` dependencies.
2. Support macOS Arm64/x86-64, Windows x86-64, and Linux x86-64 first; add Windows Arm64 and Linux Arm64 in the same pipeline if native smoke tests are reliable.
3. Configure GoReleaser for version injection, reproducible flags, archives, and SHA-256 checksums.
4. Publish draft GitHub Releases from semantic version tags.
5. Sign/notarize both macOS binaries and sign Windows executables before public launch.
6. Publish a Homebrew tap and WinGet portable manifest from the same artifacts.
7. Provide direct download links plus `brew` and `winget` commands on the website.
8. Let package managers own upgrades; add a version notification before considering self-update.

## Decision

**Build Skill Issue's standalone CLI in Go and distribute signed per-platform binaries through GitHub Releases, automated by GoReleaser.** This gives the simplest user experience and the smallest release/maintenance burden while preserving a clean path to Homebrew, WinGet, Scoop, and native Linux packages.

Use **Deno compile** instead only if the implementation is already materially TypeScript-centric. Use **Rust** only if the thin orchestration CLI grows into a performance-, concurrency-, or security-heavy native system. Use **Dart** only when existing team expertise provides enough development advantage to justify native macOS and Windows build jobs.
