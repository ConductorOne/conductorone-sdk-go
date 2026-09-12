# MCPServerAuthOAuth2Mode

OAuth2 mode.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerAuthOAuth2Mode("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeUnspecified`          | MCP_SERVER_AUTH_OAUTH2_MODE_UNSPECIFIED                              |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeService`              | MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE                                  |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModePassthrough`          | MCP_SERVER_AUTH_OAUTH2_MODE_PASSTHROUGH                              |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeClientCredentials`    | MCP_SERVER_AUTH_OAUTH2_MODE_CLIENT_CREDENTIALS                       |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeJwtBearer`            | MCP_SERVER_AUTH_OAUTH2_MODE_JWT_BEARER                               |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeGoogleServiceAccount` | MCP_SERVER_AUTH_OAUTH2_MODE_GOOGLE_SERVICE_ACCOUNT                   |
| `MCPServerAuthOAuth2ModeMcpServerAuthOauth2ModeAuthorizationCode`    | MCP_SERVER_AUTH_OAUTH2_MODE_AUTHORIZATION_CODE                       |