# SSOApplicationServiceSearchRequest

SSOApplicationServiceSearchRequest searches SSO applications with filters.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppIds`                                                                         | []`string`                                                                       | :heavy_minus_sign:                                                               | Optional filter by applications in your catalog. Empty matches any<br/> application. |
| `PageSize`                                                                       | `*int`                                                                           | :heavy_minus_sign:                                                               | Maximum number of results to return per page.                                    |
| `PageToken`                                                                      | `*string`                                                                        | :heavy_minus_sign:                                                               | Pagination token from a previous response.                                       |
| `Query`                                                                          | `*string`                                                                        | :heavy_minus_sign:                                                               | Optional text query matched against display_name and description.                |