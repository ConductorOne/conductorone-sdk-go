# EvaluateEntitlementSelectionRequest

EvaluateEntitlementSelectionRequest selects analyzed entitlements using an
 inclusive coverage cutoff plus optional manual overrides.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ExplicitlyExcluded`                                                       | [][shared.EntitlementRef](../../../pkg/models/shared/entitlementref.md)    | :heavy_minus_sign:                                                         | Analyzed entitlements to exclude when they meet the cutoff.                |
| `ExplicitlyIncluded`                                                       | [][shared.EntitlementRef](../../../pkg/models/shared/entitlementref.md)    | :heavy_minus_sign:                                                         | Analyzed entitlements to include even when they fall below the cutoff.     |
| `IncludeFacets`                                                            | `*bool`                                                                    | :heavy_minus_sign:                                                         | Whether to return profile attribute facets for exact holders.              |
| `MinimumCoverageBasisPoints`                                               | `*int`                                                                     | :heavy_minus_sign:                                                         | Inclusive minimum entitlement coverage in basis points, where 8000 is 80%. |