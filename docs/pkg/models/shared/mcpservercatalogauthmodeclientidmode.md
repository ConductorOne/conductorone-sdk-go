# MCPServerCatalogAuthModeClientIDMode

How the OAuth2 client_id is acquired for this mode. Set by the impl bundle
 and shown read-only on the form. authorization_code grant only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerCatalogAuthModeClientIDModeMcpServerCatalogClientIDModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerCatalogAuthModeClientIDMode("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `MCPServerCatalogAuthModeClientIDModeMcpServerCatalogClientIDModeUnspecified` | MCP_SERVER_CATALOG_CLIENT_ID_MODE_UNSPECIFIED                                 |
| `MCPServerCatalogAuthModeClientIDModeMcpServerCatalogClientIDModeManual`      | MCP_SERVER_CATALOG_CLIENT_ID_MODE_MANUAL                                      |
| `MCPServerCatalogAuthModeClientIDModeMcpServerCatalogClientIDModeDcr`         | MCP_SERVER_CATALOG_CLIENT_ID_MODE_DCR                                         |
| `MCPServerCatalogAuthModeClientIDModeMcpServerCatalogClientIDModeCimd`        | MCP_SERVER_CATALOG_CLIENT_ID_MODE_CIMD                                        |