# TaskActionsServiceEscalateToEmergencyAccessRequest

The TaskActionsServiceEscalateToEmergencyAccessRequest object lets you escalate a task to the emergency access workflow.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Comment`                                                              | `*string`                                                              | :heavy_minus_sign:                                                     | An optional comment attached to the escalation.                        |
| `ExpandMask`                                                           | [*shared.TaskExpandMask](../../../pkg/models/shared/taskexpandmask.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PolicyStepID`                                                         | `*string`                                                              | :heavy_minus_sign:                                                     | The ID of the current policy step being escalated from.                |