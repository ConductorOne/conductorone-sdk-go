# WebhooksSearchResponse

The WebhooksSearchResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.WebhookEndpoint](../../../pkg/models/shared/webhookendpoint.md)            | :heavy_minus_sign:                                                                   | The list of webhooks matching the search criteria.                                   |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |