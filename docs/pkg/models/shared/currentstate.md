# CurrentState

The currentState field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CurrentStateTaskStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CurrentState("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CurrentStateTaskStateUnspecified` | TASK_STATE_UNSPECIFIED             |
| `CurrentStateTaskStateOpen`        | TASK_STATE_OPEN                    |
| `CurrentStateTaskStateClosed`      | TASK_STATE_CLOSED                  |