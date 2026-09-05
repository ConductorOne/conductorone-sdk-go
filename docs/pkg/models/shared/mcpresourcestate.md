# MCPResourceState

Resource approval/lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPResourceStateMcpResourceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPResourceState("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `MCPResourceStateMcpResourceStateUnspecified`   | MCP_RESOURCE_STATE_UNSPECIFIED                  |
| `MCPResourceStateMcpResourceStatePendingReview` | MCP_RESOURCE_STATE_PENDING_REVIEW               |
| `MCPResourceStateMcpResourceStateApproved`      | MCP_RESOURCE_STATE_APPROVED                     |
| `MCPResourceStateMcpResourceStateDisabled`      | MCP_RESOURCE_STATE_DISABLED                     |
| `MCPResourceStateMcpResourceStateRemoved`       | MCP_RESOURCE_STATE_REMOVED                      |