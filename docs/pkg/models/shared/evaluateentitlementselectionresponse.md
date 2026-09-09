# EvaluateEntitlementSelectionResponse

EvaluateEntitlementSelectionResponse contains the exact impact of the
 resolved entitlement selection.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `CoreHolderFacets`                                                              | [][shared.AttributeFacet](../../../pkg/models/shared/attributefacet.md)         | :heavy_minus_sign:                                                              | Profile attribute facets narrowed to users who hold every selected entitlement. |
| `SelectedEntitlementCount`                                                      | `*int`                                                                          | :heavy_minus_sign:                                                              | Number of entitlements in the resolved selection.                               |
| `UsersWithAllEntitlements`                                                      | `*int`                                                                          | :heavy_minus_sign:                                                              | Exact number of cohort users who hold every selected entitlement.               |