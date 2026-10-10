# RequestTaskState

Current state of the associated access request task.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestTaskStateTaskStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestTaskState("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `RequestTaskStateTaskStateUnspecified` | TASK_STATE_UNSPECIFIED                 |
| `RequestTaskStateTaskStateOpen`        | TASK_STATE_OPEN                        |
| `RequestTaskStateTaskStateClosed`      | TASK_STATE_CLOSED                      |