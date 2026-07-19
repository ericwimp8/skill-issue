# Internal Vendor Telemetry Reporting Service: A-to-B Plan

## A: Starting Position

- Vendor API credentials exist before the planned work begins, and the security team manages their rotation outside this plan.
- Analysts currently download vendor telemetry data each week.
- Analysts manually assemble the vendor telemetry report and decide whether its data is fresh.

## B: Desired Position

An internal reporting service replaces the manually assembled vendor telemetry report by ingesting the vendor data, classifying its freshness, exposing the report's status to analysts, and supporting analyst review before distribution.

## Path from A to B

1. **Establish the internal report contract**
   Define the report's telemetry boundaries and production cadence, the freshness rules and source signals, the lifecycle and status semantics, and the conditions that complete analyst review and authorize distribution.

2. **Create vendor telemetry ingestion**
   Build the service capability that uses the existing vendor access to acquire the required telemetry and records each run's source timestamp and completion or failure state alongside the report data needed by later processing.

3. **Classify report freshness**
   Add freshness evaluation to the ingested report data so that the service determines and records the report's freshness state consistently.

4. **Expose observable report status**
   Give analysts a service-owned view that exposes each ingestion run's source timestamp, freshness classification, and completion or failure state, together with the generated report's availability.

5. **Enable analyst review before distribution**
   Place each generated report into an analyst review stage and prevent it from being distributed until that review is complete.

6. **Replace the weekly manual reporting workflow**
   Move analysts from weekly downloads and manual freshness decisions to the internal service as the operational source for the vendor telemetry report.

## C: Observable Outcomes

- The service ingests the vendor telemetry required for the internal report.
- Every ingestion run exposes its source timestamp, service-determined freshness classification, and completion or failure state.
- Each generated report is held for analyst review and is distributed only after that review is complete.
- Analysts can obtain the vendor telemetry report without performing the existing weekly download-and-assembly workflow.
