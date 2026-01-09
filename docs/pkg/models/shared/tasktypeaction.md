# TaskTypeAction

The TaskTypeAction message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ActionID`                                                                           | **string*                                                                            | :heavy_minus_sign:                                                                   | The ID of the action to execute.                                                     |
| `FormValues`                                                                         | map[string]*any*                                                                     | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Outcome`                                                                            | [*shared.TaskTypeActionOutcome](../../../pkg/models/shared/tasktypeactionoutcome.md) | :heavy_minus_sign:                                                                   | The outcome field.                                                                   |
| `OutcomeTime`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |