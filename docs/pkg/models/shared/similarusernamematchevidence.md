# SimilarUsernameMatchEvidence

The SimilarUsernameMatchEvidence message.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AppUsername`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | The username on the app user account at detection time.           |
| `IdentityUsername`                                                | `*string`                                                         | :heavy_minus_sign:                                                | The username on the identity user that matched at detection time. |
| `SimilarityScore`                                                 | `*float64`                                                        | :heavy_minus_sign:                                                | Similarity score [0.0, 1.0] at detection time.                    |