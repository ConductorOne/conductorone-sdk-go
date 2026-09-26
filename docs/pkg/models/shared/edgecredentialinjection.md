# EdgeCredentialInjection

EdgeCredentialInjection binds one C1 vault secret to one request header.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ContentType`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | The secret's stored content type. Defaults to "api_key".          |
| `HeaderName`                                                      | `*string`                                                         | :heavy_minus_sign:                                                | Header to set, replacing any value supplied by the caller.        |
| `HostPattern`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | Destination host glob, e.g. "api.example.com" or "*.example.com". |
| `PathPattern`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | Request path glob, e.g. "/v1/*". Empty matches every path.        |
| `SecretID`                                                        | `*string`                                                         | :heavy_minus_sign:                                                | Secret to inject.                                                 |
| `ValuePrefix`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | Literal text placed before the secret value, e.g. "Bearer ".      |
| `VaultID`                                                         | `*string`                                                         | :heavy_minus_sign:                                                | Vault holding the secret.                                         |