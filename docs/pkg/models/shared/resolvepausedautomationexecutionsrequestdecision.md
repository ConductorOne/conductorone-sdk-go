# ResolvePausedAutomationExecutionsRequestDecision

Whether to run or cancel the paused executions.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ResolvePausedAutomationExecutionsRequestDecisionPausedExecutionDecisionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ResolvePausedAutomationExecutionsRequestDecision("custom_value")
```


## Values

| Name                                                                                 | Value                                                                                |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ResolvePausedAutomationExecutionsRequestDecisionPausedExecutionDecisionUnspecified` | PAUSED_EXECUTION_DECISION_UNSPECIFIED                                                |
| `ResolvePausedAutomationExecutionsRequestDecisionPausedExecutionDecisionRun`         | PAUSED_EXECUTION_DECISION_RUN                                                        |
| `ResolvePausedAutomationExecutionsRequestDecisionPausedExecutionDecisionCancel`      | PAUSED_EXECUTION_DECISION_CANCEL                                                     |