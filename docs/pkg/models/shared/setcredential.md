# SetCredential

SetCredential submits a RotateCredentials baton task to the target connector,
 re-encrypting the given password CEL expression with the connector's public JWK.

This message contains a oneof named connector_identifier. Only a single field of the following list may be set at a time:
  - connectorRef



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AccountIDCel`                                                     | `*string`                                                          | :heavy_minus_sign:                                                 | The accountIdCel field.                                            |
| `ConnectorRef`                                                     | [*shared.ConnectorRef](../../../pkg/models/shared/connectorref.md) | :heavy_minus_sign:                                                 | N/A                                                                |
| `PasswordCel`                                                      | `*string`                                                          | :heavy_minus_sign:                                                 | The passwordCel field.                                             |