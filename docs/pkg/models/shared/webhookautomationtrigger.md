# WebhookAutomationTrigger

The WebhookAutomationTrigger message.

This message contains a oneof named auth_config. Only a single field of the following list may be set at a time:
  - jwt
  - hmac
  - capabilityUrl



## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `CapabilityURL`                                                                                            | [*shared.WebhookListenerAuthCapabilityURL](../../../pkg/models/shared/webhooklistenerauthcapabilityurl.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `Hmac`                                                                                                     | [*shared.WebhookListenerAuthHMAC](../../../pkg/models/shared/webhooklistenerauthhmac.md)                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `Jwt`                                                                                                      | [*shared.WebhookListenerAuthJWT](../../../pkg/models/shared/webhooklistenerauthjwt.md)                     | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `ListenerID`                                                                                               | `*string`                                                                                                  | :heavy_minus_sign:                                                                                         | Optional existing listener ID (hidden field from frontend)                                                 |