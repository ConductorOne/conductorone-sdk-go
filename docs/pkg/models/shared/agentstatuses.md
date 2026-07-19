# AgentStatuses

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgentStatusesAgentStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgentStatuses("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AgentStatusesAgentStatusUnspecified` | AGENT_STATUS_UNSPECIFIED              |
| `AgentStatusesAgentStatusReady`       | AGENT_STATUS_READY                    |
| `AgentStatusesAgentStatusDisabled`    | AGENT_STATUS_DISABLED                 |
| `AgentStatusesAgentStatusDeleted`     | AGENT_STATUS_DELETED                  |