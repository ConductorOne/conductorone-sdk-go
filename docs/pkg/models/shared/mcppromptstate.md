# MCPPromptState

Administrative approval and discovery lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPPromptStateMcpPromptStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPPromptState("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `MCPPromptStateMcpPromptStateUnspecified`   | MCP_PROMPT_STATE_UNSPECIFIED                |
| `MCPPromptStateMcpPromptStatePendingReview` | MCP_PROMPT_STATE_PENDING_REVIEW             |
| `MCPPromptStateMcpPromptStateApproved`      | MCP_PROMPT_STATE_APPROVED                   |
| `MCPPromptStateMcpPromptStateDisabled`      | MCP_PROMPT_STATE_DISABLED                   |
| `MCPPromptStateMcpPromptStateRemoved`       | MCP_PROMPT_STATE_REMOVED                    |