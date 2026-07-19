# TaskActionInstanceInput

ActionInstance is the API mirror of the internal immutable snapshot of an
 Action captured on a TaskTypeAction at ticket-creation time.

This message contains a oneof named target_ref. Only a single field of the following list may be set at a time:
  - batonResourceActionRef
  - connectorActionRef



## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `BatonResourceActionRef`                                                                         | [*shared.BatonResourceActionRefInput](../../../pkg/models/shared/batonresourceactionrefinput.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `ConnectorActionRef`                                                                             | [*shared.ConnectorActionRefInput](../../../pkg/models/shared/connectoractionrefinput.md)         | :heavy_minus_sign:                                                                               | N/A                                                                                              |