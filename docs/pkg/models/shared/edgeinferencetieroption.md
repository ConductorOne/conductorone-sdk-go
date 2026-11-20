# EdgeInferenceTierOption

EdgeInferenceTierOption is a C1 tier and the model it uses for the caller's tenant.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AliasRouteID`                                     | `*string`                                          | :heavy_minus_sign:                                 | Route id published as the model alias.             |
| `ModelDisplayName`                                 | `*string`                                          | :heavy_minus_sign:                                 | The modelDisplayName field.                        |
| `ModelID`                                          | `*string`                                          | :heavy_minus_sign:                                 | Registry model id the tier uses.                   |
| `RouteID`                                          | `*string`                                          | :heavy_minus_sign:                                 | Route id published for the tier, e.g. c1-fast.     |
| `Tier`                                             | [*shared.Tier](../../../pkg/models/shared/tier.md) | :heavy_minus_sign:                                 | The tier field.                                    |