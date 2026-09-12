# FundsSpendExtension

SpendExtension replaces the base total temporarily; it is not additive credit.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ExpiresAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Limit`                                                                  | [*shared.FundsSpendLimit](../../../pkg/models/shared/fundsspendlimit.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Reason`                                                                 | `*string`                                                                | :heavy_minus_sign:                                                       | The reason field.                                                        |