# PolicyActionInstanceState

The current state of the action execution.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyActionInstanceStateActionInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyActionInstanceState("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `PolicyActionInstanceStateActionInstanceStateUnspecified` | ACTION_INSTANCE_STATE_UNSPECIFIED                         |
| `PolicyActionInstanceStateActionInstanceStateInit`        | ACTION_INSTANCE_STATE_INIT                                |
| `PolicyActionInstanceStateActionInstanceStateRunning`     | ACTION_INSTANCE_STATE_RUNNING                             |
| `PolicyActionInstanceStateActionInstanceStateDone`        | ACTION_INSTANCE_STATE_DONE                                |
| `PolicyActionInstanceStateActionInstanceStateError`       | ACTION_INSTANCE_STATE_ERROR                               |