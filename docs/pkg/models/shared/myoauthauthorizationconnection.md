# MyOAuthAuthorizationConnection

MyOAuthAuthorizationConnection is a user's current consent, including connections whose access was subsequently removed.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Authorization`                                                                    | [*shared.MyOAuthAuthorization](../../../pkg/models/shared/myoauthauthorization.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `CreatedAt`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `GrantID`                                                                          | `*string`                                                                          | :heavy_minus_sign:                                                                 | The grantId field.                                                                 |
| `LastUsedAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `SelectedRoleIds`                                                                  | []`string`                                                                         | :heavy_minus_sign:                                                                 | The selectedRoleIds field.                                                         |