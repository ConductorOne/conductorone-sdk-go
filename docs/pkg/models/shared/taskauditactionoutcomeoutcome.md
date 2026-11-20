# TaskAuditActionOutcomeOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditActionOutcomeOutcomeActionOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditActionOutcomeOutcome("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `TaskAuditActionOutcomeOutcomeActionOutcomeUnspecified` | ACTION_OUTCOME_UNSPECIFIED                              |
| `TaskAuditActionOutcomeOutcomeActionOutcomeSuccess`     | ACTION_OUTCOME_SUCCESS                                  |
| `TaskAuditActionOutcomeOutcomeActionOutcomeDenied`      | ACTION_OUTCOME_DENIED                                   |
| `TaskAuditActionOutcomeOutcomeActionOutcomeError`       | ACTION_OUTCOME_ERROR                                    |
| `TaskAuditActionOutcomeOutcomeActionOutcomeCancelled`   | ACTION_OUTCOME_CANCELLED                                |