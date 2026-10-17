# Tier

The tier field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TierEdgeInferenceTierUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Tier("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `TierEdgeInferenceTierUnspecified` | EDGE_INFERENCE_TIER_UNSPECIFIED    |
| `TierEdgeInferenceTierFast`        | EDGE_INFERENCE_TIER_FAST           |
| `TierEdgeInferenceTierBalanced`    | EDGE_INFERENCE_TIER_BALANCED       |
| `TierEdgeInferenceTierBest`        | EDGE_INFERENCE_TIER_BEST           |