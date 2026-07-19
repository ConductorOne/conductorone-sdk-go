# StateAtEvent

The stateAtEvent field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StateAtEventFindingStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StateAtEvent("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `StateAtEventFindingStateUnspecified`  | FINDING_STATE_UNSPECIFIED              |
| `StateAtEventFindingStateOpen`         | FINDING_STATE_OPEN                     |
| `StateAtEventFindingStateInProgress`   | FINDING_STATE_IN_PROGRESS              |
| `StateAtEventFindingStateResolved`     | FINDING_STATE_RESOLVED                 |
| `StateAtEventFindingStateSnoozed`      | FINDING_STATE_SNOOZED                  |
| `StateAtEventFindingStateRiskAccepted` | FINDING_STATE_RISK_ACCEPTED            |
| `StateAtEventFindingStateSuppressed`   | FINDING_STATE_SUPPRESSED               |