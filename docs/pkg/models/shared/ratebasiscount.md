# RateBasisCount

RateBasisCount summarizes settled calls that used one recorded price basis.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Calls`                                                                          | `*int64`                                                                         | :heavy_minus_sign:                                                               | Number of settled calls that used this price basis.                              |
| `SettledNano`                                                                    | `*int64`                                                                         | :heavy_minus_sign:                                                               | Total settled spend for these calls.                                             |
| `Snapshot`                                                                       | [*shared.ClawLLMRateSnapshot](../../../pkg/models/shared/clawllmratesnapshot.md) | :heavy_minus_sign:                                                               | N/A                                                                              |