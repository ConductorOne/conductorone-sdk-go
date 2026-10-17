# AuthzenServerMode

The e-stop, now a tri-state. DISABLED (and UNSPECIFIED) denies every
 call without touching the entitlement stack or its grants; OBSERVE
 evaluates fully but answers allow, recording what ENFORCE would have
 said; ENFORCE answers the policy's decision.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthzenServerModeAuthzenModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthzenServerMode("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AuthzenServerModeAuthzenModeUnspecified` | AUTHZEN_MODE_UNSPECIFIED                  |
| `AuthzenServerModeAuthzenModeDisabled`    | AUTHZEN_MODE_DISABLED                     |
| `AuthzenServerModeAuthzenModeObserve`     | AUTHZEN_MODE_OBSERVE                      |
| `AuthzenServerModeAuthzenModeEnforce`     | AUTHZEN_MODE_ENFORCE                      |