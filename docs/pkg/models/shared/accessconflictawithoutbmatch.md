# AccessConflictAWithoutBMatch

AccessConflictAWithoutBMatch is the group A entitlement behind an
 access-conflict finding for a user who holds none of the monitor's group B
 entitlements.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `GroupAAppEntitlementID`                        | `*string`                                       | :heavy_minus_sign:                              | The group A entitlement the user holds.         |
| `GroupAAppID`                                   | `*string`                                       | :heavy_minus_sign:                              | The app of the group A entitlement.             |
| `GroupAAppUserID`                               | `*string`                                       | :heavy_minus_sign:                              | The account that holds the group A entitlement. |