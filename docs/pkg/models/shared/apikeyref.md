# APIKeyRef

APIKeyRef identifies an API key, or another credential that makes API requests.

This message contains a oneof named credential. Only a single field of the following list may be set at a time:
  - personalClient
  - deviceClient
  - servicePrincipalCredential
  - connectedAppClientGrant
  - connectedAppOauthGrant
  - systemClient



## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ConnectedAppClientGrant`                                                                            | [*shared.ConnectedAppClientGrantRef](../../../pkg/models/shared/connectedappclientgrantref.md)       | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `ConnectedAppOauthGrant`                                                                             | [*shared.ConnectedAppOAuthGrantRef](../../../pkg/models/shared/connectedappoauthgrantref.md)         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `DeviceClient`                                                                                       | [*shared.DeviceClientRef](../../../pkg/models/shared/deviceclientref.md)                             | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `PersonalClient`                                                                                     | [*shared.PersonalClientRef](../../../pkg/models/shared/personalclientref.md)                         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `ServicePrincipalCredential`                                                                         | [*shared.ServicePrincipalCredentialRef](../../../pkg/models/shared/serviceprincipalcredentialref.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `SystemClient`                                                                                       | [*shared.SystemClientRef](../../../pkg/models/shared/systemclientref.md)                             | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |