# AuthorityRef

AuthorityRef identifies one authority row at the revision admission read.
 The reservation already carries tenant_id; the remaining fields are the
 row's typed primary-key components. Absent rows have version zero.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Absent`                                           | `*bool`                                            | :heavy_minus_sign:                                 | The absent field.                                  |
| `AppEntitlementID`                                 | `*string`                                          | :heavy_minus_sign:                                 | The appEntitlementId field.                        |
| `AppID`                                            | `*string`                                          | :heavy_minus_sign:                                 | The appId field.                                   |
| `AppUserID`                                        | `*string`                                          | :heavy_minus_sign:                                 | The appUserId field.                               |
| `Kind`                                             | [*shared.Kind](../../../pkg/models/shared/kind.md) | :heavy_minus_sign:                                 | The kind field.                                    |
| `RuleID`                                           | `*string`                                          | :heavy_minus_sign:                                 | The ruleId field.                                  |
| `UserID`                                           | `*string`                                          | :heavy_minus_sign:                                 | The userId field.                                  |
| `Version`                                          | `*int64`                                           | :heavy_minus_sign:                                 | The version field.                                 |