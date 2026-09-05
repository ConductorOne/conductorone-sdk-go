# MCPToolServiceSearchRequestStateFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPToolServiceSearchRequestStateFilterMcpToolStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPToolServiceSearchRequestStateFilter("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `MCPToolServiceSearchRequestStateFilterMcpToolStateUnspecified`   | MCP_TOOL_STATE_UNSPECIFIED                                        |
| `MCPToolServiceSearchRequestStateFilterMcpToolStatePendingReview` | MCP_TOOL_STATE_PENDING_REVIEW                                     |
| `MCPToolServiceSearchRequestStateFilterMcpToolStateApproved`      | MCP_TOOL_STATE_APPROVED                                           |
| `MCPToolServiceSearchRequestStateFilterMcpToolStateDisabled`      | MCP_TOOL_STATE_DISABLED                                           |
| `MCPToolServiceSearchRequestStateFilterMcpToolStateRemoved`       | MCP_TOOL_STATE_REMOVED                                            |