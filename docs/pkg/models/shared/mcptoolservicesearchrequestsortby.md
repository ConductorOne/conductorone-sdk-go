# MCPToolServiceSearchRequestSortBy

Sort order for results. UNSPECIFIED sorts by tool name ascending.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPToolServiceSearchRequestSortByMcpToolSortByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPToolServiceSearchRequestSortBy("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `MCPToolServiceSearchRequestSortByMcpToolSortByUnspecified`    | MCP_TOOL_SORT_BY_UNSPECIFIED                                   |
| `MCPToolServiceSearchRequestSortByMcpToolSortByToolName`       | MCP_TOOL_SORT_BY_TOOL_NAME                                     |
| `MCPToolServiceSearchRequestSortByMcpToolSortByVisibility`     | MCP_TOOL_SORT_BY_VISIBILITY                                    |
| `MCPToolServiceSearchRequestSortByMcpToolSortByClassification` | MCP_TOOL_SORT_BY_CLASSIFICATION                                |
| `MCPToolServiceSearchRequestSortByMcpToolSortByState`          | MCP_TOOL_SORT_BY_STATE                                         |
| `MCPToolServiceSearchRequestSortByMcpToolSortByUpdatedAt`      | MCP_TOOL_SORT_BY_UPDATED_AT                                    |