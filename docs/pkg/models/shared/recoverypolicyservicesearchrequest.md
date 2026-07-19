# RecoveryPolicyServiceSearchRequest

The RecoveryPolicyServiceSearchRequest message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `PageSize`                                                                    | `*int`                                                                        | :heavy_minus_sign:                                                            | The maximum number of results to return per page.                             |
| `PageToken`                                                                   | `*string`                                                                     | :heavy_minus_sign:                                                            | A pagination token from a previous Search response.                           |
| `Query`                                                                       | `*string`                                                                     | :heavy_minus_sign:                                                            | Free-text search over the policy name. Empty matches all policies.            |
| `Refs`                                                                        | [][shared.RecoveryPolicyRef](../../../pkg/models/shared/recoverypolicyref.md) | :heavy_minus_sign:                                                            | Restrict results to these specific policies. Empty matches all policies.      |