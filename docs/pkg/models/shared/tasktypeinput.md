# TaskTypeInput

Task Type provides configuration for the type of task: certify, grant, or revoke

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify
  - offboarding
  - action
  - finding



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Action`                                                                                   | [*shared.TaskTypeActionInput](../../../pkg/models/shared/tasktypeactioninput.md)           | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Certify`                                                                                  | [*shared.TaskTypeCertifyInput](../../../pkg/models/shared/tasktypecertifyinput.md)         | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Finding`                                                                                  | [*shared.TaskTypeFindingInput](../../../pkg/models/shared/tasktypefindinginput.md)         | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Grant`                                                                                    | [*shared.TaskTypeGrantInput](../../../pkg/models/shared/tasktypegrantinput.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Offboarding`                                                                              | [*shared.TaskTypeOffboardingInput](../../../pkg/models/shared/tasktypeoffboardinginput.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Revoke`                                                                                   | [*shared.TaskTypeRevokeInput](../../../pkg/models/shared/tasktyperevokeinput.md)           | :heavy_minus_sign:                                                                         | N/A                                                                                        |