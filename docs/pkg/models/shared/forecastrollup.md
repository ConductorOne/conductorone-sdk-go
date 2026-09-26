# ForecastRollup

ForecastRollup represents a single forecasted data point with uncertainty bounds.
 Note: field names are legacy (p10/p90); the actual quantile levels are p20/p50/p80.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `P10`                                                        | `*int64`                                                     | :heavy_minus_sign:                                           | The lower bound (p20) of the forecast. Field name is legacy. |
| `P50`                                                        | `*int64`                                                     | :heavy_minus_sign:                                           | The median (p50) of the forecast.                            |
| `P90`                                                        | `*int64`                                                     | :heavy_minus_sign:                                           | The upper bound (p80) of the forecast. Field name is legacy. |
| `Timestamp`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | N/A                                                          |