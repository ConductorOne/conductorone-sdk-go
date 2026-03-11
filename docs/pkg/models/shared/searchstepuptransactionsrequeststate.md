# SearchStepUpTransactionsRequestState

Filter by transaction state

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SearchStepUpTransactionsRequestStateStepUpTransactionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SearchStepUpTransactionsRequestState("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `SearchStepUpTransactionsRequestStateStepUpTransactionStateUnspecified` | STEP_UP_TRANSACTION_STATE_UNSPECIFIED                                   |
| `SearchStepUpTransactionsRequestStateStepUpTransactionStatePending`     | STEP_UP_TRANSACTION_STATE_PENDING                                       |
| `SearchStepUpTransactionsRequestStateStepUpTransactionStateVerified`    | STEP_UP_TRANSACTION_STATE_VERIFIED                                      |
| `SearchStepUpTransactionsRequestStateStepUpTransactionStateError`       | STEP_UP_TRANSACTION_STATE_ERROR                                         |