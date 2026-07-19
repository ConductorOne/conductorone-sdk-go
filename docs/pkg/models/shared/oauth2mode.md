# Oauth2Mode

OAuth2 mode in use. Read-only; derived from stored config.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.Oauth2ModeMcpServerAuthOauth2ModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Oauth2Mode("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `Oauth2ModeMcpServerAuthOauth2ModeUnspecified`          | MCP_SERVER_AUTH_OAUTH2_MODE_UNSPECIFIED                 |
| `Oauth2ModeMcpServerAuthOauth2ModeService`              | MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE                     |
| `Oauth2ModeMcpServerAuthOauth2ModePassthrough`          | MCP_SERVER_AUTH_OAUTH2_MODE_PASSTHROUGH                 |
| `Oauth2ModeMcpServerAuthOauth2ModeClientCredentials`    | MCP_SERVER_AUTH_OAUTH2_MODE_CLIENT_CREDENTIALS          |
| `Oauth2ModeMcpServerAuthOauth2ModeJwtBearer`            | MCP_SERVER_AUTH_OAUTH2_MODE_JWT_BEARER                  |
| `Oauth2ModeMcpServerAuthOauth2ModeGoogleServiceAccount` | MCP_SERVER_AUTH_OAUTH2_MODE_GOOGLE_SERVICE_ACCOUNT      |
| `Oauth2ModeMcpServerAuthOauth2ModeAuthorizationCode`    | MCP_SERVER_AUTH_OAUTH2_MODE_AUTHORIZATION_CODE          |