# DisabledReasonCircuitBreaker

DisabledReasonCircuitBreaker carries the trip context when an automation
 has been auto-disabled by its rate cap. Returned on the parent Automation
 when read; not directly settable.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ObservedCount`                                        | `*int64`                                               | :heavy_minus_sign:                                     | Observed execution count in the period at trip time.   |
| `Period`                                               | [*shared.Period](../../../pkg/models/shared/period.md) | :heavy_minus_sign:                                     | Snapshot of the period at trip time.                   |
| `Threshold`                                            | `*int64`                                               | :heavy_minus_sign:                                     | Snapshot of the threshold at trip time.                |
| `TrippedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |