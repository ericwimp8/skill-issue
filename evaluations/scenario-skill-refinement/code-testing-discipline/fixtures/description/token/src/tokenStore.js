export class TokenStore {
  #token;

  constructor(token) {
    this.#token = token;
  }

  redact() {
    if (this.#token.length <= 4) {
      return '*'.repeat(this.#token.length);
    }
    return `${this.#token.slice(0, 2)}${'*'.repeat(this.#token.length - 4)}${this.#token.slice(-2)}`;
  }
}
