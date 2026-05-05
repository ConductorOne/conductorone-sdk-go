# Int64Field

The Int64Field message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - numberField



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Int64Rules`                                                     | [*shared.Int64Rules](../../../pkg/models/shared/int64rules.md)   | :heavy_minus_sign:                                               | Int64Rules describes the constraints applied to `int64` values   |
| `NumberField`                                                    | [*shared.NumberField](../../../pkg/models/shared/numberfield.md) | :heavy_minus_sign:                                               | The NumberField message.                                         |
| `DefaultValue`                                                   | `*int64`                                                         | :heavy_minus_sign:                                               | The defaultValue field.                                          |
| `Placeholder`                                                    | `*string`                                                        | :heavy_minus_sign:                                               | The placeholder field.                                           |