# EdgeEgressCheck

EdgeEgressCheck requires an integration's authorization in addition to egress policy.

This message contains a oneof named authorizer. Only a single field of the following list may be set at a time:
  - integration



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Enabled`                                                                              | `*bool`                                                                                | :heavy_minus_sign:                                                                     | Whether this check is enforced. Disabled checks still reference the integration.       |
| `ID`                                                                                   | `*string`                                                                              | :heavy_minus_sign:                                                                     | Stable identifier unique within the Edge.                                              |
| `Integration`                                                                          | [*shared.EdgeIntegrationRef](../../../pkg/models/shared/edgeintegrationref.md)         | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Match`                                                                                | [*shared.EdgeEgressRequestMatch](../../../pkg/models/shared/edgeegressrequestmatch.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |