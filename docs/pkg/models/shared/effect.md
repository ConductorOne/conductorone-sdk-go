# Effect

The effect field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EffectEdgeEgressEffectUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Effect("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `EffectEdgeEgressEffectUnspecified` | EDGE_EGRESS_EFFECT_UNSPECIFIED      |
| `EffectEdgeEgressEffectAllow`       | EDGE_EGRESS_EFFECT_ALLOW            |
| `EffectEdgeEgressEffectDeny`        | EDGE_EGRESS_EFFECT_DENY             |