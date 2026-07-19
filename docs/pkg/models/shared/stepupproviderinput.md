# StepUpProviderInput

StepUpProvider represents a configured step-up authentication integration (e.g., Duo, custom OIDC).

This message contains a oneof named settings. Only a single field of the following list may be set at a time:
  - oauth2
  - microsoft



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ClientID`                                                                               | `*string`                                                                                | :heavy_minus_sign:                                                                       | The OAuth2 client ID used to authenticate with the step-up provider.                     |
| `DisplayName`                                                                            | `*string`                                                                                | :heavy_minus_sign:                                                                       | The human-readable name of the step-up provider.                                         |
| `Enabled`                                                                                | `*bool`                                                                                  | :heavy_minus_sign:                                                                       | Whether the step-up provider is active and available for use.                            |
| `IssuerURL`                                                                              | `*string`                                                                                | :heavy_minus_sign:                                                                       | The OIDC issuer URL for the step-up provider.                                            |
| `Microsoft`                                                                              | [*shared.StepUpMicrosoftSettings](../../../pkg/models/shared/stepupmicrosoftsettings.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Oauth2`                                                                                 | [*shared.StepUpOAuth2Settings](../../../pkg/models/shared/stepupoauth2settings.md)       | :heavy_minus_sign:                                                                       | N/A                                                                                      |