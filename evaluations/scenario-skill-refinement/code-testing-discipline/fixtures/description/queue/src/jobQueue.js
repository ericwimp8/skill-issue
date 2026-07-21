export class JobQueue {
  #pending = [];

  enqueue(job) {
    this.#pending.push(job);
  }

  drain() {
    const jobs = [...this.#pending];
    this.#pending = [];
    return jobs;
  }
}
