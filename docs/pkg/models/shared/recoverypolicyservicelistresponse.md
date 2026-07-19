# RecoveryPolicyServiceListResponse

The RecoveryPolicyServiceListResponse message.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `List`                                                                  | [][shared.RecoveryPolicy](../../../pkg/models/shared/recoverypolicy.md) | :heavy_minus_sign:                                                      | The page of policies.                                                   |
| `NextPageToken`                                                         | `*string`                                                               | :heavy_minus_sign:                                                      | A token to fetch the next page, or empty if there are no more results.  |