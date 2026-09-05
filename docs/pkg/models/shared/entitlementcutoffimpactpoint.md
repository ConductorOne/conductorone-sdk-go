# EntitlementCutoffImpactPoint

EntitlementCutoffImpactPoint reports the exact effect of an inclusive
 entitlement coverage cutoff on the analyzed cohort.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `EntitlementCount`                                                         | `*int`                                                                     | :heavy_minus_sign:                                                         | Number of analyzed entitlements included at this cutoff.                   |
| `MinimumCoverageBasisPoints`                                               | `*int`                                                                     | :heavy_minus_sign:                                                         | Inclusive minimum entitlement coverage in basis points, where 8000 is 80%. |
| `UsersWithAllEntitlements`                                                 | `*int`                                                                     | :heavy_minus_sign:                                                         | Exact number of cohort users who hold every included entitlement.          |