# StringSliceField

The StringSliceField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - chipsField


This message contains a oneof named _rules. Only a single field of the following list may be set at a time:
  - rules



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ChipsField`                                                         | [*shared.ChipsField](../../../pkg/models/shared/chipsfield.md)       | :heavy_minus_sign:                                                   | The ChipsField message.                                              |
| `RepeatedRules`                                                      | [*shared.RepeatedRules](../../../pkg/models/shared/repeatedrules.md) | :heavy_minus_sign:                                                   | RepeatedRules describe the constraints applied to `repeated` values  |
| `DefaultValues`                                                      | []*string*                                                           | :heavy_minus_sign:                                                   | The defaultValues field.                                             |
| `Placeholder`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | The placeholder field.                                               |