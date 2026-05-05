# WebhookEndpoint

The Webhook message.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CallbackTimeout`                                                    | `*string`                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DeletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Description`                                                        | `*string`                                                            | :heavy_minus_sign:                                                   | An optional description of the webhook's purpose.                    |
| `DisplayName`                                                        | `*string`                                                            | :heavy_minus_sign:                                                   | The human-readable name of the webhook.                              |
| `ID`                                                                 | `*string`                                                            | :heavy_minus_sign:                                                   | The unique identifier of the webhook.                                |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `URL`                                                                | `*string`                                                            | :heavy_minus_sign:                                                   | The destination URL that receives event notification HTTP callbacks. |