# RetryAutomationExecutionResult

The RetryAutomationExecutionResult message.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ExecutionID`                                                      | `*int64`                                                           | :heavy_minus_sign:                                                 | The automation execution this result is for.                       |
| `Retried`                                                          | `*bool`                                                            | :heavy_minus_sign:                                                 | True when the execution was queued to resume from its failed step. |
| `SkippedReason`                                                    | `*string`                                                          | :heavy_minus_sign:                                                 | Why the execution was not retried. Empty when retried is true.     |