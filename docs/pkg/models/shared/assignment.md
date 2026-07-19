# Assignment

Assignment is one principal (user) assigned to a session policy.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Source`                                                       | [*shared.Source](../../../pkg/models/shared/source.md)         | :heavy_minus_sign:                                             | Whether the assignment is direct or conferred through a group. |
| `UserID`                                                       | `*string`                                                      | :heavy_minus_sign:                                             | The assigned user's ID.                                        |