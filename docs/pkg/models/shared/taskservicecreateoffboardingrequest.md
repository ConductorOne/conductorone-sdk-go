# TaskServiceCreateOffboardingRequest

Create an offboarding task.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Description`                                                          | `*string`                                                              | :heavy_minus_sign:                                                     | The description of the offboarding request.                            |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `SubjectUserID`                                                        | `*string`                                                              | :heavy_minus_sign:                                                     | The ID of the user to offboard.                                        |