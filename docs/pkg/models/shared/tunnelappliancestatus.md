# TunnelApplianceStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TunnelApplianceStatusTunnelApplianceStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TunnelApplianceStatus("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `TunnelApplianceStatusTunnelApplianceStatusUnspecified`    | TUNNEL_APPLIANCE_STATUS_UNSPECIFIED                        |
| `TunnelApplianceStatusTunnelApplianceStatusConnected`      | TUNNEL_APPLIANCE_STATUS_CONNECTED                          |
| `TunnelApplianceStatusTunnelApplianceStatusDisconnected`   | TUNNEL_APPLIANCE_STATUS_DISCONNECTED                       |
| `TunnelApplianceStatusTunnelApplianceStatusNeverConnected` | TUNNEL_APPLIANCE_STATUS_NEVER_CONNECTED                    |