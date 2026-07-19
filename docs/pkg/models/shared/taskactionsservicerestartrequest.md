# TaskActionsServiceRestartRequest

The TaskActionsServiceRestartRequest object lets you restart a task.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Comment`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | The comment attached to the request.                                   |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PolicyStepID`                                                         | `*string`                                                              | :heavy_minus_sign:                                                     | Deprecated. This field is accepted but does not affect behavior.       |