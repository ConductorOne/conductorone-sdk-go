# TBEgressPolicyDefaultOutcome

The defaultOutcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBEgressPolicyDefaultOutcomeTbEgressOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBEgressPolicyDefaultOutcome("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `TBEgressPolicyDefaultOutcomeTbEgressOutcomeUnspecified` | TB_EGRESS_OUTCOME_UNSPECIFIED                            |
| `TBEgressPolicyDefaultOutcomeTbEgressOutcomeAllowed`     | TB_EGRESS_OUTCOME_ALLOWED                                |
| `TBEgressPolicyDefaultOutcomeTbEgressOutcomeDenied`      | TB_EGRESS_OUTCOME_DENIED                                 |