# ProvisionPolicy

ProvisionPolicy is a oneOf that indicates how a provision step should be processed.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated
  - webhook
  - multiStep
  - externalTicket
  - unconfigured
  - action
  - devicePlacement



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Action`                                                                                   | [*shared.ActionProvision](../../../pkg/models/shared/actionprovision.md)                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Connector`                                                                                | [*shared.ConnectorProvision](../../../pkg/models/shared/connectorprovision.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Delegated`                                                                                | [*shared.DelegatedProvision](../../../pkg/models/shared/delegatedprovision.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `DevicePlacement`                                                                          | [*shared.DevicePlacementProvision](../../../pkg/models/shared/deviceplacementprovision.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ExternalTicket`                                                                           | [*shared.ExternalTicketProvision](../../../pkg/models/shared/externalticketprovision.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Manual`                                                                                   | [*shared.ManualProvision](../../../pkg/models/shared/manualprovision.md)                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `MultiStep`                                                                                | [*shared.MultiStep](../../../pkg/models/shared/multistep.md)                               | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Unconfigured`                                                                             | [*shared.UnconfiguredProvision](../../../pkg/models/shared/unconfiguredprovision.md)       | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Webhook`                                                                                  | [*shared.WebhookProvision](../../../pkg/models/shared/webhookprovision.md)                 | :heavy_minus_sign:                                                                         | N/A                                                                                        |