# TBEgressDenialGroupOutcome

DENIED, WOULD_DENY or ALLOWED. A flow an observe-mode Edge's policy
 denied but a later check also denied is DENIED.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBEgressDenialGroupOutcomeTbEgressFlowOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBEgressDenialGroupOutcome("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `TBEgressDenialGroupOutcomeTbEgressFlowOutcomeUnspecified` | TB_EGRESS_FLOW_OUTCOME_UNSPECIFIED                         |
| `TBEgressDenialGroupOutcomeTbEgressFlowOutcomeDenied`      | TB_EGRESS_FLOW_OUTCOME_DENIED                              |
| `TBEgressDenialGroupOutcomeTbEgressFlowOutcomeWouldDeny`   | TB_EGRESS_FLOW_OUTCOME_WOULD_DENY                          |
| `TBEgressDenialGroupOutcomeTbEgressFlowOutcomeAllowed`     | TB_EGRESS_FLOW_OUTCOME_ALLOWED                             |