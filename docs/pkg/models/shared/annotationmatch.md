# AnnotationMatch

Matches entitlements whose annotations contain `key` with exactly `value`.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Key`                                                                                       | `*string`                                                                                   | :heavy_minus_sign:                                                                          | The key field.                                                                              |
| `Value`                                                                                     | `*string`                                                                                   | :heavy_minus_sign:                                                                          | The full value pattern is enforced server-side by pkg/annotations where criteria are saved. |