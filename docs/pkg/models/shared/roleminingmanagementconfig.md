# RoleMiningManagementConfig

The RoleMiningManagementConfig message.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `CohortHints`                                                                      | [][shared.CohortHintView](../../../pkg/models/shared/cohorthintview.md)            | :heavy_minus_sign:                                                                 | Configured cohort hints that guide which user attributes the analysis prioritizes. |
| `MaxSuggestions`                                                                   | `*int`                                                                             | :heavy_minus_sign:                                                                 | Maximum number of suggestions the analysis will produce per run.                   |
| `MinCohortSize`                                                                    | `*int`                                                                             | :heavy_minus_sign:                                                                 | Minimum number of users a cohort must contain to generate a suggestion.            |