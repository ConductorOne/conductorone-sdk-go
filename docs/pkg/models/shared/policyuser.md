# PolicyUser

PolicyUser is one user a session policy applies to.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Group`                                                                    | [*shared.AppEntitlement](../../../pkg/models/shared/appentitlement.md)     | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Source`                                                                   | [*shared.PolicyUserSource](../../../pkg/models/shared/policyusersource.md) | :heavy_minus_sign:                                                         | Why the policy applies to this user. DIRECT or GROUP.                      |
| `User`                                                                     | [*shared.User](../../../pkg/models/shared/user.md)                         | :heavy_minus_sign:                                                         | N/A                                                                        |