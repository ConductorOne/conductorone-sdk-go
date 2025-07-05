# AgentApproval

The agent to assign the task to.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AgentMode`                                                               | [*shared.AgentMode](../../../pkg/models/shared/agentmode.md)              | :heavy_minus_sign:                                                        | The mode of the agent, full control, change policy only, or comment only. |
| `AgentUserID`                                                             | **string*                                                                 | :heavy_minus_sign:                                                        | The agent user ID to assign the task to.                                  |
| `Instructions`                                                            | **string*                                                                 | :heavy_minus_sign:                                                        | Instructions for the agent.                                               |
| `PolicyIds`                                                               | []*string*                                                                | :heavy_minus_sign:                                                        | The allow list of policy IDs to re-route the task to.                     |