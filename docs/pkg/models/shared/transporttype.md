# TransportType

Transport type for the MCP connection.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TransportTypeMcpServerTransportTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TransportType("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `TransportTypeMcpServerTransportTypeUnspecified`    | MCP_SERVER_TRANSPORT_TYPE_UNSPECIFIED               |
| `TransportTypeMcpServerTransportTypeStreamableHTTP` | MCP_SERVER_TRANSPORT_TYPE_STREAMABLE_HTTP           |
| `TransportTypeMcpServerTransportTypeSse`            | MCP_SERVER_TRANSPORT_TYPE_SSE                       |