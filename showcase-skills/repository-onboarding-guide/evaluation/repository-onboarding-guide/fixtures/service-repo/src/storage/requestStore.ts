export interface RequestStore {
  append(payload: string): Promise<void>;
}

