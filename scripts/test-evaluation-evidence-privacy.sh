#!/bin/sh

set -eu

repo_root=$(git rev-parse --show-toplevel)
temporary_root=$(mktemp -d "${TMPDIR:-/tmp}/skill-issue-privacy-test.XXXXXX")
trap 'rm -rf "$temporary_root"' EXIT

raw_session="$temporary_root/raw.jsonl"
curated_evidence="$temporary_root/curated.jsonl"

cat > "$raw_session" <<'EOF'
{"timestamp":"2026-07-21T00:00:00Z","type":"session_meta","payload":{"id":"safe-session","cli_version":"test","agent_path":"/root/eval","base_instructions":{"text":"private ambient context"},"source":{"subagent":{"thread_spawn":{"agent_path":"/root/eval"}}}}}
{"timestamp":"2026-07-21T00:00:01Z","type":"turn_context","payload":{"model":"test-model","effort":"medium"}}
{"timestamp":"2026-07-21T00:00:02Z","type":"response_item","payload":{"type":"custom_tool_call","name":"exec","call_id":"safe-call","input":"cat supporting-skills/example/SKILL.md"}}
{"timestamp":"2026-07-21T00:00:03Z","type":"response_item","payload":{"type":"message","role":"assistant","content":[{"type":"output_text","text":"Safe evaluation result at /Users/private-name/project.\n<oai-mem-citation>private memory pointer</oai-mem-citation>"}]}}
EOF

node "$repo_root/plugins/skill-issue/skills/skill-evaluation-and-refinement/scripts/export-codex-evidence.mjs" \
  --input "$raw_session" \
  --target supporting-skills/example/SKILL.md > "$curated_evidence"

jq -e '
  .format == "curated-codex-evaluation-evidence/v1" and
  .source_session_id == "safe-session" and
  .candidate_read.target_path == "supporting-skills/example/SKILL.md" and
  .final_response == "Safe evaluation result at ~/project."
' "$curated_evidence" >/dev/null

if grep -E -q 'session_meta|turn_context|base_instructions|private ambient context' "$curated_evidence"; then
  echo "test failed: exporter retained raw Codex context" >&2
  exit 1
fi

test_repo="$temporary_root/repository"
mkdir -p "$test_repo/scripts"
cp "$repo_root/scripts/check-repository-privacy.sh" "$test_repo/scripts/"
cp "$curated_evidence" "$test_repo/evidence.jsonl"
printf 'private-identity\n' > "$test_repo/.privacy-denylist.local"

git -C "$test_repo" init -q
git -C "$test_repo" config user.name "Eric Wimp"
git -C "$test_repo" config user.email "ericwimp8@example.invalid"
git -C "$test_repo" add scripts/check-repository-privacy.sh evidence.jsonl

(cd "$test_repo" && ./scripts/check-repository-privacy.sh) >/dev/null

cp "$raw_session" "$test_repo/evidence.jsonl"
git -C "$test_repo" add evidence.jsonl
if (cd "$test_repo" && ./scripts/check-repository-privacy.sh) >/dev/null 2>&1; then
  echo "test failed: privacy check accepted raw Codex context" >&2
  exit 1
fi

echo "evaluation evidence privacy tests passed"
