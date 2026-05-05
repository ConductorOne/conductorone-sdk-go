# CohortHintInput

The CohortHintInput message.


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Attribute`                                                                                             | `*string`                                                                                               | :heavy_minus_sign:                                                                                      | The user attribute name to use for cohort grouping (e.g., "department", "job_title").                   |
| `Priority`                                                                                              | `*int`                                                                                                  | :heavy_minus_sign:                                                                                      | Relative priority of this hint. Higher values cause the analysis to weight this attribute more heavily. |
| `Values`                                                                                                | []`string`                                                                                              | :heavy_minus_sign:                                                                                      | Specific attribute values to focus on. If empty, all values for the attribute are considered.           |