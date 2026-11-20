# MCPToolEnablementConfigMode

How newly discovered tools are enabled.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPToolEnablementConfigModeMcpToolEnablementModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPToolEnablementConfigMode("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `MCPToolEnablementConfigModeMcpToolEnablementModeUnspecified`       | MCP_TOOL_ENABLEMENT_MODE_UNSPECIFIED                                |
| `MCPToolEnablementConfigModeMcpToolEnablementModeAutoEnableAll`     | MCP_TOOL_ENABLEMENT_MODE_AUTO_ENABLE_ALL                            |
| `MCPToolEnablementConfigModeMcpToolEnablementModeRequireEnablement` | MCP_TOOL_ENABLEMENT_MODE_REQUIRE_ENABLEMENT                         |
| `MCPToolEnablementConfigModeMcpToolEnablementModeByClassification`  | MCP_TOOL_ENABLEMENT_MODE_BY_CLASSIFICATION                          |