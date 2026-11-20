# ReviewSubjectInput

The explicitly stored item under review. Legacy tasks may have no subject.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - entitlementAccess
  - profileComponent



## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `EntitlementAccess`                                                                                                        | [*shared.EntitlementAccessReviewSubject](../../../pkg/models/shared/entitlementaccessreviewsubject.md)                     | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `ProfileComponent`                                                                                                         | [*shared.AccessProfileComponentReviewSubjectInput](../../../pkg/models/shared/accessprofilecomponentreviewsubjectinput.md) | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |