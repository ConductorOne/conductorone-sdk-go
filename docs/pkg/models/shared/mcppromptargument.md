# MCPPromptArgument

MCPPromptArgument describes a native prompt argument, rather than an argument value.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Description`                                         | `*string`                                             | :heavy_minus_sign:                                    | Argument description, supplied by discovery.          |
| `Name`                                                | `*string`                                             | :heavy_minus_sign:                                    | Native argument name, supplied by discovery.          |
| `Required`                                            | `*bool`                                               | :heavy_minus_sign:                                    | Whether the native prompt requires this argument.     |
| `Title`                                               | `*string`                                             | :heavy_minus_sign:                                    | Human-readable argument title, supplied by discovery. |