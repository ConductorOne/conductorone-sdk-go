# TenantAuthConfigInput

The TenantAuthConfig message.

This message contains a oneof named provider_config. Only a single field of the following list may be set at a time:
  - google
  - microsoft
  - okta
  - onelogin
  - jumpcloud
  - pingone
  - oidc
  - c1Local



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `BootstrapDomains`                                                                     | []`string`                                                                             | :heavy_minus_sign:                                                                     | Bootstrap routing: email domains that route unknown users to this config.              |
| `C1Local`                                                                              | [*shared.AuthConfigC1Local](../../../pkg/models/shared/authconfigc1local.md)           | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `DeprecationDeadline`                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `DeprecationMessage`                                                                   | `*string`                                                                              | :heavy_minus_sign:                                                                     | User-visible message shown when status=DEPRECATED.                                     |
| `DisplayName`                                                                          | `*string`                                                                              | :heavy_minus_sign:                                                                     | The displayName field.                                                                 |
| `Google`                                                                               | [*shared.AuthConfigGoogleInput](../../../pkg/models/shared/authconfiggoogleinput.md)   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `ID`                                                                                   | `*string`                                                                              | :heavy_minus_sign:                                                                     | The id field.                                                                          |
| `IsDefaultBootstrap`                                                                   | `*bool`                                                                                | :heavy_minus_sign:                                                                     | The isDefaultBootstrap field.                                                          |
| `Jumpcloud`                                                                            | [*shared.AuthConfigJumpCloud](../../../pkg/models/shared/authconfigjumpcloud.md)       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Microsoft`                                                                            | [*shared.AuthConfigMicrosoft](../../../pkg/models/shared/authconfigmicrosoft.md)       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Oidc`                                                                                 | [*shared.AuthConfigOIDC](../../../pkg/models/shared/authconfigoidc.md)                 | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Okta`                                                                                 | [*shared.AuthConfigOkta](../../../pkg/models/shared/authconfigokta.md)                 | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Onelogin`                                                                             | [*shared.AuthConfigOneLogin](../../../pkg/models/shared/authconfigonelogin.md)         | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Pingone`                                                                              | [*shared.AuthConfigPingOne](../../../pkg/models/shared/authconfigpingone.md)           | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Status`                                                                               | [*shared.TenantAuthConfigStatus](../../../pkg/models/shared/tenantauthconfigstatus.md) | :heavy_minus_sign:                                                                     | The status field.                                                                      |