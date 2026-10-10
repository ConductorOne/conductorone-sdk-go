# GetEdgeFindingSettingsResponse

The GetEdgeFindingSettingsResponse message.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `List`                                                                                                | [][shared.EdgeFindingCategorySetting](../../../pkg/models/shared/edgefindingcategorysetting.md)       | :heavy_minus_sign:                                                                                    | One entry per Edge finding category, in EdgeFindingCategory declaration<br/> order.                   |
| `Subcategories`                                                                                       | [][shared.EdgeFindingSubcategorySetting](../../../pkg/models/shared/edgefindingsubcategorysetting.md) | :heavy_minus_sign:                                                                                    | One entry per AI usage subcategory, in EdgeFindingSubcategory<br/> declaration order.                 |