# SecretAgingFilter

SecretAgingFilter restricts a resource search to secrets (credential_type != 0)
 whose secret-trait timestamps fall in the given half-open ranges. Each bound is
 optional; leave one unset for an open-ended range. All set bounds are ANDed.
 Callers pass absolute timestamps (computed against their reference "now").


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `LastUsedAfter`                            | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `LastUsedBefore`                           | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `SecretCreatedAfter`                       | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `SecretCreatedBefore`                      | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `SecretExpiresAfter`                       | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `SecretExpiresBefore`                      | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |