# WaitInstanceState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WaitInstanceStateWaitInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.WaitInstanceState("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `WaitInstanceStateWaitInstanceStateUnspecified` | WAIT_INSTANCE_STATE_UNSPECIFIED                 |
| `WaitInstanceStateWaitInstanceStateWaiting`     | WAIT_INSTANCE_STATE_WAITING                     |
| `WaitInstanceStateWaitInstanceStateCompleted`   | WAIT_INSTANCE_STATE_COMPLETED                   |
| `WaitInstanceStateWaitInstanceStateTimedOut`    | WAIT_INSTANCE_STATE_TIMED_OUT                   |