# VaultServiceCreateRequest

VaultServiceCreateRequest is the request message for creating a new vault.

This message contains a oneof named vault. Only a single field of the following list may be set at a time:
  - groupAuthzVault
  - magicVault



## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `GroupAuthzVault`                                                                                               | [*shared.GroupAuthzVault](../../../pkg/models/shared/groupauthzvault.md)                                        | :heavy_minus_sign:                                                                                              | GroupAuthzVault configures a vault that uses group-based authorization to control access to stored credentials. |
| `MagicVault`                                                                                                    | [*shared.MagicVault](../../../pkg/models/shared/magicvault.md)                                                  | :heavy_minus_sign:                                                                                              | MagicVault configures a vault that grants time-limited credential access via magic links.                       |
| `Description`                                                                                                   | `*string`                                                                                                       | :heavy_minus_sign:                                                                                              | A free-text description of the vault's purpose or configuration.                                                |
| `DisplayName`                                                                                                   | `string`                                                                                                        | :heavy_check_mark:                                                                                              | The human-readable name for the new vault.                                                                      |
| `OwnerIds`                                                                                                      | []`string`                                                                                                      | :heavy_minus_sign:                                                                                              | The IDs of users to assign as owners of this vault.                                                             |