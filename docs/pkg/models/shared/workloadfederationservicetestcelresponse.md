# WorkloadFederationServiceTestCELResponse

The WorkloadFederationServiceTestCELResponse message.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Error`                                            | **string*                                          | :heavy_minus_sign:                                 | Error message if compilation or evaluation failed. |
| `Expression`                                       | **string*                                          | :heavy_minus_sign:                                 | The expression that was evaluated (echo back).     |
| `Matched`                                          | **bool*                                            | :heavy_minus_sign:                                 | Whether the expression matched (returned true).    |