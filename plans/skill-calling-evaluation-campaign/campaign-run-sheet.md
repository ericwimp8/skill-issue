# Campaign Run Sheet — All 30 Commands

Operator run sheet for executing the 30 campaign evaluations by hand. Model identifiers below were resolved from the installed Cursor agent's `--list-models` on 2026-07-22 — the catalog drifts daily, so re-check before running Cursor lanes on a later date (prefer an exact `-medium` variant when one appears). Sol and Fable on Cursor had no medium variant and use the recorded fallbacks; note them as deviations from the medium target when recording results.

Ground rules: within a lane run `-01` → `-02` → `-03` in order; never two `claude-code` runs at once; the two Claude lanes go last, Fable dead last. Beyond those invariants, run as many lanes in parallel as you judge the machine and provider accounts can sustain — the orchestration prompt sets no fixed concurrency cap. Five lanes bill the same OpenAI account, so if rate-limit-shaped failures appear, reduce parallelism instead of retrying hot. Prefix a command with `caffeinate -i` if the machine might sleep. Success ends with `"status": "complete"` and results in `chat-<n>/output/`; a failure names its `failure.json`.

## Setup (paste first in every terminal)

```zsh
cd "$(git rev-parse --show-toplevel)"   # run from anywhere inside the repository
REPO=$(pwd); CHATS="$(dirname "$REPO")/chats"
CUR="$REPO/.skill-issue/cursor/home/.local/bin/agent"
for n in {6..35}; do mkdir -p "$CHATS/chat-$n/workspace" "$CHATS/chat-$n/output"; done
run() { "$REPO/cli/scripts/local-cli.sh" evaluate run --workspace "$CHATS/chat-$1/workspace" --output "$CHATS/chat-$1/output" --events --transcript --yes "${@:2}"; }
```

## Cursor — Composer (`composer-2.5`)

```zsh
run 6  --harness cursor --executable "$CUR" --model composer-2.5 --evaluation gardening-web-application                    # CUR-COM-01
run 8  --harness cursor --executable "$CUR" --model composer-2.5 --evaluation community-archive-desktop-application       # CUR-COM-02
run 9  --harness cursor --executable "$CUR" --model composer-2.5 --evaluation neighborhood-emergency-preparedness-program # CUR-COM-03
```

## Cursor — Grok (exact medium: `cursor-grok-4.5-medium`)

```zsh
run 7  --harness cursor --executable "$CUR" --model cursor-grok-4.5-medium --evaluation gardening-web-application                    # CUR-GRO-01
run 10 --harness cursor --executable "$CUR" --model cursor-grok-4.5-medium --evaluation community-archive-desktop-application       # CUR-GRO-02
run 11 --harness cursor --executable "$CUR" --model cursor-grok-4.5-medium --evaluation neighborhood-emergency-preparedness-program # CUR-GRO-03
```

## Cursor — Sol (fallback: `gpt-5.6-sol-high`)

```zsh
run 12 --harness cursor --executable "$CUR" --model gpt-5.6-sol-high --evaluation gardening-web-application                    # CUR-COD-01
run 13 --harness cursor --executable "$CUR" --model gpt-5.6-sol-high --evaluation community-archive-desktop-application       # CUR-COD-02
run 14 --harness cursor --executable "$CUR" --model gpt-5.6-sol-high --evaluation neighborhood-emergency-preparedness-program # CUR-COD-03
```

## Cursor — Fable (fallback: `claude-fable-5-thinking-high`)

```zsh
run 15 --harness cursor --executable "$CUR" --model claude-fable-5-thinking-high --evaluation gardening-web-application                    # CUR-FAB-01
run 16 --harness cursor --executable "$CUR" --model claude-fable-5-thinking-high --evaluation community-archive-desktop-application       # CUR-FAB-02
run 17 --harness cursor --executable "$CUR" --model claude-fable-5-thinking-high --evaluation neighborhood-emergency-preparedness-program # CUR-FAB-03
```

