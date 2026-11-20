# FormStringMapField

The StringMapField message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `DefaultValue`                                                                       | map[string]`string`                                                                  | :heavy_minus_sign:                                                                   | The defaultValue field.                                                              |
| `Rules`                                                                              | [*shared.StringMapRules](../../../pkg/models/shared/stringmaprules.md)               | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Secret`                                                                             | `*bool`                                                                              | :heavy_minus_sign:                                                                   | If secret, the map's values are write-only and rendered masked. Keys stay<br/> readable. |