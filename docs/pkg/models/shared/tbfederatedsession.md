# TBFederatedSession

TBFederatedSession is TB's secret-free hosted OAuth session lineage for a
 cryptographically attested call (see AuditFederatedSession in the TB
 sink). grant_jti_hash and proof_key_thumbprint are a hash and a public key
 thumbprint, not secrets.


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `ClientID`                    | `*string`                     | :heavy_minus_sign:            | The clientId field.           |
| `GrantJtiHash`                | `*string`                     | :heavy_minus_sign:            | The grantJtiHash field.       |
| `IssuerName`                  | `*string`                     | :heavy_minus_sign:            | The issuerName field.         |
| `ProofKeyThumbprint`          | `*string`                     | :heavy_minus_sign:            | The proofKeyThumbprint field. |
| `Resource`                    | `*string`                     | :heavy_minus_sign:            | The resource field.           |
| `Scopes`                      | []`string`                    | :heavy_minus_sign:            | The scopes field.             |
| `SessionID`                   | `*string`                     | :heavy_minus_sign:            | The sessionId field.          |
| `TenantID`                    | `*string`                     | :heavy_minus_sign:            | The tenantId field.           |