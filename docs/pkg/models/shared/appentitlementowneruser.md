# AppEntitlementOwnerUser

AppEntitlementOwnerUser represents a user ownership source for an app entitlement.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AppID`                                            | `*string`                                          | :heavy_minus_sign:                                 | The appId field.                                   |
| `CreatedAt`                                        | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |
| `EntitlementID`                                    | `*string`                                          | :heavy_minus_sign:                                 | The entitlementId field.                           |
| `RoleSlug`                                         | `*string`                                          | :heavy_minus_sign:                                 | The roleSlug field.                                |
| `User`                                             | [*shared.User](../../../pkg/models/shared/user.md) | :heavy_minus_sign:                                 | N/A                                                |