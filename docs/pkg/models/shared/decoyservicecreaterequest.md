# DecoyServiceCreateRequest

The DecoyServiceCreateRequest message.

This message contains a oneof named create_input. Only a single field of the following list may be set at a time:
  - userClientCredential
  - connectorClient
  - workloadFed
  - accessToken



## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `AccessToken`                                                                                          | [*shared.DecoyAccessTokenInput](../../../pkg/models/shared/decoyaccesstokeninput.md)                   | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `Annotations`                                                                                          | map[string]`string`                                                                                    | :heavy_minus_sign:                                                                                     | The annotations field.                                                                                 |
| `ConnectorClient`                                                                                      | [*shared.DecoyConnectorClientInput](../../../pkg/models/shared/decoyconnectorclientinput.md)           | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `Description`                                                                                          | `*string`                                                                                              | :heavy_minus_sign:                                                                                     | The description field.                                                                                 |
| `DisplayName`                                                                                          | `*string`                                                                                              | :heavy_minus_sign:                                                                                     | The displayName field.                                                                                 |
| `UserClientCredential`                                                                                 | [*shared.DecoyUserClientCredentialInput](../../../pkg/models/shared/decoyuserclientcredentialinput.md) | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `WorkloadFed`                                                                                          | [*shared.DecoyWorkloadFederationInput](../../../pkg/models/shared/decoyworkloadfederationinput.md)     | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |