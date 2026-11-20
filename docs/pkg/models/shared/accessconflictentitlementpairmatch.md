# AccessConflictEntitlementPairMatch

AccessConflictEntitlementPairMatch is the entitlement pair behind an
 access-conflict finding.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `GroupAAppEntitlementID`                        | `*string`                                       | :heavy_minus_sign:                              | The group A entitlement the user holds.         |
| `GroupAAppID`                                   | `*string`                                       | :heavy_minus_sign:                              | The app of the group A entitlement.             |
| `GroupAAppUserID`                               | `*string`                                       | :heavy_minus_sign:                              | The account that holds the group A entitlement. |
| `GroupBAppEntitlementID`                        | `*string`                                       | :heavy_minus_sign:                              | The group B entitlement the user holds.         |
| `GroupBAppID`                                   | `*string`                                       | :heavy_minus_sign:                              | The app of the group B entitlement.             |
| `GroupBAppUserID`                               | `*string`                                       | :heavy_minus_sign:                              | The account that holds the group B entitlement. |