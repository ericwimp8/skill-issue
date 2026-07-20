import type { RequestStore } from "../storage/requestStore.js";

export function createRequestHandler(store: RequestStore) {
  return async function handleRequest(request: Request): Promise<Response> {
    if (request.method !== "POST") return new Response("Not found", { status: 404 });
    try {
      await store.append(await request.text());
      return new Response("Accepted", { status: 202 });
    } catch {
      return new Response("Unavailable", { status: 503 });
    }
  };
}

