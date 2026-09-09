# AppCapServiceSetLimitRequest

The AppCapServiceSetLimitRequest message.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Limit`                                                                      | [*shared.SpendLimit](../../../pkg/models/shared/spendlimit.md)               | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Period`                                                                     | [*shared.Period](../../../pkg/models/shared/period.md)                       | :heavy_minus_sign:                                                           | Optional period override. Only valid together with the limit it denominates. |