# PolicyStepInstanceState

The state of the step, which is either active or done.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyStepInstanceStatePolicyStepStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyStepInstanceState("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `PolicyStepInstanceStatePolicyStepStateUnspecified` | POLICY_STEP_STATE_UNSPECIFIED                       |
| `PolicyStepInstanceStatePolicyStepStateActive`      | POLICY_STEP_STATE_ACTIVE                            |
| `PolicyStepInstanceStatePolicyStepStateDone`        | POLICY_STEP_STATE_DONE                              |