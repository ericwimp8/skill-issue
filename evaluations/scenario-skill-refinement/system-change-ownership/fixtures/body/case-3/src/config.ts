export type ExportPolicy = {
  encryptionRequired: boolean;
};

export class AppConfig {
  constructor(readonly exportPolicy: ExportPolicy) {}
}
