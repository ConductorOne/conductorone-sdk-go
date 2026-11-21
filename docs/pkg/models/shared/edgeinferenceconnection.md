# EdgeInferenceConnection

EdgeInferenceConnection is a Vault-backed credential for one inference catalog provider.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `CatalogEntryID`                                                            | `*string`                                                                   | :heavy_minus_sign:                                                          | Catalog entry, e.g. anthropic. Unique within the Edge.                      |
| `Region`                                                                    | `*string`                                                                   | :heavy_minus_sign:                                                          | The region field.                                                           |
| `SecretID`                                                                  | `*string`                                                                   | :heavy_minus_sign:                                                          | The secretId field.                                                         |
| `VariantID`                                                                 | `*string`                                                                   | :heavy_minus_sign:                                                          | Catalog variant of the entry; it sets the provider profile and destination. |
| `VaultBoundaryID`                                                           | `*string`                                                                   | :heavy_minus_sign:                                                          | The vaultBoundaryId field.                                                  |