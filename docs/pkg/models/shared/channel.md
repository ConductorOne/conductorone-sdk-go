# Channel

Release channel for this catalog entry.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ChannelMcpServerCatalogChannelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Channel("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ChannelMcpServerCatalogChannelUnspecified` | MCP_SERVER_CATALOG_CHANNEL_UNSPECIFIED      |
| `ChannelMcpServerCatalogChannelStable`      | MCP_SERVER_CATALOG_CHANNEL_STABLE           |
| `ChannelMcpServerCatalogChannelBeta`        | MCP_SERVER_CATALOG_CHANNEL_BETA             |
| `ChannelMcpServerCatalogChannelAlpha`       | MCP_SERVER_CATALOG_CHANNEL_ALPHA            |