# BulkUpdateFindingStateRequest

The BulkUpdateFindingStateRequest message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - snooze
  - suppress
  - acceptRisk
  - unsuppress
  - assignOwner



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `BulkAcceptRiskAction`                                                               | [*shared.BulkAcceptRiskAction](../../../pkg/models/shared/bulkacceptriskaction.md)   | :heavy_minus_sign:                                                                   | The BulkAcceptRiskAction message.                                                    |
| `BulkAssignOwnerAction`                                                              | [*shared.BulkAssignOwnerAction](../../../pkg/models/shared/bulkassignowneraction.md) | :heavy_minus_sign:                                                                   | The BulkAssignOwnerAction message.                                                   |
| `BulkSnoozeAction`                                                                   | [*shared.BulkSnoozeAction](../../../pkg/models/shared/bulksnoozeaction.md)           | :heavy_minus_sign:                                                                   | The BulkSnoozeAction message.                                                        |
| `BulkSuppressAction`                                                                 | [*shared.BulkSuppressAction](../../../pkg/models/shared/bulksuppressaction.md)       | :heavy_minus_sign:                                                                   | The BulkSuppressAction message.                                                      |
| `BulkUnsuppressAction`                                                               | [*shared.BulkUnsuppressAction](../../../pkg/models/shared/bulkunsuppressaction.md)   | :heavy_minus_sign:                                                                   | The BulkUnsuppressAction message.                                                    |
| `FindingSearchRequest`                                                               | [*shared.FindingSearchRequest](../../../pkg/models/shared/findingsearchrequest.md)   | :heavy_minus_sign:                                                                   | The FindingSearchRequest message.                                                    |
| `Refs`                                                                               | [][shared.FindingRef](../../../pkg/models/shared/findingref.md)                      | :heavy_minus_sign:                                                                   | By-ID mode: specify individual finding refs.                                         |