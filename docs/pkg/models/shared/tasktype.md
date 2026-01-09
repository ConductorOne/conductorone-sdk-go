# TaskType

Task Type provides configuration for the type of task: certify, grant, or revoke

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify
  - offboarding
  - action



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `TaskTypeAction`                                                                             | [*shared.TaskTypeAction](../../../pkg/models/shared/tasktypeaction.md)                       | :heavy_minus_sign:                                                                           | The TaskTypeAction message.                                                                  |
| `TaskTypeCertify`                                                                            | [*shared.TaskTypeCertify](../../../pkg/models/shared/tasktypecertify.md)                     | :heavy_minus_sign:                                                                           | The TaskTypeCertify message indicates that a task is a certify task and all related details. |
| `TaskTypeGrant`                                                                              | [*shared.TaskTypeGrant](../../../pkg/models/shared/tasktypegrant.md)                         | :heavy_minus_sign:                                                                           | The TaskTypeGrant message indicates that a task is a grant task and all related details.     |
| `TaskTypeOffboarding`                                                                        | [*shared.TaskTypeOffboarding](../../../pkg/models/shared/tasktypeoffboarding.md)             | :heavy_minus_sign:                                                                           | The TaskTypeOffboarding message.                                                             |
| `TaskTypeRevoke`                                                                             | [*shared.TaskTypeRevoke](../../../pkg/models/shared/tasktyperevoke.md)                       | :heavy_minus_sign:                                                                           | The TaskTypeRevoke message indicates that a task is a revoke task and all related details.   |