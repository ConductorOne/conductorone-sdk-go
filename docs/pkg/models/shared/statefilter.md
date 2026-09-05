# StateFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StateFilterMcpResourceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StateFilter("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `StateFilterMcpResourceStateUnspecified`   | MCP_RESOURCE_STATE_UNSPECIFIED             |
| `StateFilterMcpResourceStatePendingReview` | MCP_RESOURCE_STATE_PENDING_REVIEW          |
| `StateFilterMcpResourceStateApproved`      | MCP_RESOURCE_STATE_APPROVED                |
| `StateFilterMcpResourceStateDisabled`      | MCP_RESOURCE_STATE_DISABLED                |
| `StateFilterMcpResourceStateRemoved`       | MCP_RESOURCE_STATE_REMOVED                 |