# DatadogContext

The DatadogContext message.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `RumSessionID`                             | `*string`                                  | :heavy_minus_sign:                         | The rumSessionId field.                    |
| `RumViewID`                                | `*string`                                  | :heavy_minus_sign:                         | The rumViewId field.                       |
| `RumViewName`                              | `*string`                                  | :heavy_minus_sign:                         | The rumViewName field.                     |
| `TraceIds`                                 | []`string`                                 | :heavy_minus_sign:                         | The traceIds field.                        |
| `WindowEnd`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `WindowStart`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |