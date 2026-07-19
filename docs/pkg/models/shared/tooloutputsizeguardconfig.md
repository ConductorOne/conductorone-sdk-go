# ToolOutputSizeGuardConfig

ToolOutputSizeGuardConfig caps post-tool-use output size in bytes.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `MaxBytes`                                                            | `*int`                                                                | :heavy_minus_sign:                                                    | Maximum tool output size in bytes. Outputs exceeding this are denied. |