import { ExportService } from './exporter';

export class ExportCommand {
  constructor(private readonly exports: ExportService) {}

  run(data: Uint8Array): Promise<string> {
    return this.exports.export(data);
  }
}
