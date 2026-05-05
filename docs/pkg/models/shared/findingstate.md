# FindingState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingStateFindingStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingState("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `FindingStateFindingStateUnspecified`  | FINDING_STATE_UNSPECIFIED              |
| `FindingStateFindingStateOpen`         | FINDING_STATE_OPEN                     |
| `FindingStateFindingStateInProgress`   | FINDING_STATE_IN_PROGRESS              |
| `FindingStateFindingStateResolved`     | FINDING_STATE_RESOLVED                 |
| `FindingStateFindingStateSnoozed`      | FINDING_STATE_SNOOZED                  |
| `FindingStateFindingStateRiskAccepted` | FINDING_STATE_RISK_ACCEPTED            |
| `FindingStateFindingStateSuppressed`   | FINDING_STATE_SUPPRESSED               |