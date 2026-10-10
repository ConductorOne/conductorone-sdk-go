# Kind

Restrict resources and scopes to this capability. UNSPECIFIED returns all
 authorized capabilities; unknown enum values are rejected.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.KindEdgeCapabilityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.Kind("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `KindEdgeCapabilityKindUnspecified` | EDGE_CAPABILITY_KIND_UNSPECIFIED    |
| `KindEdgeCapabilityKindEgress`      | EDGE_CAPABILITY_KIND_EGRESS         |
| `KindEdgeCapabilityKindInference`   | EDGE_CAPABILITY_KIND_INFERENCE      |