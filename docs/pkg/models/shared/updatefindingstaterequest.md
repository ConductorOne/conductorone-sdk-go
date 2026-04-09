# UpdateFindingStateRequest

The UpdateFindingStateRequest message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - snooze
  - suppress
  - acceptRisk
  - unsuppress
  - resolve



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AcceptRiskAction`                                                               | [*shared.AcceptRiskAction](../../../pkg/models/shared/acceptriskaction.md)       | :heavy_minus_sign:                                                               | AcceptRiskAction parameters for UpdateFindingState.                              |
| `ResolveAction`                                                                  | [*shared.ResolveAction](../../../pkg/models/shared/resolveaction.md)             | :heavy_minus_sign:                                                               | ResolveAction parameters for UpdateFindingState (manual resolve).                |
| `SnoozeAction`                                                                   | [*shared.SnoozeAction](../../../pkg/models/shared/snoozeaction.md)               | :heavy_minus_sign:                                                               | SnoozeAction parameters for UpdateFindingState.                                  |
| `SuppressStateAction`                                                            | [*shared.SuppressStateAction](../../../pkg/models/shared/suppressstateaction.md) | :heavy_minus_sign:                                                               | SuppressStateAction parameters for UpdateFindingState.                           |
| `UnsuppressAction`                                                               | [*shared.UnsuppressAction](../../../pkg/models/shared/unsuppressaction.md)       | :heavy_minus_sign:                                                               | UnsuppressAction parameters for UpdateFindingState.                              |