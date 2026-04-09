# AgentMode

The mode of the agent, full control, change policy only, or comment only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgentModeApprovalAgentModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgentMode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AgentModeApprovalAgentModeUnspecified`      | APPROVAL_AGENT_MODE_UNSPECIFIED              |
| `AgentModeApprovalAgentModeFullControl`      | APPROVAL_AGENT_MODE_FULL_CONTROL             |
| `AgentModeApprovalAgentModeChangePolicyOnly` | APPROVAL_AGENT_MODE_CHANGE_POLICY_ONLY       |
| `AgentModeApprovalAgentModeCommentOnly`      | APPROVAL_AGENT_MODE_COMMENT_ONLY             |