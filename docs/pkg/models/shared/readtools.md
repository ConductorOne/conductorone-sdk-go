# ReadTools

Enablement for tools that only query data.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ReadToolsMcpToolEnablementUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ReadTools("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ReadToolsMcpToolEnablementUnspecified`       | MCP_TOOL_ENABLEMENT_UNSPECIFIED               |
| `ReadToolsMcpToolEnablementAutoEnable`        | MCP_TOOL_ENABLEMENT_AUTO_ENABLE               |
| `ReadToolsMcpToolEnablementRequireEnablement` | MCP_TOOL_ENABLEMENT_REQUIRE_ENABLEMENT        |