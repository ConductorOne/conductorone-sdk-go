# SSFOutboundAuthOAuth2

SSFOutboundAuthOAuth2 uses OAuth2 client credentials for outbound auth.
 client_secret is write-only: accepted on create/update, never returned.


## Fields

| Field                   | Type                    | Required                | Description             |
| ----------------------- | ----------------------- | ----------------------- | ----------------------- |
| `ClientID`              | `*string`               | :heavy_minus_sign:      | The clientId field.     |
| `ClientSecret`          | `*string`               | :heavy_minus_sign:      | The clientSecret field. |
| `Scopes`                | []`string`              | :heavy_minus_sign:      | The scopes field.       |
| `TokenURL`              | `*string`               | :heavy_minus_sign:      | The tokenUrl field.     |