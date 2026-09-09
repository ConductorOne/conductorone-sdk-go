# RequestTaskState

Current state of the associated access request task.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestTaskStateTicketStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestTaskState("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RequestTaskStateTicketStateUnspecified` | TICKET_STATE_UNSPECIFIED                 |
| `RequestTaskStateTicketStateOpen`        | TICKET_STATE_OPEN                        |
| `RequestTaskStateTicketStateClosed`      | TICKET_STATE_CLOSED                      |