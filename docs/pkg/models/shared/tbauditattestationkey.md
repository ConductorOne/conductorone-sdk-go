# TBAuditAttestationKey

TBAuditAttestationKey is a trusted manifest signing key.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `KeyID`                                                                                          | `*string`                                                                                        | :heavy_minus_sign:                                                                               | First 16 bytes of the SHA-256 of the raw public key, hex.                                        |
| `PublicKey`                                                                                      | `*string`                                                                                        | :heavy_minus_sign:                                                                               | Raw Ed25519 public key, 32 bytes, hex.                                                           |
| `Status`                                                                                         | [*shared.TBAuditAttestationKeyStatus](../../../pkg/models/shared/tbauditattestationkeystatus.md) | :heavy_minus_sign:                                                                               | The status field.                                                                                |