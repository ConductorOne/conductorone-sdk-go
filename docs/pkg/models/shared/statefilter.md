# StateFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StateFilterMcpPromptStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StateFilter("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `StateFilterMcpPromptStateUnspecified`   | MCP_PROMPT_STATE_UNSPECIFIED             |
| `StateFilterMcpPromptStatePendingReview` | MCP_PROMPT_STATE_PENDING_REVIEW          |
| `StateFilterMcpPromptStateApproved`      | MCP_PROMPT_STATE_APPROVED                |
| `StateFilterMcpPromptStateDisabled`      | MCP_PROMPT_STATE_DISABLED                |
| `StateFilterMcpPromptStateRemoved`       | MCP_PROMPT_STATE_REMOVED                 |