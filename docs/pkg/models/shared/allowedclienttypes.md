# AllowedClientTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AllowedClientTypesMcpClientTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AllowedClientTypes("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AllowedClientTypesMcpClientTypeUnspecified` | MCP_CLIENT_TYPE_UNSPECIFIED                  |
| `AllowedClientTypesMcpClientTypePersonal`    | MCP_CLIENT_TYPE_PERSONAL                     |
| `AllowedClientTypesMcpClientTypeShared`      | MCP_CLIENT_TYPE_SHARED                       |
| `AllowedClientTypesMcpClientTypeService`     | MCP_CLIENT_TYPE_SERVICE                      |
| `AllowedClientTypesMcpClientTypeEphemeral`   | MCP_CLIENT_TYPE_EPHEMERAL                    |