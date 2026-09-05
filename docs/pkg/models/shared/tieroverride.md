# TierOverride

Author tier override; may only tighten the derived tier.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TierOverrideFindingDispatchTierUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TierOverride("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `TierOverrideFindingDispatchTierUnspecified`      | FINDING_DISPATCH_TIER_UNSPECIFIED                 |
| `TierOverrideFindingDispatchTierAuto`             | FINDING_DISPATCH_TIER_AUTO                        |
| `TierOverrideFindingDispatchTierRequiresApproval` | FINDING_DISPATCH_TIER_REQUIRES_APPROVAL           |