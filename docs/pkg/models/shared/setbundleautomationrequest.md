# SetBundleAutomationRequest

The SetBundleAutomationRequest message.

This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
  - entitlements



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `BundleAutomationRuleEntitlement`                                                                        | [*shared.BundleAutomationRuleEntitlement](../../../pkg/models/shared/bundleautomationruleentitlement.md) | :heavy_minus_sign:                                                                                       | The BundleAutomationRuleEntitlement message.                                                             |
| `CreateTasks`                                                                                            | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | The createTasks field.                                                                                   |
| `DisableCircuitBreaker`                                                                                  | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | The disableCircuitBreaker field.                                                                         |
| `Enabled`                                                                                                | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | The enabled field.                                                                                       |