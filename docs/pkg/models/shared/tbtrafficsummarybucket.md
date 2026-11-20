# TBTrafficSummaryBucket

TBTrafficSummaryBucket is one point in GetTrafficSummary's time series.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AllowedEvents`                            | `*string`                                  | :heavy_minus_sign:                         | The allowedEvents field.                   |
| `BucketStart`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `DeniedEvents`                             | `*string`                                  | :heavy_minus_sign:                         | The deniedEvents field.                    |
| `UnspecifiedEvents`                        | `*string`                                  | :heavy_minus_sign:                         | The unspecifiedEvents field.               |