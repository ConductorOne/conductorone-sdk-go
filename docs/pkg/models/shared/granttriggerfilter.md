# GrantTriggerFilter

The GrantTriggerFilter message.

This message contains a oneof named entitlement_inclusion. Only a single field of the following list may be set at a time:
  - inclusionList
  - inclusionAll
  - inclusionCriteria
  - inclusionListCel



## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `AccountFilter`                                                                                    | [*shared.AccountFilter](../../../pkg/models/shared/accountfilter.md)                               | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `GrantFilter`                                                                                      | [*shared.GrantFilter](../../../pkg/models/shared/grantfilter.md)                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `InclusionAll`                                                                                     | [*shared.EntitlementInclusionAll](../../../pkg/models/shared/entitlementinclusionall.md)           | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `InclusionCriteria`                                                                                | [*shared.EntitlementInclusionCriteria](../../../pkg/models/shared/entitlementinclusioncriteria.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `InclusionList`                                                                                    | [*shared.EntitlementInclusionList](../../../pkg/models/shared/entitlementinclusionlist.md)         | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `InclusionListCel`                                                                                 | [*shared.EntitlementInclusionListCel](../../../pkg/models/shared/entitlementinclusionlistcel.md)   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |