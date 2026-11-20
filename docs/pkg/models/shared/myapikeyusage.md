# MyAPIKeyUsage

MyAPIKeyUsage is the live usage of one of the caller's API keys.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `APIKey`                                                                | [*shared.APIKeyIdentity](../../../pkg/models/shared/apikeyidentity.md)  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Usage`                                                                 | [][shared.RateLimitUsage](../../../pkg/models/shared/ratelimitusage.md) | :heavy_minus_sign:                                                      | Live usage of the key's own limits.                                     |