export function nextJob(jobs) {
  return jobs.sort((left, right) => left.priority > right.priority)[0];
}
