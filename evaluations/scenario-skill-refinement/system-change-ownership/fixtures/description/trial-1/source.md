# Export Retry Placement

The application exports reports through this path:

- `ReportController` validates the request and calls `ExportCoordinator`.
- `ExportCoordinator` creates an `ExportJob` and sends it to `JobDispatcher`.
- `JobDispatcher` selects a queue implementation and records dispatch state.
- `WorkerRunner` executes any queued job and reports success or failure.
- `AdminExportPage` polls the controller and displays job state.

The operations team asks for three retries with exponential backoff after report exports fail. The request was raised because the admin page currently shows a terminal failure immediately. A proposed quick fix adds retry timers to `AdminExportPage` and resubmits the controller request.

Other job types use the same dispatcher and worker but have different retry policies. `ExportJob` already carries export-specific timeout and priority settings.
