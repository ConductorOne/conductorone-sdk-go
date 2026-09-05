# FormStringField

The StringField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - textField
  - passwordField
  - selectField
  - pickerField
  - dateField



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `DateField`                                                          | [*shared.DateField](../../../pkg/models/shared/datefield.md)         | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DefaultValue`                                                       | `*string`                                                            | :heavy_minus_sign:                                                   | The defaultValue field.                                              |
| `PasswordField`                                                      | [*shared.PasswordField](../../../pkg/models/shared/passwordfield.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `PickerField`                                                        | [*shared.PickerField](../../../pkg/models/shared/pickerfield.md)     | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Placeholder`                                                        | `*string`                                                            | :heavy_minus_sign:                                                   | The placeholder field.                                               |
| `Rules`                                                              | [*shared.StringRules](../../../pkg/models/shared/stringrules.md)     | :heavy_minus_sign:                                                   | N/A                                                                  |
| `SelectField`                                                        | [*shared.SelectField](../../../pkg/models/shared/selectfield.md)     | :heavy_minus_sign:                                                   | N/A                                                                  |
| `TextField`                                                          | [*shared.TextField](../../../pkg/models/shared/textfield.md)         | :heavy_minus_sign:                                                   | N/A                                                                  |