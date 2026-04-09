# TaskTypeGrantOutcome

The outcome of the grant.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeGrantOutcomeGrantOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeGrantOutcome("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TaskTypeGrantOutcomeGrantOutcomeUnspecified`  | GRANT_OUTCOME_UNSPECIFIED                      |
| `TaskTypeGrantOutcomeGrantOutcomeGranted`      | GRANT_OUTCOME_GRANTED                          |
| `TaskTypeGrantOutcomeGrantOutcomeDenied`       | GRANT_OUTCOME_DENIED                           |
| `TaskTypeGrantOutcomeGrantOutcomeError`        | GRANT_OUTCOME_ERROR                            |
| `TaskTypeGrantOutcomeGrantOutcomeCancelled`    | GRANT_OUTCOME_CANCELLED                        |
| `TaskTypeGrantOutcomeGrantOutcomeWaitTimedOut` | GRANT_OUTCOME_WAIT_TIMED_OUT                   |