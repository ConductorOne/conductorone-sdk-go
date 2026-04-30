# UpdateRoleMiningConfigRequest

The UpdateRoleMiningConfigRequest message.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `CohortHints`                                                                                         | [][shared.CohortHintInput](../../../pkg/models/shared/cohorthintinput.md)                             | :heavy_minus_sign:                                                                                    | Hints that guide the analysis to prioritize specific user attributes and values when forming cohorts. |
| `MaxSuggestions`                                                                                      | `*int`                                                                                                | :heavy_minus_sign:                                                                                    | Maximum number of suggestions the analysis should produce per run.                                    |
| `MinCohortSize`                                                                                       | `*int`                                                                                                | :heavy_minus_sign:                                                                                    | Minimum number of users a cohort must contain to generate a suggestion.                               |