# UserOwnerUser

UserOwnerUser represents a user ownership source for a canonical User (a service account).


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `CreatedAt`                                        | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |
| `RoleSlug`                                         | `*string`                                          | :heavy_minus_sign:                                 | The roleSlug field.                                |
| `User`                                             | [*shared.User](../../../pkg/models/shared/user.md) | :heavy_minus_sign:                                 | N/A                                                |
| `UserID`                                           | `*string`                                          | :heavy_minus_sign:                                 | The userId field.                                  |