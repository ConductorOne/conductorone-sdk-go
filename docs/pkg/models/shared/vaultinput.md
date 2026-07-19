# VaultInput

Vault represents an external secret storage integration used to store connector credentials securely.

This message contains a oneof named vault. Only a single field of the following list may be set at a time:
  - groupAuthzVault
  - magicVault



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CredentialExpirationDuration`                                           | `*string`                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | A free-text description of the vault's purpose or configuration.         |
| `DisplayName`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The human-readable name of the vault.                                    |
| `GroupAuthzVault`                                                        | [*shared.GroupAuthzVault](../../../pkg/models/shared/groupauthzvault.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `ID`                                                                     | `*string`                                                                | :heavy_minus_sign:                                                       | The unique identifier of the vault.                                      |
| `MagicVault`                                                             | [*shared.MagicVault](../../../pkg/models/shared/magicvault.md)           | :heavy_minus_sign:                                                       | N/A                                                                      |