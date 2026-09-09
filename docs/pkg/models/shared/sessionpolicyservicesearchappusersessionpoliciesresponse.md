# SessionPolicyServiceSearchAppUserSessionPoliciesResponse

Distinct effective policies, ordered by display name and policy ID.
 A tenant default appears only if it is selected for at least one user.
 An empty list means no effective policies were found, including when no
 current users have accounts holding the sign-in entitlement.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `List`                                                                | [][shared.SessionPolicy](../../../pkg/models/shared/sessionpolicy.md) | :heavy_minus_sign:                                                    | The list field.                                                       |
| `NextPageToken`                                                       | `*string`                                                             | :heavy_minus_sign:                                                    | The token for the next page. Empty on the last page.                  |