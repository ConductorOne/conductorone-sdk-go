# SpendSuspension

SpendSuspension freezes a scope without erasing the limit it must restore
 on unsuspend, which is why it lives beside the SpendLimit oneof rather than
 inside it.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Reason`                                   | `*string`                                  | :heavy_minus_sign:                         | The reason field.                          |
| `SuspendedAt`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |