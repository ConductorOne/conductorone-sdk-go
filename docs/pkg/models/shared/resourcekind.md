# ResourceKind

Kind of the hosted service.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ResourceKindC1EdgeResourceKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ResourceKind("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ResourceKindC1EdgeResourceKindUnspecified`   | C1_EDGE_RESOURCE_KIND_UNSPECIFIED             |
| `ResourceKindC1EdgeResourceKindEdge`          | C1_EDGE_RESOURCE_KIND_EDGE                    |
| `ResourceKindC1EdgeResourceKindAuthzenServer` | C1_EDGE_RESOURCE_KIND_AUTHZEN_SERVER          |