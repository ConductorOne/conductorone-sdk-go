# CredentialIssueTarget

CredentialIssueTarget describes one approved credential request: who receives
 the credential, the terms asked for, and the offering it came from.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Duration`                                                            | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `IdentityAppUserID`                                                   | `*string`                                                             | :heavy_minus_sign:                                                    | The app user tying the recipient to the connector identity.           |
| `IdentityResourceID`                                                  | `*string`                                                             | :heavy_minus_sign:                                                    | The connector-side identity the credential is minted against.         |
| `IdentityUserID`                                                      | `*string`                                                             | :heavy_minus_sign:                                                    | The user who receives the credential and may open its delivery vault. |
| `OfferingID`                                                          | `*string`                                                             | :heavy_minus_sign:                                                    | The offering the requester selected.                                  |
| `RequestCatalogID`                                                    | `*string`                                                             | :heavy_minus_sign:                                                    | The Access Profile that published it.                                 |
| `Scopes`                                                              | []`string`                                                            | :heavy_minus_sign:                                                    | Provider permissions approved for this credential.                    |