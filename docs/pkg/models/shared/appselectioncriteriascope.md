# AppSelectionCriteriaScope

Selects entitlements matching every set field; values within a field are alternatives.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AppIds`                                                                  | []`string`                                                                | :heavy_minus_sign:                                                        | The appIds field.                                                         |
| `ComplianceFrameworkAttributeValueIds`                                    | []`string`                                                                | :heavy_minus_sign:                                                        | The complianceFrameworkAttributeValueIds field.                           |
| `ResourceTypeRefs`                                                        | [][shared.ResourceTypeRef](../../../pkg/models/shared/resourcetyperef.md) | :heavy_minus_sign:                                                        | The resourceTypeRefs field.                                               |
| `RiskLevelAttributeValueIds`                                              | []`string`                                                                | :heavy_minus_sign:                                                        | The riskLevelAttributeValueIds field.                                     |