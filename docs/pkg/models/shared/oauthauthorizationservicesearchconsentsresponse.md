# OAuthAuthorizationServiceSearchConsentsResponse

OAuthAuthorizationServiceSearchConsentsResponse contains current consents in
 descending creation order, with consent ID breaking creation-time ties.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `List`                                                                                        | [][shared.OAuthAuthorizationConsent](../../../pkg/models/shared/oauthauthorizationconsent.md) | :heavy_minus_sign:                                                                            | Current consents matching all requested filters.                                              |
| `NextPageToken`                                                                               | `*string`                                                                                     | :heavy_minus_sign:                                                                            | Token for the next page. Empty when there are no more matching consents.                      |