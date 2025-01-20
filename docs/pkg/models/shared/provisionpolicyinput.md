# ProvisionPolicyInput

ProvisionPolicy is a oneOf that indicates how a provision step should be processed.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated
  - webhook
  - multiStep
  - externalTicket



## Fields

| Field                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ConnectorProvision`                                                                                                                                                                                                                                         | [*shared.ConnectorProvision](../../../pkg/models/shared/connectorprovision.md)                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                           | Indicates that a connector should perform the provisioning. This object has no fields.                                                                                                                                                                       |
| `DelegatedProvision`                                                                                                                                                                                                                                         | [*shared.DelegatedProvision](../../../pkg/models/shared/delegatedprovision.md)                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                           | This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement. |
| `ExternalTicketProvision`                                                                                                                                                                                                                                    | [*shared.ExternalTicketProvision](../../../pkg/models/shared/externalticketprovision.md)                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                           | This provision step indicates that we should check an external ticket to provision this entitlement                                                                                                                                                          |
| `ManualProvision`                                                                                                                                                                                                                                            | [*shared.ManualProvision](../../../pkg/models/shared/manualprovision.md)                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                           | Manual provisioning indicates that a human must intervene for the provisioning of this step.                                                                                                                                                                 |
| `MultiStep`                                                                                                                                                                                                                                                  | [*shared.MultiStep](../../../pkg/models/shared/multistep.md)                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                           | MultiStep indicates that this provision step has multiple steps to process.                                                                                                                                                                                  |
| `WebhookProvision`                                                                                                                                                                                                                                           | [*shared.WebhookProvision](../../../pkg/models/shared/webhookprovision.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                           | This provision step indicates that a webhook should be called to provision this entitlement.                                                                                                                                                                 |