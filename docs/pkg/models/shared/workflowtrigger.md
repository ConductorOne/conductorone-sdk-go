# WorkflowTrigger

The WorkflowTrigger message.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - manual
  - userProfileChange



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ManualWorkflowTrigger`                                                                    | [*shared.ManualWorkflowTrigger](../../../pkg/models/shared/manualworkflowtrigger.md)       | :heavy_minus_sign:                                                                         | The ManualWorkflowTrigger message.                                                         |
| `UserProfileChangeTrigger`                                                                 | [*shared.UserProfileChangeTrigger](../../../pkg/models/shared/userprofilechangetrigger.md) | :heavy_minus_sign:                                                                         | The UserProfileChangeTrigger message.                                                      |