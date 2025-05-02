# TaskAuditConnectorActionResult

The TaskAuditConnectorActionResult message.

This message contains a oneof named result. Only a single field of the following list may be set at a time:
  - success
  - error
  - cancelled



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `TaskAuditCancelledResult`                                                                 | [*shared.TaskAuditCancelledResult](../../../pkg/models/shared/taskauditcancelledresult.md) | :heavy_minus_sign:                                                                         | The TaskAuditCancelledResult message.                                                      |
| `TaskAuditErrorResult`                                                                     | [*shared.TaskAuditErrorResult](../../../pkg/models/shared/taskauditerrorresult.md)         | :heavy_minus_sign:                                                                         | The TaskAuditErrorResult message.                                                          |
| `TaskAuditSuccessResult`                                                                   | [*shared.TaskAuditSuccessResult](../../../pkg/models/shared/taskauditsuccessresult.md)     | :heavy_minus_sign:                                                                         | The TaskAuditSuccessResult message.                                                        |
| `AppEntitlementID`                                                                         | **string*                                                                                  | :heavy_minus_sign:                                                                         | The appEntitlementId field.                                                                |
| `AppID`                                                                                    | **string*                                                                                  | :heavy_minus_sign:                                                                         | The appId field.                                                                           |
| `ConnectorActionID`                                                                        | **string*                                                                                  | :heavy_minus_sign:                                                                         | The connectorActionId field.                                                               |
| `ConnectorID`                                                                              | **string*                                                                                  | :heavy_minus_sign:                                                                         | The connectorId field.                                                                     |