# SearchDenialsRequest

SearchDenialsRequest specifies denial filters and pagination.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Filters`                                                                       | [*shared.DenialFilters](../../../pkg/models/shared/denialfilters.md)            | :heavy_minus_sign:                                                              | N/A                                                                             |
| `PageSize`                                                                      | `*int`                                                                          | :heavy_minus_sign:                                                              | Maximum number of episodes to return. The default is 25 and the maximum is 100. |
| `PageToken`                                                                     | `*string`                                                                       | :heavy_minus_sign:                                                              | Pagination token from a previous response with the same filters.                |