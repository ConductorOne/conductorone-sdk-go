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
| `ExclusionCriteria`                                                                                          | [*shared.GrantEntitlementExclusionCriteria](../../../pkg/models/shared/grantentitlementexclusioncriteria.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ExclusionList`                                                                                              | [*shared.GrantEntitlementExclusionList](../../../pkg/models/shared/grantentitlementexclusionlist.md)         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ExclusionListCel`                                                                                           | [*shared.GrantEntitlementExclusionListCel](../../../pkg/models/shared/grantentitlementexclusionlistcel.md)   | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `ExclusionNone`                                                                                              | [*shared.GrantEntitlementExclusionNone](../../../pkg/models/shared/grantentitlementexclusionnone.md)         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `InclusionCriteria`                                                                                          | [*shared.GrantEntitlementInclusionCriteria](../../../pkg/models/shared/grantentitlementinclusioncriteria.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `InclusionList`                                                                                              | [*shared.GrantEntitlementInclusionList](../../../pkg/models/shared/grantentitlementinclusionlist.md)         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `InclusionListCel`                                                                                           | [*shared.GrantEntitlementInclusionListCel](../../../pkg/models/shared/grantentitlementinclusionlistcel.md)   | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `UseSubjectUser`                                                                                             | `*bool`                                                                                                      | :heavy_minus_sign:                                                                                           | If true, the step will use the subject user of the automation as the subject.                                |
| `UserIDCel`                                                                                                  | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | The userIdCel field.                                                                                         |
| `UserRef`                                                                                                    | [*shared.UserRef](../../../pkg/models/shared/userref.md)                                                     | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |