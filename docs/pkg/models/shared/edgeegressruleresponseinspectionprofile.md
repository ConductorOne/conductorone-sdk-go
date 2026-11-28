# EdgeEgressRuleResponseInspectionProfile

Response-body inspection override, only applied to ALLOW rules.
 Unspecified inherits the Edge's profile; DENY and AUTHZEN clear this field.
 Requests always use strict inspection.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeEgressRuleResponseInspectionProfileEdgeResponseInspectionProfileUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeEgressRuleResponseInspectionProfile("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `EdgeEgressRuleResponseInspectionProfileEdgeResponseInspectionProfileUnspecified` | EDGE_RESPONSE_INSPECTION_PROFILE_UNSPECIFIED                                      |
| `EdgeEgressRuleResponseInspectionProfileEdgeResponseInspectionProfileStrict`      | EDGE_RESPONSE_INSPECTION_PROFILE_STRICT                                           |
| `EdgeEgressRuleResponseInspectionProfileEdgeResponseInspectionProfileWebBalanced` | EDGE_RESPONSE_INSPECTION_PROFILE_WEB_BALANCED                                     |