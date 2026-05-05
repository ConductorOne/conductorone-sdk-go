# ExecutionStates

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExecutionStatesAutomationExecutionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExecutionStates("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ExecutionStatesAutomationExecutionStateUnspecified`  | AUTOMATION_EXECUTION_STATE_UNSPECIFIED                |
| `ExecutionStatesAutomationExecutionStatePending`      | AUTOMATION_EXECUTION_STATE_PENDING                    |
| `ExecutionStatesAutomationExecutionStateCreating`     | AUTOMATION_EXECUTION_STATE_CREATING                   |
| `ExecutionStatesAutomationExecutionStateGetStep`      | AUTOMATION_EXECUTION_STATE_GET_STEP                   |
| `ExecutionStatesAutomationExecutionStateProcessStep`  | AUTOMATION_EXECUTION_STATE_PROCESS_STEP               |
| `ExecutionStatesAutomationExecutionStateCompleteStep` | AUTOMATION_EXECUTION_STATE_COMPLETE_STEP              |
| `ExecutionStatesAutomationExecutionStateDone`         | AUTOMATION_EXECUTION_STATE_DONE                       |
| `ExecutionStatesAutomationExecutionStateError`        | AUTOMATION_EXECUTION_STATE_ERROR                      |
| `ExecutionStatesAutomationExecutionStateTerminate`    | AUTOMATION_EXECUTION_STATE_TERMINATE                  |
| `ExecutionStatesAutomationExecutionStateWaiting`      | AUTOMATION_EXECUTION_STATE_WAITING                    |