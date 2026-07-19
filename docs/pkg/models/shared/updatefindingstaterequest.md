# UpdateFindingStateRequest

The UpdateFindingStateRequest message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - snooze
  - suppress
  - acceptRisk
  - unsuppress
  - resolve
  - reopen



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AcceptRisk`                                                                     | [*shared.AcceptRiskAction](../../../pkg/models/shared/acceptriskaction.md)       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Reopen`                                                                         | [*shared.ReopenAction](../../../pkg/models/shared/reopenaction.md)               | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Resolve`                                                                        | [*shared.ResolveAction](../../../pkg/models/shared/resolveaction.md)             | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Snooze`                                                                         | [*shared.SnoozeAction](../../../pkg/models/shared/snoozeaction.md)               | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Suppress`                                                                       | [*shared.SuppressStateAction](../../../pkg/models/shared/suppressstateaction.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Unsuppress`                                                                     | [*shared.UnsuppressAction](../../../pkg/models/shared/unsuppressaction.md)       | :heavy_minus_sign:                                                               | N/A                                                                              |