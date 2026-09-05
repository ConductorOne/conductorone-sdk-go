# TBEgressRuleOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBEgressRuleOutcomeTbEgressOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBEgressRuleOutcome("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `TBEgressRuleOutcomeTbEgressOutcomeUnspecified` | TB_EGRESS_OUTCOME_UNSPECIFIED                   |
| `TBEgressRuleOutcomeTbEgressOutcomeAllowed`     | TB_EGRESS_OUTCOME_ALLOWED                       |
| `TBEgressRuleOutcomeTbEgressOutcomeDenied`      | TB_EGRESS_OUTCOME_DENIED                        |