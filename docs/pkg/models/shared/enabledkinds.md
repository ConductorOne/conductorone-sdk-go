# EnabledKinds

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EnabledKindsEdgeCapabilityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EnabledKinds("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `EnabledKindsEdgeCapabilityKindUnspecified` | EDGE_CAPABILITY_KIND_UNSPECIFIED            |
| `EnabledKindsEdgeCapabilityKindEgress`      | EDGE_CAPABILITY_KIND_EGRESS                 |
| `EnabledKindsEdgeCapabilityKindInference`   | EDGE_CAPABILITY_KIND_INFERENCE              |