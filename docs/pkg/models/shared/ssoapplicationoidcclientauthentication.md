# SSOApplicationOIDCClientAuthentication

SSOApplicationOIDCClientAuthentication is the exact token-endpoint client
 authentication method assigned to an OIDC client.

This message contains a oneof named method. Only a single field of the following list may be set at a time:
  - none
  - clientSecretBasic
  - clientSecretPost
  - privateKeyJwt



## Fields

| Field                                                                                                                                | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ClientSecretBasic`                                                                                                                  | [*shared.SSOApplicationOIDCClientAuthClientSecretBasic](../../../pkg/models/shared/ssoapplicationoidcclientauthclientsecretbasic.md) | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |
| `ClientSecretPost`                                                                                                                   | [*shared.SSOApplicationOIDCClientAuthClientSecretPost](../../../pkg/models/shared/ssoapplicationoidcclientauthclientsecretpost.md)   | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |
| `None`                                                                                                                               | [*shared.SSOApplicationOIDCClientAuthNone](../../../pkg/models/shared/ssoapplicationoidcclientauthnone.md)                           | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |
| `PrivateKeyJwt`                                                                                                                      | [*shared.SSOApplicationOIDCClientAuthPrivateKeyJWT](../../../pkg/models/shared/ssoapplicationoidcclientauthprivatekeyjwt.md)         | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |