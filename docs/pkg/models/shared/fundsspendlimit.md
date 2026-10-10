# FundsSpendLimit

SpendLimit distinguishes an amount, unlimited tracking and blocked supply.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - unlimited
  - amount
  - blocked



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Amount`                                                                                   | [*shared.FundsSpendLimitAmount](../../../pkg/models/shared/fundsspendlimitamount.md)       | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Blocked`                                                                                  | [*shared.FundsSpendLimitBlocked](../../../pkg/models/shared/fundsspendlimitblocked.md)     | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Unlimited`                                                                                | [*shared.FundsSpendLimitUnlimited](../../../pkg/models/shared/fundsspendlimitunlimited.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |