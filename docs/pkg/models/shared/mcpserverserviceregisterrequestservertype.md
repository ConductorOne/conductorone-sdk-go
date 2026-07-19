# MCPServerServiceRegisterRequestServerType

The type of MCP server being registered.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerServiceRegisterRequestServerTypeMcpServerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerServiceRegisterRequestServerType("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `MCPServerServiceRegisterRequestServerTypeMcpServerTypeUnspecified` | MCP_SERVER_TYPE_UNSPECIFIED                                         |
| `MCPServerServiceRegisterRequestServerTypeMcpServerTypeHosted`      | MCP_SERVER_TYPE_HOSTED                                              |
| `MCPServerServiceRegisterRequestServerTypeMcpServerTypeExternal`    | MCP_SERVER_TYPE_EXTERNAL                                            |