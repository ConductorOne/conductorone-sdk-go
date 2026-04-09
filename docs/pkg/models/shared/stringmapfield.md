# StringMapField

The StringMapField message.

This message contains a oneof named _rules. Only a single field of the following list may be set at a time:
  - rules



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `StringMapRules`                                                       | [*shared.StringMapRules](../../../pkg/models/shared/stringmaprules.md) | :heavy_minus_sign:                                                     | The StringMapRules message.                                            |
| `DefaultValue`                                                         | map[string]`string`                                                    | :heavy_minus_sign:                                                     | The defaultValue field.                                                |