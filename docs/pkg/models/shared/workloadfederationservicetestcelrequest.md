# WorkloadFederationServiceTestCELRequest

The WorkloadFederationServiceTestCELRequest message.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ClaimsJSON`                                                                                      | **string*                                                                                         | :heavy_minus_sign:                                                                                | The claims to evaluate against, as a JSON string.<br/> Parsed into map[string]any for CEL evaluation. |
| `Expression`                                                                                      | **string*                                                                                         | :heavy_minus_sign:                                                                                | The CEL expression to evaluate. Must return bool.                                                 |