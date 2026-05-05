# WebhooksSearchRequest

The WebhooksSearchRequest message.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `PageSize`                                                                      | `*int`                                                                          | :heavy_minus_sign:                                                              | The maximum number of webhooks to return per page.                              |
| `PageToken`                                                                     | `*string`                                                                       | :heavy_minus_sign:                                                              | The pagination token from a previous search response to fetch the next page.    |
| `Query`                                                                         | `*string`                                                                       | :heavy_minus_sign:                                                              | A text query to match against webhook names and descriptions.                   |
| `Refs`                                                                          | [][shared.WebhookRef](../../../pkg/models/shared/webhookref.md)                 | :heavy_minus_sign:                                                              | Optional set of webhook references to restrict the search to specific webhooks. |