# TaskAction

The TaskAction message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - close
  - reassign



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Close`                                                                    | [*shared.CloseAction](../../../pkg/models/shared/closeaction.md)           | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Reassign`                                                                 | [*shared.ReassignAction](../../../pkg/models/shared/reassignaction.md)     | :heavy_minus_sign:                                                         | N/A                                                                        |
| `TaskTypes`                                                                | [][shared.TaskTypes](../../../pkg/models/shared/tasktypes.md)              | :heavy_minus_sign:                                                         | The taskTypes field.                                                       |
| `TaskUserRelation`                                                         | [*shared.TaskUserRelation](../../../pkg/models/shared/taskuserrelation.md) | :heavy_minus_sign:                                                         | The taskUserRelation field.                                                |