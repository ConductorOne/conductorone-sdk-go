# VaultInput

The Vault message.

This message contains a oneof named vault. Only a single field of the following list may be set at a time:
  - groupAuthzVault
  - magicVault



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `GroupAuthzVault`                                                        | [*shared.GroupAuthzVault](../../../pkg/models/shared/groupauthzvault.md) | :heavy_minus_sign:                                                       | The GroupAuthzVault message.                                             |
| `MagicVault`                                                             | [*shared.MagicVault](../../../pkg/models/shared/magicvault.md)           | :heavy_minus_sign:                                                       | The MagicVault message.                                                  |
| `CredentialExpirationDuration`                                           | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | The description field.                                                   |
| `DisplayName`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | The displayName field.                                                   |
| `ID`                                                                     | **string*                                                                | :heavy_minus_sign:                                                       | The id field.                                                            |