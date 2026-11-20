# TaskTypeActionInput

The TaskTypeAction message.

This message contains a oneof named target_object. Only a single field of the following list may be set at a time:
  - scopeRole
  - toolCall
  - finding
  - credentialIssue
  - accessProfileComponent



## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `AccessProfileComponent`                                                                                     | [*shared.AccessProfileComponentTargetInput](../../../pkg/models/shared/accessprofilecomponenttargetinput.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ActionInstance`                                                                                             | [*shared.TaskActionInstanceInput](../../../pkg/models/shared/taskactioninstanceinput.md)                     | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `CredentialIssue`                                                                                            | [*shared.CredentialIssueTargetInput](../../../pkg/models/shared/credentialissuetargetinput.md)               | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `Finding`                                                                                                    | [*shared.FindingTargetInput](../../../pkg/models/shared/findingtargetinput.md)                               | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ScopeRole`                                                                                                  | [*shared.ScopeRoleInput](../../../pkg/models/shared/scoperoleinput.md)                                       | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `Source`                                                                                                     | [*shared.TaskActionSourceInput](../../../pkg/models/shared/taskactionsourceinput.md)                         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ToolCall`                                                                                                   | [*shared.GatedToolCallTargetInput](../../../pkg/models/shared/gatedtoolcalltargetinput.md)                   | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |