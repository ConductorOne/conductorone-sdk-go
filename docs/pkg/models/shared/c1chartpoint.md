# C1ChartPoint

C1ChartPoint is a single (timestamp, value) observation.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Ts`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | N/A                                                               |
| `Value`                                                           | `*float64`                                                        | :heavy_minus_sign:                                                | Bounds double as a NaN/Inf rejection: NaN fails every comparison. |