# CreateStepUpProviderRequest

The CreateStepUpProviderRequest message.

This message contains a oneof named settings. Only a single field of the following list may be set at a time:
  - oauth2
  - microsoft



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ClientID`                                                                               | `*string`                                                                                | :heavy_minus_sign:                                                                       | The OAuth2 client ID used to authenticate with the step-up provider.                     |
| `ClientSecret`                                                                           | `*string`                                                                                | :heavy_minus_sign:                                                                       | The OAuth2 client secret. Write-only; never returned in responses.                       |
| `DisplayName`                                                                            | `*string`                                                                                | :heavy_minus_sign:                                                                       | The human-readable name for the new step-up provider.                                    |
| `IssuerURL`                                                                              | `*string`                                                                                | :heavy_minus_sign:                                                                       | The OIDC issuer URL for the step-up provider.                                            |
| `Microsoft`                                                                              | [*shared.StepUpMicrosoftSettings](../../../pkg/models/shared/stepupmicrosoftsettings.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Oauth2`                                                                                 | [*shared.StepUpOAuth2Settings](../../../pkg/models/shared/stepupoauth2settings.md)       | :heavy_minus_sign:                                                                       | N/A                                                                                      |