# OAuthAuthorizationClientAuthentication

OAuthAuthorizationClientAuthentication describes the authentication method
 assigned to an OAuth authorization. It never contains a client secret.

This message contains a oneof named method. Only a single field of the following list may be set at a time:
  - clientSecretBasic
  - clientSecretPost
  - privateKeyJwt



## Fields

| Field                                                                                                                                | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ClientSecretBasic`                                                                                                                  | [*shared.OAuthAuthorizationClientAuthClientSecretBasic](../../../pkg/models/shared/oauthauthorizationclientauthclientsecretbasic.md) | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |
| `ClientSecretPost`                                                                                                                   | [*shared.OAuthAuthorizationClientAuthClientSecretPost](../../../pkg/models/shared/oauthauthorizationclientauthclientsecretpost.md)   | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |
| `PrivateKeyJwt`                                                                                                                      | [*shared.OAuthAuthorizationClientAuthPrivateKeyJWT](../../../pkg/models/shared/oauthauthorizationclientauthprivatekeyjwt.md)         | :heavy_minus_sign:                                                                                                                   | N/A                                                                                                                                  |