# DefaultOutcome

The defaultOutcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultOutcomeTbEgressOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultOutcome("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `DefaultOutcomeTbEgressOutcomeUnspecified` | TB_EGRESS_OUTCOME_UNSPECIFIED              |
| `DefaultOutcomeTbEgressOutcomeAllowed`     | TB_EGRESS_OUTCOME_ALLOWED                  |
| `DefaultOutcomeTbEgressOutcomeDenied`      | TB_EGRESS_OUTCOME_DENIED                   |