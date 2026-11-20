# EdgeMode

Hosting state for the service configuration. DISABLED and UNSPECIFIED
 deny every request. ENFORCE and OBSERVE both still require capability
 authorization; content policy and routing remain configured on the Edge
 host. OBSERVE is currently identical to ENFORCE -- it exists so an
 operator can stage an Edge ahead of destination-level policy support,
 which does not exist yet.

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