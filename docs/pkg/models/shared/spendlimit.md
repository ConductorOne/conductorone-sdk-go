# SpendLimit

SpendLimit is the three-way behavior fork. Which arms are legal depends on the
 scope carrying it; pkg/funds enforces that matrix, not the schema, because one
 SpendControls shape is shared by every scope.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - unlimited
  - amount
  - blocked



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Amount`                                                                         | [*shared.SpendLimitAmount](../../../pkg/models/shared/spendlimitamount.md)       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Blocked`                                                                        | [*shared.SpendLimitBlocked](../../../pkg/models/shared/spendlimitblocked.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Unlimited`                                                                      | [*shared.SpendLimitUnlimited](../../../pkg/models/shared/spendlimitunlimited.md) | :heavy_minus_sign:                                                               | N/A                                                                              |