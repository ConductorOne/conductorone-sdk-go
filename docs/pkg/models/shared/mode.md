# Mode

OAuth2 mode.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ModeMcpServerAuthOauth2ModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Mode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ModeMcpServerAuthOauth2ModeUnspecified`           | MCP_SERVER_AUTH_OAUTH2_MODE_UNSPECIFIED            |
| `ModeMcpServerAuthOauth2ModeService`               | MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE                |
| `ModeMcpServerAuthOauth2ModePassthrough`           | MCP_SERVER_AUTH_OAUTH2_MODE_PASSTHROUGH            |
| `ModeMcpServerAuthOauth2ModeClientCredentials`     | MCP_SERVER_AUTH_OAUTH2_MODE_CLIENT_CREDENTIALS     |
| `ModeMcpServerAuthOauth2ModeJwtBearer`             | MCP_SERVER_AUTH_OAUTH2_MODE_JWT_BEARER             |
| `ModeMcpServerAuthOauth2ModeGoogleServiceAccount`  | MCP_SERVER_AUTH_OAUTH2_MODE_GOOGLE_SERVICE_ACCOUNT |
| `ModeMcpServerAuthOauth2ModeAuthorizationCode`     | MCP_SERVER_AUTH_OAUTH2_MODE_AUTHORIZATION_CODE     |