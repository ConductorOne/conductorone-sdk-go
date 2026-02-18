# ExpressionProvisioner

ExpressionProvisioner evaluates CEL expressions to determine provisioners.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AllowReassignment`                                         | **bool*                                                     | :heavy_minus_sign:                                          | Whether the provisioner can reassign the task.              |
| `Expressions`                                               | []*string*                                                  | :heavy_minus_sign:                                          | The CEL expressions to evaluate.                            |
| `FallbackUserIds`                                           | []*string*                                                  | :heavy_minus_sign:                                          | Fallback user IDs if expression evaluation yields no users. |