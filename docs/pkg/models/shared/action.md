# Action

The Action message.

This message contains a oneof named target. Only a single field of the following list may be set at a time:
  - automation
  - batonResourceAction
  - clientIdApproval



## Fields

| Field                                                                                                                                 | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ActionTargetAutomation`                                                                                                              | [*shared.ActionTargetAutomation](../../../pkg/models/shared/actiontargetautomation.md)                                                | :heavy_minus_sign:                                                                                                                    | ActionTargetAutomation targets automation templates for policy actions.                                                               |
| `ActionTargetBatonResourceAction`                                                                                                     | [*shared.ActionTargetBatonResourceAction](../../../pkg/models/shared/actiontargetbatonresourceaction.md)                              | :heavy_minus_sign:                                                                                                                    | ActionTargetResource targets resource actions for policy actions.                                                                     |
| `ActionTargetClientIDApproval`                                                                                                        | [*shared.ActionTargetClientIDApproval](../../../pkg/models/shared/actiontargetclientidapproval.md)                                    | :heavy_minus_sign:                                                                                                                    | ActionTargetClientIdApproval targets administrator review of an external<br/> OAuth client registration (CIMD or DCR) for policy actions. |