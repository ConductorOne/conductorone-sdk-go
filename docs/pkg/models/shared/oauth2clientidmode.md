# Oauth2ClientIDMode

How the OAuth2 client_id was acquired (manual / DCR / CIMD). Read-only;
 derived from stored config. Empty for non-authorization_code servers.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.Oauth2ClientIDModeMcpServerAuthOauth2ClientIDModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Oauth2ClientIDMode("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Oauth2ClientIDModeMcpServerAuthOauth2ClientIDModeUnspecified` | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_UNSPECIFIED              |
| `Oauth2ClientIDModeMcpServerAuthOauth2ClientIDModeDcr`         | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_DCR                      |
| `Oauth2ClientIDModeMcpServerAuthOauth2ClientIDModeCimd`        | MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_CIMD                     |