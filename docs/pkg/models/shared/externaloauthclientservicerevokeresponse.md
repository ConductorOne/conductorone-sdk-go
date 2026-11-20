# ExternalOAuthClientServiceRevokeResponse

The ExternalOAuthClientServiceRevokeResponse message.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Complete`                                                                                                 | `*bool`                                                                                                    | :heavy_minus_sign:                                                                                         | False when grants remain to retire. The client is already revoked; call<br/> Revoke again to finish the sweep. |