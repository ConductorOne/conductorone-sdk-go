# EntitlementAccessReviewSubject

Access held by an account or resource principal.

This message contains a oneof named principal. Only a single field of the following list may be set at a time:
  - appUser
  - resource



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppUser`                                                                    | [*shared.AppUserRef](../../../pkg/models/shared/appuserref.md)               | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Entitlement`                                                                | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Resource`                                                                   | [*shared.AppResourceRef](../../../pkg/models/shared/appresourceref.md)       | :heavy_minus_sign:                                                           | N/A                                                                          |