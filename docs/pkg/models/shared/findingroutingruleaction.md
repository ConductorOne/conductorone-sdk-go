# FindingRoutingRuleAction

The FindingRoutingRuleAction message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - createTask
  - suppress
  - snooze
  - acceptRisk



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `AcceptRisk`                                                                             | [*shared.AcceptRiskRoutingAction](../../../pkg/models/shared/acceptriskroutingaction.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `CreateTask`                                                                             | [*shared.CreateTaskAction](../../../pkg/models/shared/createtaskaction.md)               | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Snooze`                                                                                 | [*shared.SnoozeRoutingAction](../../../pkg/models/shared/snoozeroutingaction.md)         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Suppress`                                                                               | [*shared.SuppressRoutingAction](../../../pkg/models/shared/suppressroutingaction.md)     | :heavy_minus_sign:                                                                       | N/A                                                                                      |