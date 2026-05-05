# ActionInstanceState

The current state of the action execution.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ActionInstanceStateActionInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ActionInstanceState("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `ActionInstanceStateActionInstanceStateUnspecified` | ACTION_INSTANCE_STATE_UNSPECIFIED                   |
| `ActionInstanceStateActionInstanceStateInit`        | ACTION_INSTANCE_STATE_INIT                          |
| `ActionInstanceStateActionInstanceStateRunning`     | ACTION_INSTANCE_STATE_RUNNING                       |
| `ActionInstanceStateActionInstanceStateDone`        | ACTION_INSTANCE_STATE_DONE                          |
| `ActionInstanceStateActionInstanceStateError`       | ACTION_INSTANCE_STATE_ERROR                         |