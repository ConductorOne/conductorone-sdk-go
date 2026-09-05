# FundRuleServiceSearchRequest

The FundRuleServiceSearchRequest message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `PageSize`                                                             | `*int`                                                                 | :heavy_minus_sign:                                                     | The pageSize field.                                                    |
| `PageToken`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | The pageToken field.                                                   |
| `Query`                                                                | `*string`                                                              | :heavy_minus_sign:                                                     | Case-insensitive search over the rule display name; empty returns all. |