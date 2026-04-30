# AuthConfigOIDC

The AuthConfigOIDC message.


## Fields

| Field                       | Type                        | Required                    | Description                 |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `ExactMatchClaims`          | map[string]`string`         | :heavy_minus_sign:          | The exactMatchClaims field. |
| `IssuerID`                  | `*string`                   | :heavy_minus_sign:          | The issuerId field.         |
| `OidcClientID`              | `*string`                   | :heavy_minus_sign:          | The oidcClientId field.     |
| `OidcClientSecret`          | `*string`                   | :heavy_minus_sign:          | The oidcClientSecret field. |
| `Scopes`                    | []`string`                  | :heavy_minus_sign:          | The scopes field.           |