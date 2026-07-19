# AccessReviewTemplateSetScopeByResourceTypeRequest

The AccessReviewTemplateSetScopeByResourceTypeRequest message.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ResourceTypeSelections`                                                                              | [][shared.ResourceTypeIDRef](../../../pkg/models/shared/resourcetypeidref.md)                         | :heavy_minus_sign:                                                                                    | The resource types to include in the template scope. Replaces all previously selected resource types. |
| `Scope`                                                                                               | [*shared.AccessReviewScopeV2](../../../pkg/models/shared/accessreviewscopev2.md)                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |