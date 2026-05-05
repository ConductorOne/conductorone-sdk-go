# TaskTypeActionOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeActionOutcomeActionOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeActionOutcome("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `TaskTypeActionOutcomeActionOutcomeUnspecified` | ACTION_OUTCOME_UNSPECIFIED                      |
| `TaskTypeActionOutcomeActionOutcomeSuccess`     | ACTION_OUTCOME_SUCCESS                          |
| `TaskTypeActionOutcomeActionOutcomeDenied`      | ACTION_OUTCOME_DENIED                           |
| `TaskTypeActionOutcomeActionOutcomeError`       | ACTION_OUTCOME_ERROR                            |
| `TaskTypeActionOutcomeActionOutcomeCancelled`   | ACTION_OUTCOME_CANCELLED                        |