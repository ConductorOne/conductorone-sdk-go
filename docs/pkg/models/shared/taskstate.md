# TaskState

The current state of the task as defined by the `state_enum`

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskStateTaskStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskState("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `TaskStateTaskStateUnspecified` | TASK_STATE_UNSPECIFIED          |
| `TaskStateTaskStateOpen`        | TASK_STATE_OPEN                 |
| `TaskStateTaskStateClosed`      | TASK_STATE_CLOSED               |