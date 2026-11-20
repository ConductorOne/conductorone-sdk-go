# APIKey

APIKey is one API key search result: a personal client or a service principal key.

This message contains a oneof named key. Only a single field of the following list may be set at a time:
  - personalClient
  - servicePrincipalKey



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `PersonalClient`                                                                       | [*shared.PersonalClient](../../../pkg/models/shared/personalclient.md)                 | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `ServicePrincipalKey`                                                                  | [*shared.ServicePrincipalAPIKey](../../../pkg/models/shared/serviceprincipalapikey.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |