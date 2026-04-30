# HookFunctionRef

HookFunctionRef identifies a customer-authored function to invoke.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CommitID`                                                            | `*string`                                                             | :heavy_minus_sign:                                                    | If empty, the function's published commit is used at invocation time. |
| `FunctionID`                                                          | `*string`                                                             | :heavy_minus_sign:                                                    | The functionId field.                                                 |