# Vault

The Vault message.

This message contains a oneof named vault. Only a single field of the following list may be set at a time:
  - groupAuthzVault
  - magicVault



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `GroupAuthzVault`                                                        | [*shared.GroupAuthzVault](../../../pkg/models/shared/groupauthzvault.md) | :heavy_minus_sign:                                                       | The GroupAuthzVault message.                                             |
| `MagicVault`                                                             | [*shared.MagicVault](../../../pkg/models/shared/magicvault.md)           | :heavy_minus_sign:                                                       | The MagicVault message.                                                  |
| `CreatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `CredentialExpirationDuration`                                           | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `DeletedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | The description field.                                                   |
| `DisplayName`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | The displayName field.                                                   |
| `ID`                                                                     | **string*                                                                | :heavy_minus_sign:                                                       | The id field.                                                            |
| `UpdatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |