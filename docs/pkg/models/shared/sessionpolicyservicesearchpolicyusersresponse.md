# SessionPolicyServiceSearchPolicyUsersResponse

SessionPolicyServiceSearchPolicyUsersResponse carries one page of the users
 a policy applies to.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `NextPageToken`                                                 | `*string`                                                       | :heavy_minus_sign:                                              | The token for the next page. Empty when this is the last page.  |
| `Users`                                                         | [][shared.PolicyUser](../../../pkg/models/shared/policyuser.md) | :heavy_minus_sign:                                              | The users on this page.                                         |