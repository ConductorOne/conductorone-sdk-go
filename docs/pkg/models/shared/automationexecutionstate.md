# AutomationExecutionState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AutomationExecutionStateAutomationExecutionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AutomationExecutionState("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AutomationExecutionStateAutomationExecutionStateUnspecified`            | AUTOMATION_EXECUTION_STATE_UNSPECIFIED                                   |
| `AutomationExecutionStateAutomationExecutionStatePending`                | AUTOMATION_EXECUTION_STATE_PENDING                                       |
| `AutomationExecutionStateAutomationExecutionStateCreating`               | AUTOMATION_EXECUTION_STATE_CREATING                                      |
| `AutomationExecutionStateAutomationExecutionStateGetStep`                | AUTOMATION_EXECUTION_STATE_GET_STEP                                      |
| `AutomationExecutionStateAutomationExecutionStateProcessStep`            | AUTOMATION_EXECUTION_STATE_PROCESS_STEP                                  |
| `AutomationExecutionStateAutomationExecutionStateCompleteStep`           | AUTOMATION_EXECUTION_STATE_COMPLETE_STEP                                 |
| `AutomationExecutionStateAutomationExecutionStateDone`                   | AUTOMATION_EXECUTION_STATE_DONE                                          |
| `AutomationExecutionStateAutomationExecutionStateError`                  | AUTOMATION_EXECUTION_STATE_ERROR                                         |
| `AutomationExecutionStateAutomationExecutionStateTerminate`              | AUTOMATION_EXECUTION_STATE_TERMINATE                                     |
| `AutomationExecutionStateAutomationExecutionStateWaiting`                | AUTOMATION_EXECUTION_STATE_WAITING                                       |
| `AutomationExecutionStateAutomationExecutionStatePausedByCircuitBreaker` | AUTOMATION_EXECUTION_STATE_PAUSED_BY_CIRCUIT_BREAKER                     |