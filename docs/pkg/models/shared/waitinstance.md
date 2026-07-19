# WaitInstance

Used by the policy engine to describe an instantiated wait step.

This message contains a oneof named until. Only a single field of the following list may be set at a time:
  - condition
  - untilTime


This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - succeeded
  - timedOut
  - skipped



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `CommentOnFirstWait`                                                                 | `*string`                                                                            | :heavy_minus_sign:                                                                   | The comment to post on first failed check.                                           |
| `CommentOnTimeout`                                                                   | `*string`                                                                            | :heavy_minus_sign:                                                                   | The comment to post if we timeout.                                                   |
| `Condition`                                                                          | [*shared.WaitConditionInstance](../../../pkg/models/shared/waitconditioninstance.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Name`                                                                               | `*string`                                                                            | :heavy_minus_sign:                                                                   | The name field.                                                                      |
| `Skipped`                                                                            | [*shared.SkippedAction](../../../pkg/models/shared/skippedaction.md)                 | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `StartedWaitingAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `State`                                                                              | [*shared.WaitInstanceState](../../../pkg/models/shared/waitinstancestate.md)         | :heavy_minus_sign:                                                                   | The state field.                                                                     |
| `Succeeded`                                                                          | [*shared.ConditionSucceeded](../../../pkg/models/shared/conditionsucceeded.md)       | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `TimedOut`                                                                           | [*shared.ConditionTimedOut](../../../pkg/models/shared/conditiontimedout.md)         | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Timeout`                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `TimeoutDuration`                                                                    | `*string`                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `UntilTime`                                                                          | [*shared.WaitUntilTimeInstance](../../../pkg/models/shared/waituntiltimeinstance.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |