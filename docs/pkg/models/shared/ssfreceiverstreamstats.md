# SSFReceiverStreamStats

SSFReceiverStreamStats is a lightweight read-only stats object.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `EventsActedOnCount`                       | `*int64`                                   | :heavy_minus_sign:                         | The eventsActedOnCount field.              |
| `EventsFailedCount`                        | `*int64`                                   | :heavy_minus_sign:                         | The eventsFailedCount field.               |
| `EventsReceivedCount`                      | `*int64`                                   | :heavy_minus_sign:                         | The eventsReceivedCount field.             |
| `LastErrorAt`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `LastErrorMessage`                         | `*string`                                  | :heavy_minus_sign:                         | The lastErrorMessage field.                |
| `LastEventReceivedAt`                      | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `LastVerifiedAt`                           | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `StreamID`                                 | `*string`                                  | :heavy_minus_sign:                         | The streamId field.                        |
| `TransmitterStatus`                        | `*string`                                  | :heavy_minus_sign:                         | The transmitterStatus field.               |
| `TransmitterStatusReason`                  | `*string`                                  | :heavy_minus_sign:                         | The transmitterStatusReason field.         |