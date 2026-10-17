# GetMetricForecastResponse

GetMetricForecastResponse contains the forecast for the requested metric.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ForecastPoints`                                                                         | [][shared.ForecastRollup](../../../pkg/models/shared/forecastrollup.md)                  | :heavy_minus_sign:                                                                       | The forecasted data points (future points only).                                         |
| `HistoryDaysLoaded`                                                                      | `*int64`                                                                                 | :heavy_minus_sign:                                                                       | The number of historical days loaded as context for the model (includes hidden history). |
| `ModelName`                                                                              | `*string`                                                                                | :heavy_minus_sign:                                                                       | The model identifier used for inference.                                                 |
| `ObservedThrough`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |