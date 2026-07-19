# TaskServiceCreateActionRequest

The TaskServiceCreateActionRequest message submits a request action (requestable automation).


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ActionID`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | The ID of the action to request.                                       |
| `Description`                                                          | `*string`                                                              | :heavy_minus_sign:                                                     | An optional description of the request.                                |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `FormValues`                                                           | map[string]`any`                                                       | :heavy_minus_sign:                                                     | N/A                                                                    |