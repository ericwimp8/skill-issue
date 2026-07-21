#!/usr/bin/env node

import { readFileSync } from 'node:fs';

const argumentsByName = new Map();
for (let index = 2; index < process.argv.length; index += 2) {
  argumentsByName.set(process.argv[index], process.argv[index + 1]);
}

const inputPath = argumentsByName.get('--input');
const targetPath = argumentsByName.get('--target');

if (!inputPath || !targetPath) {
  console.error(
    'usage: export-codex-evidence.mjs --input <rollout.jsonl> --target <candidate-path>',
  );
  process.exit(2);
}

const events = readFileSync(inputPath, 'utf8')
  .split('\n')
  .filter(Boolean)
  .map((line) => JSON.parse(line));

const sessionEvent = events.find((event) => event.type === 'session_meta');
const turnEvent = events.find((event) => event.type === 'turn_context');
const candidateEvent = events.find(
  (event) =>
    event.type === 'response_item' &&
    event.payload?.type === 'custom_tool_call' &&
    event.payload?.name === 'exec' &&
    event.payload?.input?.includes(targetPath),
);

if (!sessionEvent || !candidateEvent) {
  console.error(
    'input does not contain the required session and candidate-read evidence',
  );
  process.exit(1);
}

const completedResponse = events
  .filter(
    (event) =>
      event.type === 'event_msg' &&
      event.payload?.type === 'task_complete' &&
      event.payload?.last_agent_message,
  )
  .at(-1)?.payload.last_agent_message;

const assistantResponses = events.flatMap((event) => {
  if (
    event.type !== 'response_item' ||
    event.payload?.type !== 'message' ||
    event.payload?.role !== 'assistant'
  ) {
    return [];
  }

  return (event.payload.content ?? [])
    .filter((content) => content.type === 'output_text')
    .map((content) => content.text);
});

const finalResponse = completedResponse ?? assistantResponses.at(-1);
if (!finalResponse) {
  console.error('input does not contain an observable assistant response');
  process.exit(1);
}

function sanitize(text) {
  return text
    .replace(/<oai-mem-citation>[\s\S]*?<\/oai-mem-citation>/g, '')
    .replace(/\/Users\/[^/\s]+/g, '~')
    .replace(/[A-Za-z]:\\Users\\[^\\\s]+/g, '~')
    .trim();
}

const session = sessionEvent.payload;
const spawn = session.source?.subagent?.thread_spawn;
const evidence = {
  format: 'curated-codex-evaluation-evidence/v1',
  evidence_scope: 'Pre-output candidate read and observable final response.',
  source_session_id: session.id,
  agent_path: session.agent_path ?? spawn?.agent_path,
  cli_version: session.cli_version,
  model: turnEvent?.payload?.model,
  reasoning: turnEvent?.payload?.effort,
  candidate_read: {
    timestamp: candidateEvent.timestamp,
    tool_name: candidateEvent.payload.name,
    call_id: candidateEvent.payload.call_id,
    target_path: targetPath,
  },
  final_response: sanitize(finalResponse),
};

console.log(JSON.stringify(evidence));
