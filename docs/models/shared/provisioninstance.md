# ProvisionInstance

The ProvisionInstance message.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - completed
  - cancelled
  - errored
  - reassignedByError



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `NotificationID`                                                             | **string*                                                                    | :heavy_minus_sign:                                                           | The notificationId field.                                                    |
| `Outcome`                                                                    | [*ProvisionInstanceOutcome](../../models/shared/provisioninstanceoutcome.md) | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Provision`                                                                  | [*Provision](../../models/shared/provision.md)                               | :heavy_minus_sign:                                                           | The Provision message.                                                       |
| `State`                                                                      | [*ProvisionInstanceState](../../models/shared/provisioninstancestate.md)     | :heavy_minus_sign:                                                           | The state field.                                                             |