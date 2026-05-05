# CohortHintView

The CohortHintView message.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Attribute`                                          | `*string`                                            | :heavy_minus_sign:                                   | The user attribute name used for cohort grouping.    |
| `Priority`                                           | `*int`                                               | :heavy_minus_sign:                                   | Relative priority of this hint.                      |
| `Values`                                             | []`string`                                           | :heavy_minus_sign:                                   | The specific attribute values targeted by this hint. |