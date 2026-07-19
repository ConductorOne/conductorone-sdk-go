# ~~DefaultAuthMethod~~

Deprecated: read auth_modes instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultAuthMethodMcpServerAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultAuthMethod("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `DefaultAuthMethodMcpServerAuthMethodUnspecified`  | MCP_SERVER_AUTH_METHOD_UNSPECIFIED                 |
| `DefaultAuthMethodMcpServerAuthMethodNone`         | MCP_SERVER_AUTH_METHOD_NONE                        |
| `DefaultAuthMethodMcpServerAuthMethodBearerToken`  | MCP_SERVER_AUTH_METHOD_BEARER_TOKEN                |
| `DefaultAuthMethodMcpServerAuthMethodOauth2`       | MCP_SERVER_AUTH_METHOD_OAUTH2                      |
| `DefaultAuthMethodMcpServerAuthMethodCustomHeader` | MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER               |
| `DefaultAuthMethodMcpServerAuthMethodAwsSigv4`     | MCP_SERVER_AUTH_METHOD_AWS_SIGV4                   |
| `DefaultAuthMethodMcpServerAuthMethodBasicAuth`    | MCP_SERVER_AUTH_METHOD_BASIC_AUTH                  |