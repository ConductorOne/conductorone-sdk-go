# AuthMethod

Auth method on the connector. Drives the FE connect-flow choice:
 OAUTH2 → redirect via CreateAuthorizeURL; BEARER_TOKEN / CUSTOM_HEADER
 → form dialog via SubmitUserCredential.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthMethodMcpServerAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AuthMethodMcpServerAuthMethodUnspecified`  | MCP_SERVER_AUTH_METHOD_UNSPECIFIED          |
| `AuthMethodMcpServerAuthMethodNone`         | MCP_SERVER_AUTH_METHOD_NONE                 |
| `AuthMethodMcpServerAuthMethodBearerToken`  | MCP_SERVER_AUTH_METHOD_BEARER_TOKEN         |
| `AuthMethodMcpServerAuthMethodOauth2`       | MCP_SERVER_AUTH_METHOD_OAUTH2               |
| `AuthMethodMcpServerAuthMethodCustomHeader` | MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER        |
| `AuthMethodMcpServerAuthMethodAwsSigv4`     | MCP_SERVER_AUTH_METHOD_AWS_SIGV4            |
| `AuthMethodMcpServerAuthMethodBasicAuth`    | MCP_SERVER_AUTH_METHOD_BASIC_AUTH           |