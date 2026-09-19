# AppSecretAdminServiceRevokeResponse

The AppSecretAdminServiceRevokeResponse message.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `CredentialRevocation`                                                             | [*shared.CredentialRevocation](../../../pkg/models/shared/credentialrevocation.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Revoked`                                                                          | `*bool`                                                                            | :heavy_minus_sign:                                                                 | Successful calls queue a revoke and return true.                                   |