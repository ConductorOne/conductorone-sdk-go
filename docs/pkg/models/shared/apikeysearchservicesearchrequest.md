# APIKeySearchServiceSearchRequest

The APIKeySearchServiceSearchRequest message.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `PageSize`                                                    | `*int`                                                        | :heavy_minus_sign:                                            | The maximum number of results to return per page.             |
| `PageToken`                                                   | `*string`                                                     | :heavy_minus_sign:                                            | A pagination token returned from a previous Search call.      |
| `Query`                                                       | `*string`                                                     | :heavy_minus_sign:                                            | A text query to filter API keys by display name.              |
| `Types`                                                       | [][shared.Types](../../../pkg/models/shared/types.md)         | :heavy_minus_sign:                                            | The key types to return. When empty, both types are returned. |