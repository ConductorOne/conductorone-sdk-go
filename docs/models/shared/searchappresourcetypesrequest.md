# SearchAppResourceTypesRequest

 Search for app resources based on some filters.



## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AppIds`                                                        | []*string*                                                      | :heavy_minus_sign:                                              |  A list of app IDs to restrict the search by.<br/>              |
| `ExcludeResourceTypeIds`                                        | []*string*                                                      | :heavy_minus_sign:                                              |  A list of resource type IDs to exclude from the search.<br/>   |
| `ExcludeResourceTypeTraitIds`                                   | []*string*                                                      | :heavy_minus_sign:                                              |  A list of resource type trait IDs to exclude from the search.<br/> |
| `PageSize`                                                      | **float64*                                                      | :heavy_minus_sign:                                              |  The pageSize where 10 <= pageSize <= 100, default 25.<br/>     |
| `PageToken`                                                     | **string*                                                       | :heavy_minus_sign:                                              |  The pageToken field.<br/>                                      |
| `Query`                                                         | **string*                                                       | :heavy_minus_sign:                                              |  Fuzzy search the display name of resource types.<br/>          |
| `ResourceTypeIds`                                               | []*string*                                                      | :heavy_minus_sign:                                              |  A list of resource type IDs to restrict the search by.<br/>    |
| `ResourceTypeTraitIds`                                          | []*string*                                                      | :heavy_minus_sign:                                              |  A list of resource type trait IDs to restrict the search by.<br/> |