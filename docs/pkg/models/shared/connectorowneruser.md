# ConnectorOwnerUser

ConnectorOwnerUser represents a user ownership source for a connector.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `User`                                                                                  | [*shared.User](../../../pkg/models/shared/user.md)                                      | :heavy_minus_sign:                                                                      | The User object provides all of the details for an user, as well as some configuration. |
| `CreatedAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `RoleSlug`                                                                              | `*string`                                                                               | :heavy_minus_sign:                                                                      | The roleSlug field.                                                                     |