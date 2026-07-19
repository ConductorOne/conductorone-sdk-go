# ActionInstance

The ActionInstance message.

This message contains a oneof named target_instance. Only a single field of the following list may be set at a time:
  - automation
  - batonResourceActionInstance
  - clientIdApprovalInstance


This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - success
  - denied
  - error
  - cancelled



## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Action`                                                                                                                 | [*shared.Action](../../../pkg/models/shared/action.md)                                                                   | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Automation`                                                                                                             | [*shared.ActionTargetAutomationInstance](../../../pkg/models/shared/actiontargetautomationinstance.md)                   | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `BatonResourceActionInstance`                                                                                            | [*shared.ActionTargetBatonResourceActionInstance](../../../pkg/models/shared/actiontargetbatonresourceactioninstance.md) | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Cancelled`                                                                                                              | [*shared.ActionOutcomeCancelled](../../../pkg/models/shared/actionoutcomecancelled.md)                                   | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `ClientIDApprovalInstance`                                                                                               | [*shared.ActionTargetClientIDApprovalInstance](../../../pkg/models/shared/actiontargetclientidapprovalinstance.md)       | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Denied`                                                                                                                 | [*shared.ActionOutcomeDenied](../../../pkg/models/shared/actionoutcomedenied.md)                                         | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Error`                                                                                                                  | [*shared.ActionOutcomeError](../../../pkg/models/shared/actionoutcomeerror.md)                                           | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `State`                                                                                                                  | [*shared.ActionInstanceState](../../../pkg/models/shared/actioninstancestate.md)                                         | :heavy_minus_sign:                                                                                                       | The current state of the action execution.                                                                               |
| `Success`                                                                                                                | [*shared.ActionOutcomeSuccess](../../../pkg/models/shared/actionoutcomesuccess.md)                                       | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |