# EdgeIntegrationVaultSecretRef

EdgeIntegrationVaultSecretRef names an existing modern shared Vault credential.
 The integration type determines the content type; Socket requires api_key.


## Fields

| Field                               | Type                                | Required                            | Description                         |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `SecretID`                          | `*string`                           | :heavy_minus_sign:                  | Secret identifier within the Vault. |
| `VaultBoundaryID`                   | `*string`                           | :heavy_minus_sign:                  | Vault containing the secret.        |