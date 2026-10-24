# MCPServerViewRequireToolApproval

Per-server override for tool approval on MCP servers without a saved
 registration access level. UNSPECIFIED inherits the tenant setting.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewRequireToolApprovalOptionalBoolUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewRequireToolApproval("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `MCPServerViewRequireToolApprovalOptionalBoolUnspecified` | OPTIONAL_BOOL_UNSPECIFIED                                 |
| `MCPServerViewRequireToolApprovalOptionalBoolTrue`        | OPTIONAL_BOOL_TRUE                                        |
| `MCPServerViewRequireToolApprovalOptionalBoolFalse`       | OPTIONAL_BOOL_FALSE                                       |