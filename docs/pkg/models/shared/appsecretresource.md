# AppSecretResource

The AppSecretResource message.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AppID`                                                   | `*string`                                                 | :heavy_minus_sign:                                        | The app that this secret belongs to.                      |
| `AppResourceTypeID`                                       | `*string`                                                 | :heavy_minus_sign:                                        | The secret type that this secret is.                      |
| `CreatedAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `DeletedAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `DisplayName`                                             | `*string`                                                 | :heavy_minus_sign:                                        | The display name for this secret.                         |
| `ID`                                                      | `*string`                                                 | :heavy_minus_sign:                                        | The id of the secret.                                     |
| `IdentityAppUserID`                                       | `*string`                                                 | :heavy_minus_sign:                                        | The ID of the user that is associated with the AppSecret. |
| `LastUsedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `SecretCreatedAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `SecretExpiresAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `UpdatedAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |