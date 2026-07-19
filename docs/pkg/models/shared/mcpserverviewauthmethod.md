# MCPServerViewAuthMethod

Authentication method in use. Read-only; derived from stored config.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewAuthMethodMcpServerAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewAuthMethod("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `MCPServerViewAuthMethodMcpServerAuthMethodUnspecified`  | MCP_SERVER_AUTH_METHOD_UNSPECIFIED                       |
| `MCPServerViewAuthMethodMcpServerAuthMethodNone`         | MCP_SERVER_AUTH_METHOD_NONE                              |
| `MCPServerViewAuthMethodMcpServerAuthMethodBearerToken`  | MCP_SERVER_AUTH_METHOD_BEARER_TOKEN                      |
| `MCPServerViewAuthMethodMcpServerAuthMethodOauth2`       | MCP_SERVER_AUTH_METHOD_OAUTH2                            |
| `MCPServerViewAuthMethodMcpServerAuthMethodCustomHeader` | MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER                     |
| `MCPServerViewAuthMethodMcpServerAuthMethodAwsSigv4`     | MCP_SERVER_AUTH_METHOD_AWS_SIGV4                         |
| `MCPServerViewAuthMethodMcpServerAuthMethodBasicAuth`    | MCP_SERVER_AUTH_METHOD_BASIC_AUTH                        |