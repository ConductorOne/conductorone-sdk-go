# MCPConnectionViewAuthMethod

Auth method on the connector. Drives the FE connect-flow choice:
 OAUTH2 → redirect via CreateAuthorizeURL; BEARER_TOKEN / CUSTOM_HEADER
 → form dialog via SubmitUserCredential.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPConnectionViewAuthMethodMcpServerAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPConnectionViewAuthMethod("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodUnspecified`  | MCP_SERVER_AUTH_METHOD_UNSPECIFIED                           |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodNone`         | MCP_SERVER_AUTH_METHOD_NONE                                  |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodBearerToken`  | MCP_SERVER_AUTH_METHOD_BEARER_TOKEN                          |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodOauth2`       | MCP_SERVER_AUTH_METHOD_OAUTH2                                |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodCustomHeader` | MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER                         |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodAwsSigv4`     | MCP_SERVER_AUTH_METHOD_AWS_SIGV4                             |
| `MCPConnectionViewAuthMethodMcpServerAuthMethodBasicAuth`    | MCP_SERVER_AUTH_METHOD_BASIC_AUTH                            |