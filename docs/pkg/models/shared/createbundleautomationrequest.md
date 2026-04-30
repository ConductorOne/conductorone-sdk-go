# CreateBundleAutomationRequest

The request message for creating a new bundle automation rule on a catalog.

This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
  - entitlements
  - cel



## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `BundleAutomationRuleCEL`                                                                                         | [*shared.BundleAutomationRuleCEL](../../../pkg/models/shared/bundleautomationrulecel.md)                          | :heavy_minus_sign:                                                                                                | The BundleAutomationRuleCEL message.                                                                              |
| `BundleAutomationRuleEntitlement`                                                                                 | [*shared.BundleAutomationRuleEntitlement](../../../pkg/models/shared/bundleautomationruleentitlement.md)          | :heavy_minus_sign:                                                                                                | The BundleAutomationRuleEntitlement message.                                                                      |
| `CreateTasks`                                                                                                     | `*bool`                                                                                                           | :heavy_minus_sign:                                                                                                | Whether to create access request tasks for matched users instead of granting directly.                            |
| `DisableCircuitBreaker`                                                                                           | `*bool`                                                                                                           | :heavy_minus_sign:                                                                                                | Whether to disable the circuit breaker that pauses the automation when excessive membership changes are detected. |
| `Enabled`                                                                                                         | `*bool`                                                                                                           | :heavy_minus_sign:                                                                                                | Whether the automation should actively run on its schedule.                                                       |