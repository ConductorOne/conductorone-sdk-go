# TaskActionsServiceSkipStepRequest

The TaskActionsServiceSkipStepRequest object lets you skip a policy step in a task.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Comment`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | The comment attached to the request.                                   |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PolicyStepID`                                                         | `string`                                                               | :heavy_check_mark:                                                     | The ID of the policy step to skip.                                     |