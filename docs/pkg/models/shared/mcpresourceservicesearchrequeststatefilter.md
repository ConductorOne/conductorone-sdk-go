# MCPResourceServiceSearchRequestStateFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPResourceServiceSearchRequestStateFilterMcpResourceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPResourceServiceSearchRequestStateFilter("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `MCPResourceServiceSearchRequestStateFilterMcpResourceStateUnspecified`   | MCP_RESOURCE_STATE_UNSPECIFIED                                            |
| `MCPResourceServiceSearchRequestStateFilterMcpResourceStatePendingReview` | MCP_RESOURCE_STATE_PENDING_REVIEW                                         |
| `MCPResourceServiceSearchRequestStateFilterMcpResourceStateApproved`      | MCP_RESOURCE_STATE_APPROVED                                               |
| `MCPResourceServiceSearchRequestStateFilterMcpResourceStateDisabled`      | MCP_RESOURCE_STATE_DISABLED                                               |
| `MCPResourceServiceSearchRequestStateFilterMcpResourceStateRemoved`       | MCP_RESOURCE_STATE_REMOVED                                                |