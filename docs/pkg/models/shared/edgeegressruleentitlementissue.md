# EdgeEgressRuleEntitlementIssue

EdgeEgressRuleEntitlementIssue is one egress rule whose entitlement does not
 resolve.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Cause`                                                          | [*shared.Cause](../../../pkg/models/shared/cause.md)             | :heavy_minus_sign:                                               | The cause field.                                                 |
| `DeletedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `Enforcement`                                                    | [*shared.Enforcement](../../../pkg/models/shared/enforcement.md) | :heavy_minus_sign:                                               | What the Edge enforces for this rule.                            |
| `EntitlementDisplayName`                                         | `*string`                                                        | :heavy_minus_sign:                                               | The entitlement's name, when its row could be read.              |
| `EntitlementID`                                                  | `*string`                                                        | :heavy_minus_sign:                                               | The entitlementId field.                                         |
| `RuleIndex`                                                      | `*int`                                                           | :heavy_minus_sign:                                               | Zero-based index into the saved rules.                           |