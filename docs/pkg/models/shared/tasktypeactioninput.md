# TaskTypeActionInput

The TaskTypeAction message.

This message contains a oneof named target_object. Only a single field of the following list may be set at a time:
  - scopeRole
  - toolCall
  - finding



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ActionInstance`                                                                           | [*shared.TaskActionInstanceInput](../../../pkg/models/shared/taskactioninstanceinput.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Finding`                                                                                  | [*shared.FindingTargetInput](../../../pkg/models/shared/findingtargetinput.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ScopeRole`                                                                                | [*shared.ScopeRoleInput](../../../pkg/models/shared/scoperoleinput.md)                     | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ToolCall`                                                                                 | [*shared.GatedToolCallTargetInput](../../../pkg/models/shared/gatedtoolcalltargetinput.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |