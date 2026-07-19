# BoolField

The BoolField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - checkboxField
  - toggleField



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CheckboxField`                                                      | [*shared.CheckboxField](../../../pkg/models/shared/checkboxfield.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DefaultValue`                                                       | `*bool`                                                              | :heavy_minus_sign:                                                   | The defaultValue field.                                              |
| `Rules`                                                              | [*shared.BoolRules](../../../pkg/models/shared/boolrules.md)         | :heavy_minus_sign:                                                   | N/A                                                                  |
| `ToggleField`                                                        | [*shared.ToggleField](../../../pkg/models/shared/togglefield.md)     | :heavy_minus_sign:                                                   | N/A                                                                  |