# WaitInstance

Used by the policy engine to describe an instantiated wait step.

This message contains a oneof named until. Only a single field of the following list may be set at a time:
  - condition


This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - succeeded
  - timedOut



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ConditionSucceeded`                                                                 | [*shared.ConditionSucceeded](../../../pkg/models/shared/conditionsucceeded.md)       | :heavy_minus_sign:                                                                   | The ConditionSucceeded message.                                                      |
| `ConditionTimedOut`                                                                  | [*shared.ConditionTimedOut](../../../pkg/models/shared/conditiontimedout.md)         | :heavy_minus_sign:                                                                   | The ConditionTimedOut message.                                                       |
| `WaitConditionInstance`                                                              | [*shared.WaitConditionInstance](../../../pkg/models/shared/waitconditioninstance.md) | :heavy_minus_sign:                                                                   | Used by the policy engine to describe an instantiated condition to wait on.          |
| `CommentOnFirstWait`                                                                 | **string*                                                                            | :heavy_minus_sign:                                                                   | The comment to post on first failed check.                                           |
| `CommentOnTimeout`                                                                   | **string*                                                                            | :heavy_minus_sign:                                                                   | The comment to post if we timeout.                                                   |
| `Name`                                                                               | **string*                                                                            | :heavy_minus_sign:                                                                   | The name field.                                                                      |
| `StartedWaitingAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `State`                                                                              | [*shared.WaitInstanceState](../../../pkg/models/shared/waitinstancestate.md)         | :heavy_minus_sign:                                                                   | The state field.                                                                     |
| `Timeout`                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `TimeoutDuration`                                                                    | **string*                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |