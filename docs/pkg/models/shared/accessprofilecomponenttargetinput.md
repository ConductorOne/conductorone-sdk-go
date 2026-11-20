# AccessProfileComponentTargetInput

One component of an access profile's definition that an action task changes.

This message contains a oneof named component. Only a single field of the following list may be set at a time:
  - compositionEntitlement
  - membershipRule
  - membershipEntitlement



## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `AccessProfile`                                                                                            | [*shared.AccessProfileRef](../../../pkg/models/shared/accessprofileref.md)                                 | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `CompositionEntitlement`                                                                                   | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)                               | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `MembershipEntitlement`                                                                                    | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)                               | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `MembershipRule`                                                                                           | [*shared.AccessProfileMembershipRuleInput](../../../pkg/models/shared/accessprofilemembershipruleinput.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |