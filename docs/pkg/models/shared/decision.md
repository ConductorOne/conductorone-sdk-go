# Decision

What to do with paused executions. UNSPECIFIED means clear the breaker
 only (backward-compatible default). RUN or CANCEL creates a bulk action
 to resolve them asynchronously.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DecisionPausedExecutionDecisionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Decision("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `DecisionPausedExecutionDecisionUnspecified` | PAUSED_EXECUTION_DECISION_UNSPECIFIED        |
| `DecisionPausedExecutionDecisionRun`         | PAUSED_EXECUTION_DECISION_RUN                |
| `DecisionPausedExecutionDecisionCancel`      | PAUSED_EXECUTION_DECISION_CANCEL             |