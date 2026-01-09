# TaskTypeInput

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
| `TaskTypeAction`                                                                             | [*shared.TaskTypeActionInput](../../../pkg/models/shared/tasktypeactioninput.md)             | :heavy_minus_sign:                                                                           | The TaskTypeAction message.                                                                  |
| `TaskTypeCertify`                                                                            | [*shared.TaskTypeCertifyInput](../../../pkg/models/shared/tasktypecertifyinput.md)           | :heavy_minus_sign:                                                                           | The TaskTypeCertify message indicates that a task is a certify task and all related details. |
| `TaskTypeGrant`                                                                              | [*shared.TaskTypeGrantInput](../../../pkg/models/shared/tasktypegrantinput.md)               | :heavy_minus_sign:                                                                           | The TaskTypeGrant message indicates that a task is a grant task and all related details.     |
| `TaskTypeOffboarding`                                                                        | [*shared.TaskTypeOffboardingInput](../../../pkg/models/shared/tasktypeoffboardinginput.md)   | :heavy_minus_sign:                                                                           | The TaskTypeOffboarding message.                                                             |
| `TaskTypeRevoke`                                                                             | [*shared.TaskTypeRevokeInput](../../../pkg/models/shared/tasktyperevokeinput.md)             | :heavy_minus_sign:                                                                           | The TaskTypeRevoke message indicates that a task is a revoke task and all related details.   |