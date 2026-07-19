# DecoyVendingMaterial

DecoyVendingMaterial carries the freshly-vended secret material returned
 exactly once at Create or Rotate.

This message contains a oneof named material. Only a single field of the following list may be set at a time:
  - clientCredential
  - accessToken
  - workloadFederation



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `AccessToken`                                                                                            | [*shared.DecoyAccessTokenMaterial](../../../pkg/models/shared/decoyaccesstokenmaterial.md)               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ClientCredential`                                                                                       | [*shared.DecoyClientCredentialMaterial](../../../pkg/models/shared/decoyclientcredentialmaterial.md)     | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `WorkloadFederation`                                                                                     | [*shared.DecoyWorkloadFederationMaterial](../../../pkg/models/shared/decoyworkloadfederationmaterial.md) | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |