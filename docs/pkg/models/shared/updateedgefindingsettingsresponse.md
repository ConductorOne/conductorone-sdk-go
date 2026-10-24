# UpdateEdgeFindingSettingsResponse

The UpdateEdgeFindingSettingsResponse message.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `List`                                                                                                | [][shared.EdgeFindingCategorySetting](../../../pkg/models/shared/edgefindingcategorysetting.md)       | :heavy_minus_sign:                                                                                    | Every category after the write, in the same shape GetEdgeFindingSettings<br/> returns.                |
| `Subcategories`                                                                                       | [][shared.EdgeFindingSubcategorySetting](../../../pkg/models/shared/edgefindingsubcategorysetting.md) | :heavy_minus_sign:                                                                                    | Every AI usage subcategory after the write.                                                           |