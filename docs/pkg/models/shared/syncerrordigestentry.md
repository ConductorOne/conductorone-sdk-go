# SyncErrorDigestEntry

SyncErrorDigestEntry groups occurrences sharing the same redacted text.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Count`                                                                | `*int`                                                                 | :heavy_minus_sign:                                                     | Occurrences observed in the bounded scan, not an exhaustive log count. |
| `ErrorClass`                                                           | [*shared.ErrorClass](../../../pkg/models/shared/errorclass.md)         | :heavy_minus_sign:                                                     | Likely failure family.                                                 |
| `FirstSeenAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `LastSeenAt`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Step`                                                                 | `*string`                                                              | :heavy_minus_sign:                                                     | Known sync step name, absent if it cannot be established.              |
| `Summary`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | Distinct redacted message, capped at 256 bytes.                        |