# MCPServerViewTransportType

Transport type for external MCP servers. Read-only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewTransportTypeMcpServerTransportTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewTransportType("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `MCPServerViewTransportTypeMcpServerTransportTypeUnspecified`    | MCP_SERVER_TRANSPORT_TYPE_UNSPECIFIED                            |
| `MCPServerViewTransportTypeMcpServerTransportTypeStreamableHTTP` | MCP_SERVER_TRANSPORT_TYPE_STREAMABLE_HTTP                        |
| `MCPServerViewTransportTypeMcpServerTransportTypeSse`            | MCP_SERVER_TRANSPORT_TYPE_SSE                                    |