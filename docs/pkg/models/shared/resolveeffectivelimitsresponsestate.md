# ResolveEffectiveLimitsResponseState

Result of resolving the user's effective limits.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ResolveEffectiveLimitsResponseStateEffectiveLimitsStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ResolveEffectiveLimitsResponseState("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ResolveEffectiveLimitsResponseStateEffectiveLimitsStateUnspecified` | EFFECTIVE_LIMITS_STATE_UNSPECIFIED                                   |
| `ResolveEffectiveLimitsResponseStateEffectiveLimitsStateResolved`    | EFFECTIVE_LIMITS_STATE_RESOLVED                                      |
| `ResolveEffectiveLimitsResponseStateEffectiveLimitsStateNoPolicy`    | EFFECTIVE_LIMITS_STATE_NO_POLICY                                     |
| `ResolveEffectiveLimitsResponseStateEffectiveLimitsStateNoCeiling`   | EFFECTIVE_LIMITS_STATE_NO_CEILING                                    |
| `ResolveEffectiveLimitsResponseStateEffectiveLimitsStateDenied`      | EFFECTIVE_LIMITS_STATE_DENIED                                        |