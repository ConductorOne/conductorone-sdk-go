# TaskType

Task Type provides configuration for the type of task: certify, grant, or revoke

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify
  - offboarding
  - action
  - finding



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [*shared.TaskTypeAction](../../../pkg/models/shared/tasktypeaction.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Certify`                                                                        | [*shared.TaskTypeCertify](../../../pkg/models/shared/tasktypecertify.md)         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Finding`                                                                        | [*shared.TaskTypeFinding](../../../pkg/models/shared/tasktypefinding.md)         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Grant`                                                                          | [*shared.TaskTypeGrant](../../../pkg/models/shared/tasktypegrant.md)             | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Offboarding`                                                                    | [*shared.TaskTypeOffboarding](../../../pkg/models/shared/tasktypeoffboarding.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Revoke`                                                                         | [*shared.TaskTypeRevoke](../../../pkg/models/shared/tasktyperevoke.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |