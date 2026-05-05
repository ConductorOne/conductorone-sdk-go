# TaskTypeFindingOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeFindingOutcomeFindingTaskOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeFindingOutcome("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `TaskTypeFindingOutcomeFindingTaskOutcomeUnspecified`  | FINDING_TASK_OUTCOME_UNSPECIFIED                       |
| `TaskTypeFindingOutcomeFindingTaskOutcomeRemediated`   | FINDING_TASK_OUTCOME_REMEDIATED                        |
| `TaskTypeFindingOutcomeFindingTaskOutcomeRiskAccepted` | FINDING_TASK_OUTCOME_RISK_ACCEPTED                     |
| `TaskTypeFindingOutcomeFindingTaskOutcomeCancelled`    | FINDING_TASK_OUTCOME_CANCELLED                         |