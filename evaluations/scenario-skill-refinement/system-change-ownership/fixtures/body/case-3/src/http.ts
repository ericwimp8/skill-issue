import { ExportService } from './exporter';

export class ExportRoute {
  constructor(private readonly exports: ExportService) {}

  post(data: Uint8Array): Promise<string> {
    return this.exports.export(data);
  }
}
