# ForecastPoint

ForecastPoint represents projected spend for one future UTC day.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Day`                                      | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `P20Nano`                                  | `*int64`                                   | :heavy_minus_sign:                         | Low spend projection for the day.          |
| `P50Nano`                                  | `*int64`                                   | :heavy_minus_sign:                         | Median spend projection for the day.       |
| `P80Nano`                                  | `*int64`                                   | :heavy_minus_sign:                         | High spend projection for the day.         |