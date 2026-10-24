# Capability

The capability field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CapabilityTbTrafficCapabilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Capability("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CapabilityTbTrafficCapabilityUnspecified` | TB_TRAFFIC_CAPABILITY_UNSPECIFIED          |
| `CapabilityTbTrafficCapabilityEgress`      | TB_TRAFFIC_CAPABILITY_EGRESS               |
| `CapabilityTbTrafficCapabilityLlm`         | TB_TRAFFIC_CAPABILITY_LLM                  |