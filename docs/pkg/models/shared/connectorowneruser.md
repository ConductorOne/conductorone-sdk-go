# ConnectorOwnerUser

ConnectorOwnerUser represents a user ownership source for a connector.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AppID`                                            | `*string`                                          | :heavy_minus_sign:                                 | The appId field.                                   |
| `ConnectorID`                                      | `*string`                                          | :heavy_minus_sign:                                 | The connectorId field.                             |
| `CreatedAt`                                        | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |
| `RoleSlug`                                         | `*string`                                          | :heavy_minus_sign:                                 | The roleSlug field.                                |
| `User`                                             | [*shared.User](../../../pkg/models/shared/user.md) | :heavy_minus_sign:                                 | N/A                                                |