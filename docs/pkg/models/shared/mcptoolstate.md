# MCPToolState

Tool approval/lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPToolStateMcpToolStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPToolState("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `MCPToolStateMcpToolStateUnspecified`   | MCP_TOOL_STATE_UNSPECIFIED              |
| `MCPToolStateMcpToolStatePendingReview` | MCP_TOOL_STATE_PENDING_REVIEW           |
| `MCPToolStateMcpToolStateApproved`      | MCP_TOOL_STATE_APPROVED                 |
| `MCPToolStateMcpToolStateDisabled`      | MCP_TOOL_STATE_DISABLED                 |
| `MCPToolStateMcpToolStateRemoved`       | MCP_TOOL_STATE_REMOVED                  |