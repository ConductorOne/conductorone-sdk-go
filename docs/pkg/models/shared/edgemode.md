# EdgeMode

Global e-stop for the connection-routing configuration. DISABLED (and
 UNSPECIFIED) denies every capability; OBSERVE evaluates policy but serves
 the request, recording what ENFORCE would have answered; ENFORCE serves
 the policy decision.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeModeEdgeModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeMode("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EdgeModeEdgeModeUnspecified` | EDGE_MODE_UNSPECIFIED         |
| `EdgeModeEdgeModeDisabled`    | EDGE_MODE_DISABLED            |
| `EdgeModeEdgeModeObserve`     | EDGE_MODE_OBSERVE             |
| `EdgeModeEdgeModeEnforce`     | EDGE_MODE_ENFORCE             |