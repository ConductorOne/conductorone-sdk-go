# OAuthAuthorizationConsent

OAuthAuthorizationConsent is a current consent to a managed OAuth authorization.
 Disconnected and superseded consents are not included.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `AuthorizationDisplayName`                                                             | `*string`                                                                              | :heavy_minus_sign:                                                                     | The current display name of the managed authorization.                                 |
| `AuthorizationID`                                                                      | `*string`                                                                              | :heavy_minus_sign:                                                                     | The managed authorization ID, used with grant.id when revoking this consent.           |
| `Grant`                                                                                | [*shared.ExternalOAuthGrantInfo](../../../pkg/models/shared/externaloauthgrantinfo.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |