# StringSliceField

The StringSliceField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - chipsField
  - pickerField



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ChipsField`                                                         | [*shared.ChipsField](../../../pkg/models/shared/chipsfield.md)       | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DefaultValues`                                                      | []`string`                                                           | :heavy_minus_sign:                                                   | The defaultValues field.                                             |
| `PickerField`                                                        | [*shared.PickerField](../../../pkg/models/shared/pickerfield.md)     | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Placeholder`                                                        | `*string`                                                            | :heavy_minus_sign:                                                   | The placeholder field.                                               |
| `Rules`                                                              | [*shared.RepeatedRules](../../../pkg/models/shared/repeatedrules.md) | :heavy_minus_sign:                                                   | N/A                                                                  |