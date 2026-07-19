# Action

The Action message.

This message contains a oneof named target. Only a single field of the following list may be set at a time:
  - automation
  - batonResourceAction
  - clientIdApproval



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Automation`                                                                                             | [*shared.ActionTargetAutomation](../../../pkg/models/shared/actiontargetautomation.md)                   | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `BatonResourceAction`                                                                                    | [*shared.ActionTargetBatonResourceAction](../../../pkg/models/shared/actiontargetbatonresourceaction.md) | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ClientIDApproval`                                                                                       | [*shared.ActionTargetClientIDApproval](../../../pkg/models/shared/actiontargetclientidapproval.md)       | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |