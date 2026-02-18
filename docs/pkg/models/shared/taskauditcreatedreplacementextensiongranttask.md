# TaskAuditCreatedReplacementExtensionGrantTask

TaskAuditCreatedReplacementExtensionGrantTask is used when a replacement extension grant task is created
 (e.g. when an extension grant task is cancelled due to app user deletion).


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `NewTaskID`                                                        | **string*                                                          | :heavy_minus_sign:                                                 | The ID of the newly created replacement task                       |
| `NewTaskNumericID`                                                 | **int64*                                                           | :heavy_minus_sign:                                                 | The numeric ID of the newly created replacement task (for display) |