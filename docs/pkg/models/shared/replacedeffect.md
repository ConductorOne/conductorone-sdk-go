# ReplacedEffect

The effect of an every-user rule for exactly this host that the decision
 replaced, or unspecified when none was.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ReplacedEffectEdgeEgressEffectUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ReplacedEffect("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ReplacedEffectEdgeEgressEffectUnspecified` | EDGE_EGRESS_EFFECT_UNSPECIFIED              |
| `ReplacedEffectEdgeEgressEffectAllow`       | EDGE_EGRESS_EFFECT_ALLOW                    |
| `ReplacedEffectEdgeEgressEffectDeny`        | EDGE_EGRESS_EFFECT_DENY                     |
| `ReplacedEffectEdgeEgressEffectAuthzen`     | EDGE_EGRESS_EFFECT_AUTHZEN                  |