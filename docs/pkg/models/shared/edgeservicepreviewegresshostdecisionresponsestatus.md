# EdgeServicePreviewEgressHostDecisionResponseStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeServicePreviewEgressHostDecisionResponseStatusStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeServicePreviewEgressHostDecisionResponseStatus("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusUnspecified`        | STATUS_UNSPECIFIED                                                           |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusReady`              | STATUS_READY                                                                 |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusAlreadyDecided`     | STATUS_ALREADY_DECIDED                                                       |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusInvalidHost`        | STATUS_INVALID_HOST                                                          |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusSsrfBlocked`        | STATUS_SSRF_BLOCKED                                                          |
| `EdgeServicePreviewEgressHostDecisionResponseStatusStatusUnresolvedDenyRule` | STATUS_UNRESOLVED_DENY_RULE                                                  |