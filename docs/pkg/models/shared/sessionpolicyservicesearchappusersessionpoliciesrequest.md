# SessionPolicyServiceSearchAppUserSessionPoliciesRequest

Search the policies selected for users with accounts holding the application's
 SSO sign-in entitlement. Accounts not linked to a current user are excluded.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `PageSize`                                                                          | `*int`                                                                              | :heavy_minus_sign:                                                                  | Maximum number of policies per page. Follow next_page_token to complete the search. |
| `PageToken`                                                                         | `*string`                                                                           | :heavy_minus_sign:                                                                  | The page token from the previous response.                                          |