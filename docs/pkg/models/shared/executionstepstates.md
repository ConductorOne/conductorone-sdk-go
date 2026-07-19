# ExecutionStepStates

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExecutionStepStatesAutomationExecutionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExecutionStepStates("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ExecutionStepStatesAutomationExecutionStateUnspecified`            | AUTOMATION_EXECUTION_STATE_UNSPECIFIED                              |
| `ExecutionStepStatesAutomationExecutionStatePending`                | AUTOMATION_EXECUTION_STATE_PENDING                                  |
| `ExecutionStepStatesAutomationExecutionStateCreating`               | AUTOMATION_EXECUTION_STATE_CREATING                                 |
| `ExecutionStepStatesAutomationExecutionStateGetStep`                | AUTOMATION_EXECUTION_STATE_GET_STEP                                 |
| `ExecutionStepStatesAutomationExecutionStateProcessStep`            | AUTOMATION_EXECUTION_STATE_PROCESS_STEP                             |
| `ExecutionStepStatesAutomationExecutionStateCompleteStep`           | AUTOMATION_EXECUTION_STATE_COMPLETE_STEP                            |
| `ExecutionStepStatesAutomationExecutionStateDone`                   | AUTOMATION_EXECUTION_STATE_DONE                                     |
| `ExecutionStepStatesAutomationExecutionStateError`                  | AUTOMATION_EXECUTION_STATE_ERROR                                    |
| `ExecutionStepStatesAutomationExecutionStateTerminate`              | AUTOMATION_EXECUTION_STATE_TERMINATE                                |
| `ExecutionStepStatesAutomationExecutionStateWaiting`                | AUTOMATION_EXECUTION_STATE_WAITING                                  |
| `ExecutionStepStatesAutomationExecutionStatePausedByCircuitBreaker` | AUTOMATION_EXECUTION_STATE_PAUSED_BY_CIRCUIT_BREAKER                |