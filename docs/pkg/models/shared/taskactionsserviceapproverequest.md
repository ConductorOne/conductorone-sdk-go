# TaskActionsServiceApproveRequest

The TaskActionsServiceApproveRequest object lets you approve a task.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Comment`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | The comment attached to the request.                                   |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PolicyStepID`                                                         | `string`                                                               | :heavy_check_mark:                                                     | The ID of the policy step on the given task to approve.                |