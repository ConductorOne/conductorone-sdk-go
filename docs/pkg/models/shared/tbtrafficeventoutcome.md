# TBTrafficEventOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBTrafficEventOutcomeTbTrafficOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBTrafficEventOutcome("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `TBTrafficEventOutcomeTbTrafficOutcomeUnspecified` | TB_TRAFFIC_OUTCOME_UNSPECIFIED                     |
| `TBTrafficEventOutcomeTbTrafficOutcomeAllowed`     | TB_TRAFFIC_OUTCOME_ALLOWED                         |
| `TBTrafficEventOutcomeTbTrafficOutcomeDenied`      | TB_TRAFFIC_OUTCOME_DENIED                          |