export interface ArchiveWriter {
  write(data: Uint8Array, encrypted: boolean): Promise<string>;
}

export class ExportService {
  constructor(private readonly writer: ArchiveWriter) {}

  export(data: Uint8Array): Promise<string> {
    return this.writer.write(data, false);
  }
}
