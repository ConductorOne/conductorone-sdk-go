# EdgePublicTrustMetadataState

Available authority or explicit withdrawal.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgePublicTrustMetadataStateEdgePublicTrustStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgePublicTrustMetadataState("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `EdgePublicTrustMetadataStateEdgePublicTrustStateUnspecified` | EDGE_PUBLIC_TRUST_STATE_UNSPECIFIED                           |
| `EdgePublicTrustMetadataStateEdgePublicTrustStateAvailable`   | EDGE_PUBLIC_TRUST_STATE_AVAILABLE                             |
| `EdgePublicTrustMetadataStateEdgePublicTrustStateWithdrawn`   | EDGE_PUBLIC_TRUST_STATE_WITHDRAWN                             |