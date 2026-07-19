# Maturity

Implementation maturity level.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MaturityMcpServerCatalogMaturityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Maturity("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `MaturityMcpServerCatalogMaturityUnspecified` | MCP_SERVER_CATALOG_MATURITY_UNSPECIFIED       |
| `MaturityMcpServerCatalogMaturityStub`        | MCP_SERVER_CATALOG_MATURITY_STUB              |
| `MaturityMcpServerCatalogMaturityGenerated`   | MCP_SERVER_CATALOG_MATURITY_GENERATED         |
| `MaturityMcpServerCatalogMaturityVerified`    | MCP_SERVER_CATALOG_MATURITY_VERIFIED          |
| `MaturityMcpServerCatalogMaturityCurated`     | MCP_SERVER_CATALOG_MATURITY_CURATED           |