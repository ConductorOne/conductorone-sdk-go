# MCPServerViewServerType

Whether this is a hosted MCP server.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewServerTypeMcpServerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewServerType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `MCPServerViewServerTypeMcpServerTypeUnspecified` | MCP_SERVER_TYPE_UNSPECIFIED                       |
| `MCPServerViewServerTypeMcpServerTypeHosted`      | MCP_SERVER_TYPE_HOSTED                            |
| `MCPServerViewServerTypeMcpServerTypeExternal`    | MCP_SERVER_TYPE_EXTERNAL                          |