# TaskServiceCreateResourceActionRequest

The TaskServiceCreateResourceActionRequest submits a request to execute a connector resource-create action, for example creating a group from a group template.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ActionID`                                                             | `string`                                                               | :heavy_check_mark:                                                     | The ID of the resource-create action to execute.                       |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `FormValues`                                                           | map[string]`any`                                                       | :heavy_minus_sign:                                                     | N/A                                                                    |