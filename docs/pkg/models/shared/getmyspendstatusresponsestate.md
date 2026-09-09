# GetMySpendStatusResponseState

Result of resolving the caller's effective limits.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GetMySpendStatusResponseStateEffectiveLimitsStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GetMySpendStatusResponseState("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `GetMySpendStatusResponseStateEffectiveLimitsStateUnspecified` | EFFECTIVE_LIMITS_STATE_UNSPECIFIED                             |
| `GetMySpendStatusResponseStateEffectiveLimitsStateResolved`    | EFFECTIVE_LIMITS_STATE_RESOLVED                                |
| `GetMySpendStatusResponseStateEffectiveLimitsStateNoPolicy`    | EFFECTIVE_LIMITS_STATE_NO_POLICY                               |
| `GetMySpendStatusResponseStateEffectiveLimitsStateNoCeiling`   | EFFECTIVE_LIMITS_STATE_NO_CEILING                              |
| `GetMySpendStatusResponseStateEffectiveLimitsStateDenied`      | EFFECTIVE_LIMITS_STATE_DENIED                                  |