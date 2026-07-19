# DynamicNumber

DynamicNumber can be a literal value, a JSON pointer path, or a function call.

This message contains a oneof named value. Only a single field of the following list may be set at a time:
  - literal
  - path
  - call



## Fields

| Field                                                                                                                                  | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `Call`                                                                                                                                 | [*shared.FunctionCall](../../../pkg/models/shared/functioncall.md)                                                                     | :heavy_minus_sign:                                                                                                                     | N/A                                                                                                                                    |
| `Literal`                                                                                                                              | `*float64`                                                                                                                             | :heavy_minus_sign:                                                                                                                     | The literal field.<br/>This field is part of the `value` oneof.<br/>See the documentation for `c1.api.a2ui.v1.DynamicNumber` for more details. |
| `Path`                                                                                                                                 | `*string`                                                                                                                              | :heavy_minus_sign:                                                                                                                     | The path field.<br/>This field is part of the `value` oneof.<br/>See the documentation for `c1.api.a2ui.v1.DynamicNumber` for more details. |