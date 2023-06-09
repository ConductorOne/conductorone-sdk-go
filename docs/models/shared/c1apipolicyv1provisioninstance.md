# C1APIPolicyV1ProvisionInstance

The ProvisionInstance message.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - completed
  - cancelled
  - errored
  - reassignedByError



## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Cancelled`                                                                                          | [*C1APIPolicyV1CancelledAction](../../models/shared/c1apipolicyv1cancelledaction.md)                 | :heavy_minus_sign:                                                                                   | The CancelledAction message.                                                                         |
| `Completed`                                                                                          | [*C1APIPolicyV1CompletedAction](../../models/shared/c1apipolicyv1completedaction.md)                 | :heavy_minus_sign:                                                                                   | The CompletedAction message.                                                                         |
| `Errored`                                                                                            | [*C1APIPolicyV1ErroredAction](../../models/shared/c1apipolicyv1erroredaction.md)                     | :heavy_minus_sign:                                                                                   | The ErroredAction message.                                                                           |
| `NotificationID`                                                                                     | **string*                                                                                            | :heavy_minus_sign:                                                                                   | The notificationId field.                                                                            |
| `Provision`                                                                                          | [*C1APIPolicyV1Provision](../../models/shared/c1apipolicyv1provision.md)                             | :heavy_minus_sign:                                                                                   | The Provision message.                                                                               |
| `ReassignedByError`                                                                                  | [*C1APIPolicyV1ReassignedByErrorAction](../../models/shared/c1apipolicyv1reassignedbyerroraction.md) | :heavy_minus_sign:                                                                                   | The ReassignedByErrorAction message.                                                                 |
| `State`                                                                                              | [*C1APIPolicyV1ProvisionInstanceState](../../models/shared/c1apipolicyv1provisioninstancestate.md)   | :heavy_minus_sign:                                                                                   | The state field.                                                                                     |