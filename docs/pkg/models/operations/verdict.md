# Verdict

Restricts results to destinations with at least one row of this
 outcome. TB_TRAFFIC_OUTCOME_UNSPECIFIED applies no verdict filter.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.VerdictTbTrafficOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.Verdict("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `VerdictTbTrafficOutcomeUnspecified` | TB_TRAFFIC_OUTCOME_UNSPECIFIED       |
| `VerdictTbTrafficOutcomeAllowed`     | TB_TRAFFIC_OUTCOME_ALLOWED           |
| `VerdictTbTrafficOutcomeDenied`      | TB_TRAFFIC_OUTCOME_DENIED            |