# DestructiveTools

Enablement for tools that delete data or cannot be undone.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DestructiveToolsMcpToolEnablementUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DestructiveTools("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `DestructiveToolsMcpToolEnablementUnspecified`       | MCP_TOOL_ENABLEMENT_UNSPECIFIED                      |
| `DestructiveToolsMcpToolEnablementAutoEnable`        | MCP_TOOL_ENABLEMENT_AUTO_ENABLE                      |
| `DestructiveToolsMcpToolEnablementRequireEnablement` | MCP_TOOL_ENABLEMENT_REQUIRE_ENABLEMENT               |