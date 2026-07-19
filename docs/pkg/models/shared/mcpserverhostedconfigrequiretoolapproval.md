# MCPServerHostedConfigRequireToolApproval

Optional per-server override for tool auto-approval. See
 MCPServerView.require_tool_approval for semantics.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerHostedConfigRequireToolApprovalOptionalBoolUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerHostedConfigRequireToolApproval("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolUnspecified` | OPTIONAL_BOOL_UNSPECIFIED                                         |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolTrue`        | OPTIONAL_BOOL_TRUE                                                |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolFalse`       | OPTIONAL_BOOL_FALSE                                               |