# States

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StatesFindingStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.States("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `StatesFindingStateUnspecified`  | FINDING_STATE_UNSPECIFIED        |
| `StatesFindingStateOpen`         | FINDING_STATE_OPEN               |
| `StatesFindingStateInProgress`   | FINDING_STATE_IN_PROGRESS        |
| `StatesFindingStateResolved`     | FINDING_STATE_RESOLVED           |
| `StatesFindingStateSnoozed`      | FINDING_STATE_SNOOZED            |
| `StatesFindingStateRiskAccepted` | FINDING_STATE_RISK_ACCEPTED      |
| `StatesFindingStateSuppressed`   | FINDING_STATE_SUPPRESSED         |