# SessionPolicyServiceSearchRequest

The SessionPolicyServiceSearchRequest message.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `PageSize`                                                                  | `*int`                                                                      | :heavy_minus_sign:                                                          | The pageSize field.                                                         |
| `PageToken`                                                                 | `*string`                                                                   | :heavy_minus_sign:                                                          | The pageToken field.                                                        |
| `Query`                                                                     | `*string`                                                                   | :heavy_minus_sign:                                                          | Free-text search over the policy name. Empty matches all policies.          |
| `Refs`                                                                      | [][shared.SessionPolicyRef](../../../pkg/models/shared/sessionpolicyref.md) | :heavy_minus_sign:                                                          | Restrict results to these specific policies. Empty matches all policies.    |