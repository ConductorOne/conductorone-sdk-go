# TaskAuditListResponse

The TaskAuditListResponse message.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `List`                                                                | [][shared.TaskAuditView](../../../pkg/models/shared/taskauditview.md) | :heavy_minus_sign:                                                    | The list of audit events for the task.                                |
| `NextPageToken`                                                       | `*string`                                                             | :heavy_minus_sign:                                                    | A pagination token to retrieve the next page of results.              |