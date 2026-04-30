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
| `AccountFilter`                                                                                    | [*shared.AccountFilter](../../../pkg/models/shared/accountfilter.md)                               | :heavy_minus_sign:                                                                                 | The AccountFilter message.                                                                         |
| `EntitlementInclusionAll`                                                                          | [*shared.EntitlementInclusionAll](../../../pkg/models/shared/entitlementinclusionall.md)           | :heavy_minus_sign:                                                                                 | The EntitlementInclusionAll message.                                                               |
| `EntitlementInclusionCriteria`                                                                     | [*shared.EntitlementInclusionCriteria](../../../pkg/models/shared/entitlementinclusioncriteria.md) | :heavy_minus_sign:                                                                                 | The EntitlementInclusionCriteria message.                                                          |
| `EntitlementInclusionList`                                                                         | [*shared.EntitlementInclusionList](../../../pkg/models/shared/entitlementinclusionlist.md)         | :heavy_minus_sign:                                                                                 | The EntitlementInclusionList message.                                                              |
| `EntitlementInclusionListCel`                                                                      | [*shared.EntitlementInclusionListCel](../../../pkg/models/shared/entitlementinclusionlistcel.md)   | :heavy_minus_sign:                                                                                 | The EntitlementInclusionListCel message.                                                           |
| `GrantFilter`                                                                                      | [*shared.GrantFilter](../../../pkg/models/shared/grantfilter.md)                                   | :heavy_minus_sign:                                                                                 | The GrantFilter message.                                                                           |