# C1APIPolicyV1ProvisionPolicy

The ProvisionPolicy message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Connector`                                                                                | [*C1APIPolicyV1ConnectorProvision](../../models/shared/c1apipolicyv1connectorprovision.md) | :heavy_minus_sign:                                                                         | The ConnectorProvision message.                                                            |
| `Delegated`                                                                                | [*C1APIPolicyV1DelegatedProvision](../../models/shared/c1apipolicyv1delegatedprovision.md) | :heavy_minus_sign:                                                                         | The DelegatedProvision message.                                                            |
| `Manual`                                                                                   | [*C1APIPolicyV1ManualProvision](../../models/shared/c1apipolicyv1manualprovision.md)       | :heavy_minus_sign:                                                                         | The ManualProvision message.                                                               |