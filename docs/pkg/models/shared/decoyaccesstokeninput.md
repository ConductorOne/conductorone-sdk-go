# DecoyAccessTokenInput

DecoyAccessTokenInput mints a session access-token decoy under an
 existing User.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ExpiresIn`                                                | `*string`                                                  | :heavy_minus_sign:                                         | N/A                                                        |
| `SubjectUserID`                                            | `*string`                                                  | :heavy_minus_sign:                                         | Existing User the access token's subject claim references. |