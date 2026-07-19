# SupportedOauth2Modes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SupportedOauth2ModesMcpServerAuthOauth2ModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SupportedOauth2Modes("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeUnspecified`          | MCP_SERVER_AUTH_OAUTH2_MODE_UNSPECIFIED                           |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeService`              | MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE                               |
| `SupportedOauth2ModesMcpServerAuthOauth2ModePassthrough`          | MCP_SERVER_AUTH_OAUTH2_MODE_PASSTHROUGH                           |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeClientCredentials`    | MCP_SERVER_AUTH_OAUTH2_MODE_CLIENT_CREDENTIALS                    |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeJwtBearer`            | MCP_SERVER_AUTH_OAUTH2_MODE_JWT_BEARER                            |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeGoogleServiceAccount` | MCP_SERVER_AUTH_OAUTH2_MODE_GOOGLE_SERVICE_ACCOUNT                |
| `SupportedOauth2ModesMcpServerAuthOauth2ModeAuthorizationCode`    | MCP_SERVER_AUTH_OAUTH2_MODE_AUTHORIZATION_CODE                    |