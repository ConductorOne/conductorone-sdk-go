# PersonalClientSearchServiceSearchRequest

The PersonalClientSearchServiceSearchRequest message.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `PageSize`                                                       | `*int`                                                           | :heavy_minus_sign:                                               | The maximum number of results to return per page.                |
| `PageToken`                                                      | `*string`                                                        | :heavy_minus_sign:                                               | A pagination token returned from a previous Search call.         |
| `Query`                                                          | `*string`                                                        | :heavy_minus_sign:                                               | A text query to filter personal clients by display name.         |
| `Users`                                                          | [][shared.UserRef](../../../pkg/models/shared/userref.md)        | :heavy_minus_sign:                                               | Filter results to personal clients owned by the specified users. |