# ToolState

Which tool state to return count for on each server.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ToolStateMcpToolStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ToolState("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ToolStateMcpToolStateUnspecified`   | MCP_TOOL_STATE_UNSPECIFIED           |
| `ToolStateMcpToolStatePendingReview` | MCP_TOOL_STATE_PENDING_REVIEW        |
| `ToolStateMcpToolStateApproved`      | MCP_TOOL_STATE_APPROVED              |
| `ToolStateMcpToolStateDisabled`      | MCP_TOOL_STATE_DISABLED              |
| `ToolStateMcpToolStateRemoved`       | MCP_TOOL_STATE_REMOVED               |