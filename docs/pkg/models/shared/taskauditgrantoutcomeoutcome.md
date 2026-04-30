# TaskAuditGrantOutcomeOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditGrantOutcomeOutcomeGrantOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditGrantOutcomeOutcome("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeUnspecified`  | GRANT_OUTCOME_UNSPECIFIED                              |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeGranted`      | GRANT_OUTCOME_GRANTED                                  |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeDenied`       | GRANT_OUTCOME_DENIED                                   |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeError`        | GRANT_OUTCOME_ERROR                                    |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeCancelled`    | GRANT_OUTCOME_CANCELLED                                |
| `TaskAuditGrantOutcomeOutcomeGrantOutcomeWaitTimedOut` | GRANT_OUTCOME_WAIT_TIMED_OUT                           |