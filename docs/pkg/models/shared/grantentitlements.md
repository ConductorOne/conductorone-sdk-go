# GrantEntitlements

The GrantEntitlements message.

This message contains a oneof named inclusion. Only a single field of the following list may be set at a time:
  - inclusionList
  - inclusionCriteria
  - inclusionListCel


This message contains a oneof named exclusion. Only a single field of the following list may be set at a time:
  - exclusionNone
  - exclusionList
  - exclusionCriteria
  - exclusionListCel



## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `GrantEntitlementExclusionCriteria`                                                                          | [*shared.GrantEntitlementExclusionCriteria](../../../pkg/models/shared/grantentitlementexclusioncriteria.md) | :heavy_minus_sign:                                                                                           | The GrantEntitlementExclusionCriteria message.                                                               |
| `GrantEntitlementExclusionList`                                                                              | [*shared.GrantEntitlementExclusionList](../../../pkg/models/shared/grantentitlementexclusionlist.md)         | :heavy_minus_sign:                                                                                           | The GrantEntitlementExclusionList message.                                                                   |
| `GrantEntitlementExclusionListCel`                                                                           | [*shared.GrantEntitlementExclusionListCel](../../../pkg/models/shared/grantentitlementexclusionlistcel.md)   | :heavy_minus_sign:                                                                                           | The GrantEntitlementExclusionListCel message.                                                                |
| `GrantEntitlementExclusionNone`                                                                              | [*shared.GrantEntitlementExclusionNone](../../../pkg/models/shared/grantentitlementexclusionnone.md)         | :heavy_minus_sign:                                                                                           | The GrantEntitlementExclusionNone message.                                                                   |
| `GrantEntitlementInclusionCriteria`                                                                          | [*shared.GrantEntitlementInclusionCriteria](../../../pkg/models/shared/grantentitlementinclusioncriteria.md) | :heavy_minus_sign:                                                                                           | The GrantEntitlementInclusionCriteria message.                                                               |
| `GrantEntitlementInclusionList`                                                                              | [*shared.GrantEntitlementInclusionList](../../../pkg/models/shared/grantentitlementinclusionlist.md)         | :heavy_minus_sign:                                                                                           | The GrantEntitlementInclusionList message.                                                                   |
| `GrantEntitlementInclusionListCel`                                                                           | [*shared.GrantEntitlementInclusionListCel](../../../pkg/models/shared/grantentitlementinclusionlistcel.md)   | :heavy_minus_sign:                                                                                           | The GrantEntitlementInclusionListCel message.                                                                |
| `UserRef`                                                                                                    | [*shared.UserRef](../../../pkg/models/shared/userref.md)                                                     | :heavy_minus_sign:                                                                                           | A reference to a user.                                                                                       |
| `UseSubjectUser`                                                                                             | **bool*                                                                                                      | :heavy_minus_sign:                                                                                           | If true, the step will use the subject user of the automation as the subject.                                |
| `UserIDCel`                                                                                                  | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | The userIdCel field.                                                                                         |