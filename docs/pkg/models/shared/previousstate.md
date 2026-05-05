# PreviousState

The previousState field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PreviousStateTaskStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PreviousState("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PreviousStateTaskStateUnspecified` | TASK_STATE_UNSPECIFIED              |
| `PreviousStateTaskStateOpen`        | TASK_STATE_OPEN                     |
| `PreviousStateTaskStateClosed`      | TASK_STATE_CLOSED                   |