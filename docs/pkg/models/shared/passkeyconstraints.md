# PasskeyConstraints

PasskeyConstraints controls how users may enroll passkeys (FIDO2 / WebAuthn).


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `AllowedAaguids`                                                                   | []`string`                                                                         | :heavy_minus_sign:                                                                 | Allowed authenticator models, by AAGUID. Leave empty to permit any<br/> authenticator. |
| `Attestation`                                                                      | [*shared.Attestation](../../../pkg/models/shared/attestation.md)                   | :heavy_minus_sign:                                                                 | How strictly the authenticator's origin must be attested.                          |
| `RequireUserVerification`                                                          | `*bool`                                                                            | :heavy_minus_sign:                                                                 | Require the authenticator to verify the user (PIN or biometric) at enrollment.     |