# WriteTools

Enablement for tools that create or modify records.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WriteToolsMcpToolEnablementUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.WriteTools("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `WriteToolsMcpToolEnablementUnspecified`       | MCP_TOOL_ENABLEMENT_UNSPECIFIED                |
| `WriteToolsMcpToolEnablementAutoEnable`        | MCP_TOOL_ENABLEMENT_AUTO_ENABLE                |
| `WriteToolsMcpToolEnablementRequireEnablement` | MCP_TOOL_ENABLEMENT_REQUIRE_ENABLEMENT         |