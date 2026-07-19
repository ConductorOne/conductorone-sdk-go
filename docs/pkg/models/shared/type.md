# Type

The type field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TypeGraphEdgeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Type("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `TypeGraphEdgeTypeUnspecified`       | GRAPH_EDGE_TYPE_UNSPECIFIED          |
| `TypeGraphEdgeTypeIdentityLink`      | GRAPH_EDGE_TYPE_IDENTITY_LINK        |
| `TypeGraphEdgeTypeDirectGrant`       | GRAPH_EDGE_TYPE_DIRECT_GRANT         |
| `TypeGraphEdgeTypeAppHierarchy`      | GRAPH_EDGE_TYPE_APP_HIERARCHY        |
| `TypeGraphEdgeTypeResourceHierarchy` | GRAPH_EDGE_TYPE_RESOURCE_HIERARCHY   |
| `TypeGraphEdgeTypeProxyBinding`      | GRAPH_EDGE_TYPE_PROXY_BINDING        |