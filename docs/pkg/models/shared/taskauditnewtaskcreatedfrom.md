# TaskAuditNewTaskCreatedFrom

TaskAuditNewTaskCreatedFrom is used when a task is created from another task
 (e.g. when a replacement extension grant task is created after the original is cancelled).
 This is set on the NEW task to indicate its origin.


## Fields

| Field                            | Type                             | Required                         | Description                      |
| -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| `OriginalTaskID`                 | **string*                        | :heavy_minus_sign:               | The originalTaskId field.        |
| `OriginalTaskNumericID`          | **int64*                         | :heavy_minus_sign:               | The originalTaskNumericId field. |