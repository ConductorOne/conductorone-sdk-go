# CurrentStep

Search tasks that have this type of step as the current step.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CurrentStepTaskSearchCurrentStepUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CurrentStep("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CurrentStepTaskSearchCurrentStepUnspecified` | TASK_SEARCH_CURRENT_STEP_UNSPECIFIED          |
| `CurrentStepTaskSearchCurrentStepApproval`    | TASK_SEARCH_CURRENT_STEP_APPROVAL             |
| `CurrentStepTaskSearchCurrentStepProvision`   | TASK_SEARCH_CURRENT_STEP_PROVISION            |