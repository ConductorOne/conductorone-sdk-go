# FindingRoutingRuleAction

FindingRoutingRuleAction defines what happens when a rule matches.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - createTask
  - suppress
  - notify



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CreateTaskAction`                                                         | [*shared.CreateTaskAction](../../../pkg/models/shared/createtaskaction.md) | :heavy_minus_sign:                                                         | The CreateTaskAction message.                                              |
| `NotifyAction`                                                             | [*shared.NotifyAction](../../../pkg/models/shared/notifyaction.md)         | :heavy_minus_sign:                                                         | The NotifyAction message.                                                  |
| `SuppressAction`                                                           | [*shared.SuppressAction](../../../pkg/models/shared/suppressaction.md)     | :heavy_minus_sign:                                                         | The SuppressAction message.                                                |