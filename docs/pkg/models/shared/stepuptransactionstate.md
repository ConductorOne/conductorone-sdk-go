# StepUpTransactionState

Current state of the transaction

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StepUpTransactionStateStepUpTransactionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StepUpTransactionState("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `StepUpTransactionStateStepUpTransactionStateUnspecified` | STEP_UP_TRANSACTION_STATE_UNSPECIFIED                     |
| `StepUpTransactionStateStepUpTransactionStatePending`     | STEP_UP_TRANSACTION_STATE_PENDING                         |
| `StepUpTransactionStateStepUpTransactionStateVerified`    | STEP_UP_TRANSACTION_STATE_VERIFIED                        |
| `StepUpTransactionStateStepUpTransactionStateError`       | STEP_UP_TRANSACTION_STATE_ERROR                           |