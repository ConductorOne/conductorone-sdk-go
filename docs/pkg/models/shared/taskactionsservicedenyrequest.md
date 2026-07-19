# TaskActionsServiceDenyRequest

The TaskActionsServiceDenyRequest object lets you deny a task.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Comment`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | The comment attached to the request.                                   |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PolicyStepID`                                                         | `*string`                                                              | :heavy_minus_sign:                                                     | The ID of the current policy step. This is the step you want to deny.  |