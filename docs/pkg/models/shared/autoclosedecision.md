# AutoCloseDecision

The autoCloseDecision field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AutoCloseDecisionCloseDecisionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AutoCloseDecision("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AutoCloseDecisionCloseDecisionUnspecified` | CLOSE_DECISION_UNSPECIFIED                  |
| `AutoCloseDecisionCloseDecisionRevoked`     | CLOSE_DECISION_REVOKED                      |
| `AutoCloseDecisionCloseDecisionSkip`        | CLOSE_DECISION_SKIP                         |
| `AutoCloseDecisionCloseDecisionNoAction`    | CLOSE_DECISION_NO_ACTION                    |