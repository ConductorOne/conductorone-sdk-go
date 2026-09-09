# TrajectoryPoint

TrajectoryPoint represents settled spend for one UTC day.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ConsumedNano`                                                               | `*int64`                                                                     | :heavy_minus_sign:                                                           | Spend settled during the day, in billionths of the configured currency unit. |
| `Day`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                   | :heavy_minus_sign:                                                           | N/A                                                                          |