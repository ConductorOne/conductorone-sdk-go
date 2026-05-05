# TaskStates

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskStatesTaskStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskStates("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `TaskStatesTaskStateUnspecified` | TASK_STATE_UNSPECIFIED           |
| `TaskStatesTaskStateOpen`        | TASK_STATE_OPEN                  |
| `TaskStatesTaskStateClosed`      | TASK_STATE_CLOSED                |