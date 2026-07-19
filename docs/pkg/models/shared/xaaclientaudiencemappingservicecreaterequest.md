# XAAClientAudienceMappingServiceCreateRequest

XAAClientAudienceMappingServiceCreateRequest creates a new mapping.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AudienceClientID`                                                    | `*string`                                                             | :heavy_minus_sign:                                                    | The client's identifier at the resource authorization server.         |
| `ClientKey`                                                           | `*string`                                                             | :heavy_minus_sign:                                                    | Stable client registration key.                                       |
| `Disabled`                                                            | `*bool`                                                               | :heavy_minus_sign:                                                    | When true, the mapping is created but exchange requests are rejected. |