# FollowedTier

C1 tier this route follows; UNSPECIFIED means it follows no tier. On write, a route that sets a tier must
 use the route_id c1-<tier>; the server replaces its targets and selection with the tier's current model,
 served through the providers this Edge already has credentials for, and also publishes that model's name
 as a second route. Clearing the tier (and publication) on a published route pins it to its current
 targets.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FollowedTierEdgeInferenceTierUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FollowedTier("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `FollowedTierEdgeInferenceTierUnspecified` | EDGE_INFERENCE_TIER_UNSPECIFIED            |
| `FollowedTierEdgeInferenceTierFast`        | EDGE_INFERENCE_TIER_FAST                   |
| `FollowedTierEdgeInferenceTierBalanced`    | EDGE_INFERENCE_TIER_BALANCED               |
| `FollowedTierEdgeInferenceTierBest`        | EDGE_INFERENCE_TIER_BEST                   |