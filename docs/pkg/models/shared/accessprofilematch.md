# AccessProfileMatch

The AccessProfileMatch message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `CatalogDisplayName`                                                          | `*string`                                                                     | :heavy_minus_sign:                                                            | The catalogDisplayName field.                                                 |
| `CatalogID`                                                                   | `*string`                                                                     | :heavy_minus_sign:                                                            | The catalogId field.                                                          |
| `MatchType`                                                                   | [*shared.MatchType](../../../pkg/models/shared/matchtype.md)                  | :heavy_minus_sign:                                                            | The matchType field.                                                          |
| `MissingEntitlements`                                                         | [][shared.CohortEntitlement](../../../pkg/models/shared/cohortentitlement.md) | :heavy_minus_sign:                                                            | The missingEntitlements field.                                                |
| `OverlapRatio`                                                                | `*float64`                                                                    | :heavy_minus_sign:                                                            | The overlapRatio field.                                                       |