# WebhookDispatcher

WebhookDispatcher POSTs to a registered webhook (webhooks v3).


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `PayloadTemplate`                                                  | `*string`                                                          | :heavy_minus_sign:                                                 | Optional payload template; empty uses the default finding payload. |
| `WebhookID`                                                        | `*string`                                                          | :heavy_minus_sign:                                                 | ID of a registered webhook to POST to.                             |