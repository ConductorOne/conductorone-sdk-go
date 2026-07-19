# GraphNodeType

The type field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GraphNodeTypeGraphNodeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GraphNodeType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `GraphNodeTypeGraphNodeTypeUnspecified`  | GRAPH_NODE_TYPE_UNSPECIFIED              |
| `GraphNodeTypeGraphNodeTypeUser`         | GRAPH_NODE_TYPE_USER                     |
| `GraphNodeTypeGraphNodeTypeAppUser`      | GRAPH_NODE_TYPE_APP_USER                 |
| `GraphNodeTypeGraphNodeTypeApp`          | GRAPH_NODE_TYPE_APP                      |
| `GraphNodeTypeGraphNodeTypeResourceType` | GRAPH_NODE_TYPE_RESOURCE_TYPE            |
| `GraphNodeTypeGraphNodeTypeResource`     | GRAPH_NODE_TYPE_RESOURCE                 |
| `GraphNodeTypeGraphNodeTypeEntitlement`  | GRAPH_NODE_TYPE_ENTITLEMENT              |
| `GraphNodeTypeGraphNodeTypeGrant`        | GRAPH_NODE_TYPE_GRANT                    |