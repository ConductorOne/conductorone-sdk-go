# TaskAuditConnectorActionResult

The TaskAuditConnectorActionResult message.

This message contains a oneof named result. Only a single field of the following list may be set at a time:
  - success
  - error
  - cancelled
  - pending



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `AppEntitlementID`                                                                         | `*string`                                                                                  | :heavy_minus_sign:                                                                         | The appEntitlementId field.                                                                |
| `AppID`                                                                                    | `*string`                                                                                  | :heavy_minus_sign:                                                                         | The appId field.                                                                           |
| `Cancelled`                                                                                | [*shared.TaskAuditCancelledResult](../../../pkg/models/shared/taskauditcancelledresult.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ConnectorActionID`                                                                        | `*string`                                                                                  | :heavy_minus_sign:                                                                         | The connectorActionId field.                                                               |
| `ConnectorID`                                                                              | `*string`                                                                                  | :heavy_minus_sign:                                                                         | The connectorId field.                                                                     |
| `Error`                                                                                    | [*shared.TaskAuditErrorResult](../../../pkg/models/shared/taskauditerrorresult.md)         | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Pending`                                                                                  | [*shared.TaskAuditPendingResult](../../../pkg/models/shared/taskauditpendingresult.md)     | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Success`                                                                                  | [*shared.TaskAuditSuccessResult](../../../pkg/models/shared/taskauditsuccessresult.md)     | :heavy_minus_sign:                                                                         | N/A                                                                                        |