## OpenAI Codex — Sol (defaults)

```zsh
run 18 --harness codex --executable "$(command -v codex)" --evaluation gardening-web-application                    # COD-SOL-01
run 19 --harness codex --executable "$(command -v codex)" --evaluation community-archive-desktop-application       # COD-SOL-02
run 20 --harness codex --executable "$(command -v codex)" --evaluation neighborhood-emergency-preparedness-program # COD-SOL-03
```

## Pi — Codex (defaults)

```zsh
run 21 --harness pi --executable "$(command -v pi)" --evaluation gardening-web-application                    # PI-COD-01
run 22 --harness pi --executable "$(command -v pi)" --evaluation community-archive-desktop-application       # PI-COD-02
run 23 --harness pi --executable "$(command -v pi)" --evaluation neighborhood-emergency-preparedness-program # PI-COD-03
```

## OpenCode — Codex (env prefix required)

```zsh
XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share" run 24 --harness opencode --executable "$REPO/.skill-issue/opencode/bin/opencode" --evaluation gardening-web-application                    # OPE-COD-01
XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share" run 25 --harness opencode --executable "$REPO/.skill-issue/opencode/bin/opencode" --evaluation community-archive-desktop-application       # OPE-COD-02
XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share" run 26 --harness opencode --executable "$REPO/.skill-issue/opencode/bin/opencode" --evaluation neighborhood-emergency-preparedness-program # OPE-COD-03
```

## Kilo — Codex (env prefix required; real binary, never `.skill-issue/kilo/bin/kilo`)

```zsh
XDG_DATA_HOME="$REPO/.skill-issue/kilo/home/.local/share" run 27 --harness kilo-code --executable "$REPO/.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo" --evaluation gardening-web-application                    # KIL-COD-01
XDG_DATA_HOME="$REPO/.skill-issue/kilo/home/.local/share" run 28 --harness kilo-code --executable "$REPO/.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo" --evaluation community-archive-desktop-application       # KIL-COD-02
XDG_DATA_HOME="$REPO/.skill-issue/kilo/home/.local/share" run 29 --harness kilo-code --executable "$REPO/.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo" --evaluation neighborhood-emergency-preparedness-program # KIL-COD-03
```

## Claude Code — Codex (proxy launcher; one Claude run at a time)

```zsh
run 30 --harness claude-code --executable "$REPO/.skill-issue/claudex/claudex" --model gpt-5.6-sol --evaluation gardening-web-application                    # CLA-COD-01
run 31 --harness claude-code --executable "$REPO/.skill-issue/claudex/claudex" --model gpt-5.6-sol --evaluation community-archive-desktop-application       # CLA-COD-02
run 32 --harness claude-code --executable "$REPO/.skill-issue/claudex/claudex" --model gpt-5.6-sol --evaluation neighborhood-emergency-preparedness-program # CLA-COD-03
```

Between the Claude lanes: stop the claudex-owned proxy with its `manage` script and confirm no proxy process or localhost listener remains before starting Fable.

## Claude Code — Fable (dead last, sequential)

```zsh
run 33 --harness claude-code --executable "$(command -v claude)" --model claude-fable-5 --evaluation gardening-web-application                    # CLA-FAB-01
run 34 --harness claude-code --executable "$(command -v claude)" --model claude-fable-5 --evaluation community-archive-desktop-application       # CLA-FAB-02
run 35 --harness claude-code --executable "$(command -v claude)" --model claude-fable-5 --evaluation neighborhood-emergency-preparedness-program # CLA-FAB-03
```

## Container map

`chat-6..9` Composer · `chat-7,10,11` Grok · `chat-12..14` Sol · `chat-15..17` Cursor-Fable · `chat-18..20` Codex · `chat-21..23` Pi · `chat-24..26` OpenCode · `chat-27..29` Kilo · `chat-30..32` Claude-Codex · `chat-33..35` Claude-Fable. Record each result in `evaluation-progress.md` with its container reference (`<chats>/chat-<n>`).
