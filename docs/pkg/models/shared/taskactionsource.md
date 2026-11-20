# TaskActionSource

The TaskActionSource message indicates what created an action task.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Review`                                                                               | [*shared.TaskActionSourceReview](../../../pkg/models/shared/taskactionsourcereview.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |