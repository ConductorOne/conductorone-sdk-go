# Int64Field

The Int64Field message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - numberField


This message contains a oneof named _default_value. Only a single field of the following list may be set at a time:
  - defaultValue


This message contains a oneof named _rules. Only a single field of the following list may be set at a time:
  - rules



## Fields

| Field                                                                                                                                             | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Int64Rules`                                                                                                                                      | [*shared.Int64Rules](../../../pkg/models/shared/int64rules.md)                                                                                    | :heavy_minus_sign:                                                                                                                                | Int64Rules describes the constraints applied to `int64` values                                                                                    |
| `NumberField`                                                                                                                                     | [*shared.NumberField](../../../pkg/models/shared/numberfield.md)                                                                                  | :heavy_minus_sign:                                                                                                                                | The NumberField message.                                                                                                                          |
| `DefaultValue`                                                                                                                                    | **int64*                                                                                                                                          | :heavy_minus_sign:                                                                                                                                | The defaultValue field.<br/>This field is part of the `_default_value` oneof.<br/>See the documentation for `c1.api.form.v1.Int64Field` for more details. |
| `Placeholder`                                                                                                                                     | **string*                                                                                                                                         | :heavy_minus_sign:                                                                                                                                | The placeholder field.                                                                                                                            |