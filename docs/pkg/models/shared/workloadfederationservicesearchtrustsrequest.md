# WorkloadFederationServiceSearchTrustsRequest

The WorkloadFederationServiceSearchTrustsRequest message.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `PageSize`                                                        | **int*                                                            | :heavy_minus_sign:                                                | The pageSize field.                                               |
| `PageToken`                                                       | **string*                                                         | :heavy_minus_sign:                                                | The pageToken field.                                              |
| `ProviderID`                                                      | **string*                                                         | :heavy_minus_sign:                                                | Optional: filter trusts by provider ID.                           |
| `Query`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Optional: full-text search on trust display name and description. |
| `ServicePrincipalID`                                              | **string*                                                         | :heavy_minus_sign:                                                | Optional: filter trusts by service principal ID.                  |