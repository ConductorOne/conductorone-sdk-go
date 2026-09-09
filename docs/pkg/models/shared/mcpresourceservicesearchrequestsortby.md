# MCPResourceServiceSearchRequestSortBy

Sort order for results. UNSPECIFIED sorts by resource name ascending.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPResourceServiceSearchRequestSortByMcpResourceSortByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPResourceServiceSearchRequestSortBy("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `MCPResourceServiceSearchRequestSortByMcpResourceSortByUnspecified` | MCP_RESOURCE_SORT_BY_UNSPECIFIED                                    |
| `MCPResourceServiceSearchRequestSortByMcpResourceSortByName`        | MCP_RESOURCE_SORT_BY_NAME                                           |
| `MCPResourceServiceSearchRequestSortByMcpResourceSortByURI`         | MCP_RESOURCE_SORT_BY_URI                                            |
| `MCPResourceServiceSearchRequestSortByMcpResourceSortByState`       | MCP_RESOURCE_SORT_BY_STATE                                          |
| `MCPResourceServiceSearchRequestSortByMcpResourceSortByUpdatedAt`   | MCP_RESOURCE_SORT_BY_UPDATED_AT                                     |