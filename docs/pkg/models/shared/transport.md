# Transport

Denormalized from DeviceInventoryMcpObservation.transport at detection
 time. Distinguishes a locally-launched (STDIO) product, which has no
 address to connect to, from a remote one, which does.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TransportShadowMcpTransportUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Transport("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `TransportShadowMcpTransportUnspecified` | SHADOW_MCP_TRANSPORT_UNSPECIFIED         |
| `TransportShadowMcpTransportStdio`       | SHADOW_MCP_TRANSPORT_STDIO               |
| `TransportShadowMcpTransportRemote`      | SHADOW_MCP_TRANSPORT_REMOTE              |