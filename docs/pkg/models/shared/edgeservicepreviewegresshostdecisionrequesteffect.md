# EdgeServicePreviewEgressHostDecisionRequestEffect

The effect field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeServicePreviewEgressHostDecisionRequestEffectEdgeEgressEffectUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeServicePreviewEgressHostDecisionRequestEffect("custom_value")
```


## Values

| Name                                                                           | Value                                                                          |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `EdgeServicePreviewEgressHostDecisionRequestEffectEdgeEgressEffectUnspecified` | EDGE_EGRESS_EFFECT_UNSPECIFIED                                                 |
| `EdgeServicePreviewEgressHostDecisionRequestEffectEdgeEgressEffectAllow`       | EDGE_EGRESS_EFFECT_ALLOW                                                       |
| `EdgeServicePreviewEgressHostDecisionRequestEffectEdgeEgressEffectDeny`        | EDGE_EGRESS_EFFECT_DENY                                                        |
| `EdgeServicePreviewEgressHostDecisionRequestEffectEdgeEgressEffectAuthzen`     | EDGE_EGRESS_EFFECT_AUTHZEN                                                     |