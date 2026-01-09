# GrantsByCriteriaScope

The GrantsByCriteriaScope message.

This message contains a oneof named criteria_filter. Only a single field of the following list may be set at a time:
  - daysSinceAdded
  - daysSinceReviewed
  - grantsAddedBetween



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `GrantAccessProfileFilter`                                                                 | [*shared.GrantAccessProfileFilter](../../../pkg/models/shared/grantaccessprofilefilter.md) | :heavy_minus_sign:                                                                         | The GrantAccessProfileFilter message.                                                      |
| `GrantsAddedBetween`                                                                       | [*shared.GrantsAddedBetween](../../../pkg/models/shared/grantsaddedbetween.md)             | :heavy_minus_sign:                                                                         | The GrantsAddedBetween message.                                                            |
| `DaysSinceAdded`                                                                           | **string*                                                                                  | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `DaysSinceLastUsed`                                                                        | **string*                                                                                  | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `DaysSinceReviewed`                                                                        | **string*                                                                                  | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `SourceFilter`                                                                             | [*shared.SourceFilter](../../../pkg/models/shared/sourcefilter.md)                         | :heavy_minus_sign:                                                                         | The sourceFilter field.                                                                    |
| `TypeFilter`                                                                               | [*shared.TypeFilter](../../../pkg/models/shared/typefilter.md)                             | :heavy_minus_sign:                                                                         | The typeFilter field.                                                                      |