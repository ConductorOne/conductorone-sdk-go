# AgentStatus

AI-agent lifecycle status when this app user carries the agent trait.
 UNSPECIFIED marks a non-agent account. Read-only; translated from the
 model's agent_trait at the API boundary.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgentStatusAppUserAgentStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgentStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `AgentStatusAppUserAgentStatusUnspecified` | APP_USER_AGENT_STATUS_UNSPECIFIED          |
| `AgentStatusAppUserAgentStatusReady`       | APP_USER_AGENT_STATUS_READY                |
| `AgentStatusAppUserAgentStatusDisabled`    | APP_USER_AGENT_STATUS_DISABLED             |
| `AgentStatusAppUserAgentStatusDeleted`     | APP_USER_AGENT_STATUS_DELETED              |