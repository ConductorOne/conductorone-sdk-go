# Int64Field

The Int64Field message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - numberField



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `DefaultValue`                                                   | `*int64`                                                         | :heavy_minus_sign:                                               | The defaultValue field.                                          |
| `NumberField`                                                    | [*shared.NumberField](../../../pkg/models/shared/numberfield.md) | :heavy_minus_sign:                                               | N/A                                                              |
| `Placeholder`                                                    | `*string`                                                        | :heavy_minus_sign:                                               | The placeholder field.                                           |
| `Rules`                                                          | [*shared.Int64Rules](../../../pkg/models/shared/int64rules.md)   | :heavy_minus_sign:                                               | N/A                                                              |