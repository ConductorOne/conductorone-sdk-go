# TaskAuditListRequest

The TaskAuditListRequest message.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `PageSize`                                                                                    | `*int`                                                                                        | :heavy_minus_sign:                                                                            | The maximum number of audit events to return per page.                                        |
| `PageToken`                                                                                   | `*string`                                                                                     | :heavy_minus_sign:                                                                            | A pagination token from a previous response to retrieve the next page.                        |
| `Refs`                                                                                        | [][shared.TaskAuditViewRef](../../../pkg/models/shared/taskauditviewref.md)                   | :heavy_minus_sign:                                                                            | References to specific audit events to retrieve. If provided, only these events are returned. |
| `TaskID`                                                                                      | `*string`                                                                                     | :heavy_minus_sign:                                                                            | The ID of the task to list audit events for.                                                  |