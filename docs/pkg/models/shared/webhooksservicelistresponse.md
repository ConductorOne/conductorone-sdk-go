# WebhooksServiceListResponse

The WebhooksServiceListResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.WebhookEndpoint](../../../pkg/models/shared/webhookendpoint.md)            | :heavy_minus_sign:                                                                   | The list of webhooks for the current page.                                           |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |