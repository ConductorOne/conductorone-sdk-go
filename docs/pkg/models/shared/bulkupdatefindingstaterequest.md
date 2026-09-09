# BulkUpdateFindingStateRequest

The BulkUpdateFindingStateRequest message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - snooze
  - suppress
  - acceptRisk
  - unsuppress
  - assignOwner
  - reopen
  - reprocess



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AcceptRisk`                                                                         | [*shared.BulkAcceptRiskAction](../../../pkg/models/shared/bulkacceptriskaction.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `AssignOwner`                                                                        | [*shared.BulkAssignOwnerAction](../../../pkg/models/shared/bulkassignowneraction.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Refs`                                                                               | [][shared.FindingRef](../../../pkg/models/shared/findingref.md)                      | :heavy_minus_sign:                                                                   | By-ID mode: specify individual finding refs.                                         |
| `Reopen`                                                                             | [*shared.BulkReopenAction](../../../pkg/models/shared/bulkreopenaction.md)           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Reprocess`                                                                          | [*shared.BulkReprocessAction](../../../pkg/models/shared/bulkreprocessaction.md)     | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `SearchRequest`                                                                      | [*shared.FindingSearchRequest](../../../pkg/models/shared/findingsearchrequest.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Snooze`                                                                             | [*shared.BulkSnoozeAction](../../../pkg/models/shared/bulksnoozeaction.md)           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Suppress`                                                                           | [*shared.BulkSuppressAction](../../../pkg/models/shared/bulksuppressaction.md)       | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Unsuppress`                                                                         | [*shared.BulkUnsuppressAction](../../../pkg/models/shared/bulkunsuppressaction.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |