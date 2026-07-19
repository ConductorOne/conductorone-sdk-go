# ClientIDMode

How the client_id is acquired for authorization_code mode. When DCR (or
 CIMD), client_id / client_secret are not required on input — the gateway
 registers itself with the authorization server during Register and injects
 the result. UNSPECIFIED means manual (admin-entered client_id).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClientIDModeMcpServerAuthOauth2ClientIDModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClientIDMode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `ClientIDModeMcpServerAuthOauth2ClientIDModeUnspecified` | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_UNSPECIFIED        |
| `ClientIDModeMcpServerAuthOauth2ClientIDModeDcr`         | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_DCR                |
| `ClientIDModeMcpServerAuthOauth2ClientIDModeCimd`        | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_CIMD               |