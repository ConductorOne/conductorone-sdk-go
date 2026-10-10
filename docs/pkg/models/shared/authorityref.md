# AuthorityRef

AuthorityRef identifies a configuration and the revision admission read.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Absent`                                                         | `*bool`                                                          | :heavy_minus_sign:                                               | Whether admission found no configuration of this kind.           |
| `AppEntitlementID`                                               | `*string`                                                        | :heavy_minus_sign:                                               | Application entitlement ID, for entitlement-binding authorities. |
| `AppID`                                                          | `*string`                                                        | :heavy_minus_sign:                                               | Application ID the configuration applies to.                     |
| `AppUserID`                                                      | `*string`                                                        | :heavy_minus_sign:                                               | Application user ID, for application-user authorities.           |
| `Kind`                                                           | [*shared.Kind](../../../pkg/models/shared/kind.md)               | :heavy_minus_sign:                                               | Type of configuration.                                           |
| `RuleID`                                                         | `*string`                                                        | :heavy_minus_sign:                                               | Fund rule ID, for fund-rule authorities.                         |
| `UserID`                                                         | `*string`                                                        | :heavy_minus_sign:                                               | User ID the configuration applies to.                            |
| `Version`                                                        | `*int64`                                                         | :heavy_minus_sign:                                               | Configuration revision read during admission.                    |