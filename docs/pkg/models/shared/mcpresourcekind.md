# MCPResourceKind

Whether this is a static resource or a URI template.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPResourceKindMcpResourceKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPResourceKind("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `MCPResourceKindMcpResourceKindUnspecified` | MCP_RESOURCE_KIND_UNSPECIFIED               |
| `MCPResourceKindMcpResourceKindStatic`      | MCP_RESOURCE_KIND_STATIC                    |
| `MCPResourceKindMcpResourceKindTemplate`    | MCP_RESOURCE_KIND_TEMPLATE                  |