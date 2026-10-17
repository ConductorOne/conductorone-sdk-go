# ModelAPIKeyServiceListResponse

The ModelAPIKeyServiceListResponse message.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Keys`                                                            | [][shared.ModelAPIKey](../../../pkg/models/shared/modelapikey.md) | :heavy_minus_sign:                                                | The requested model API keys (metadata only, no secrets).         |
| `NextPageToken`                                                   | `*string`                                                         | :heavy_minus_sign:                                                | Pagination token for the next page of results.                    |