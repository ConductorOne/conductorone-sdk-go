# EffectiveUserPolicy

EffectiveUserPolicy is one session policy that applies to a user.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Group`                                                                                      | [*shared.AppEntitlement](../../../pkg/models/shared/appentitlement.md)                       | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Policy`                                                                                     | [*shared.SessionPolicy](../../../pkg/models/shared/sessionpolicy.md)                         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Source`                                                                                     | [*shared.EffectiveUserPolicySource](../../../pkg/models/shared/effectiveuserpolicysource.md) | :heavy_minus_sign:                                                                           | Why the policy applies to the user.                                                          |