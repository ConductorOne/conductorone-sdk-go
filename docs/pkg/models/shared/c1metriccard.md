# C1MetricCard

C1MetricCard is one aggregate stat: label, formatted value, optional delta
 and sparkline trend.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Delta`                                                                  | `*string`                                                                | :heavy_minus_sign:                                                       | The delta field.                                                         |
| `DeltaSentiment`                                                         | [*shared.DeltaSentiment](../../../pkg/models/shared/deltasentiment.md)   | :heavy_minus_sign:                                                       | The deltaSentiment field.                                                |
| `Label`                                                                  | `*string`                                                                | :heavy_minus_sign:                                                       | The label field.                                                         |
| `Sparkline`                                                              | []`float64`                                                              | :heavy_minus_sign:                                                       | Optional trend values, oldest first. Bounds double as NaN/Inf rejection. |
| `Value`                                                                  | `*string`                                                                | :heavy_minus_sign:                                                       | The value field.                                                         |