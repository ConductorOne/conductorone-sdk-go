# ServicePrincipalServiceCreateCredentialRequest

The ServicePrincipalServiceCreateCredentialRequest message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AllowSourceCidrs`                                                                   | []*string*                                                                           | :heavy_minus_sign:                                                                   | A list of CIDRs to restrict this credential to.                                      |
| `DisplayName`                                                                        | **string*                                                                            | :heavy_minus_sign:                                                                   | The display name for the new credential.                                             |
| `Expires`                                                                            | **string*                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `RequireDpop`                                                                        | **bool*                                                                              | :heavy_minus_sign:                                                                   | If true, requires DPoP proof-of-possession for token exchange using this credential. |
| `ScopedRoles`                                                                        | []*string*                                                                           | :heavy_minus_sign:                                                                   | The list of roles to restrict the credential to.                                     |