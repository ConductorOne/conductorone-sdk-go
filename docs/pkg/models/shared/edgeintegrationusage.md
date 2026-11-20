# EdgeIntegrationUsage

EdgeIntegrationUsage identifies saved checks that reference an integration.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CheckIds`                                                | []`string`                                                | :heavy_minus_sign:                                        | Referencing check identifiers, including disabled checks. |
| `EdgeID`                                                  | `*string`                                                 | :heavy_minus_sign:                                        | Edge containing the checks.                               |