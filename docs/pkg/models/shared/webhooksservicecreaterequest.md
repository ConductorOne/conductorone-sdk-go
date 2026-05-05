# WebhooksServiceCreateRequest

The WebhooksServiceCreateRequest message.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CallbackTimeout`                                                        | `*string`                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | An optional description of the webhook's purpose.                        |
| `DisplayName`                                                            | `string`                                                                 | :heavy_check_mark:                                                       | The human-readable name for the new webhook.                             |
| `URL`                                                                    | `string`                                                                 | :heavy_check_mark:                                                       | The destination URL that will receive event notification HTTP callbacks. |