# RequireToolApproval

Optional per-server override for tool auto-approval. See
 MCPServerView.require_tool_approval for semantics.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequireToolApprovalOptionalBoolUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequireToolApproval("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `RequireToolApprovalOptionalBoolUnspecified` | OPTIONAL_BOOL_UNSPECIFIED                    |
| `RequireToolApprovalOptionalBoolTrue`        | OPTIONAL_BOOL_TRUE                           |
| `RequireToolApprovalOptionalBoolFalse`       | OPTIONAL_BOOL_FALSE                          |