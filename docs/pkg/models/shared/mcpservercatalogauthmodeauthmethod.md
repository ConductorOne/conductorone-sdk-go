# MCPServerCatalogAuthModeAuthMethod

Authentication method enum. UNSPECIFIED entries are dropped on the way out.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerCatalogAuthModeAuthMethod("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodUnspecified`  | MCP_SERVER_AUTH_METHOD_UNSPECIFIED                                  |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodNone`         | MCP_SERVER_AUTH_METHOD_NONE                                         |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodBearerToken`  | MCP_SERVER_AUTH_METHOD_BEARER_TOKEN                                 |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodOauth2`       | MCP_SERVER_AUTH_METHOD_OAUTH2                                       |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodCustomHeader` | MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER                                |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodAwsSigv4`     | MCP_SERVER_AUTH_METHOD_AWS_SIGV4                                    |
| `MCPServerCatalogAuthModeAuthMethodMcpServerAuthMethodBasicAuth`    | MCP_SERVER_AUTH_METHOD_BASIC_AUTH                                   |