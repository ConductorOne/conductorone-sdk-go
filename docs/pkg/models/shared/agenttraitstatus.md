# AgentTraitStatus

The agent's lifecycle status (READY, DISABLED, DELETED).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgentTraitStatusAgentStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgentTraitStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AgentTraitStatusAgentStatusUnspecified` | AGENT_STATUS_UNSPECIFIED                 |
| `AgentTraitStatusAgentStatusReady`       | AGENT_STATUS_READY                       |
| `AgentTraitStatusAgentStatusDisabled`    | AGENT_STATUS_DISABLED                    |
| `AgentTraitStatusAgentStatusDeleted`     | AGENT_STATUS_DELETED                     |