# SessionPolicyServiceListUserPoliciesResponse

SessionPolicyServiceListUserPoliciesResponse carries every policy that
 applies to the user. Unpaginated: the candidate set is the user's assigned
 policies plus at most one tenant default.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Policies`                                                                        | [][shared.EffectiveUserPolicy](../../../pkg/models/shared/effectiveuserpolicy.md) | :heavy_minus_sign:                                                                | Every policy that applies to the user, in assignment order.                       |