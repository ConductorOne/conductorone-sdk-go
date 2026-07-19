# StateFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StateFilterMcpToolStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StateFilter("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `StateFilterMcpToolStateUnspecified`   | MCP_TOOL_STATE_UNSPECIFIED             |
| `StateFilterMcpToolStatePendingReview` | MCP_TOOL_STATE_PENDING_REVIEW          |
| `StateFilterMcpToolStateApproved`      | MCP_TOOL_STATE_APPROVED                |
| `StateFilterMcpToolStateDisabled`      | MCP_TOOL_STATE_DISABLED                |
| `StateFilterMcpToolStateRemoved`       | MCP_TOOL_STATE_REMOVED                 |