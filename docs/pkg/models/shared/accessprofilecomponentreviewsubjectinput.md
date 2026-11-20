# AccessProfileComponentReviewSubjectInput

A component of the profile's definition.

This message contains a oneof named detail. Only a single field of the following list may be set at a time:
  - composition
  - membershipCriteria



## Fields

| Field                                                                                                                                  | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `Composition`                                                                                                                          | [*shared.AccessProfileCompositionReviewItem](../../../pkg/models/shared/accessprofilecompositionreviewitem.md)                         | :heavy_minus_sign:                                                                                                                     | N/A                                                                                                                                    |
| `MembershipCriteria`                                                                                                                   | [*shared.AccessProfileMembershipCriteriaReviewItemInput](../../../pkg/models/shared/accessprofilemembershipcriteriareviewiteminput.md) | :heavy_minus_sign:                                                                                                                     | N/A                                                                                                                                    |
| `Profile`                                                                                                                              | [*shared.AccessProfileRef](../../../pkg/models/shared/accessprofileref.md)                                                             | :heavy_minus_sign:                                                                                                                     | N/A                                                                                                                                    |