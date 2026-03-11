# Outcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OutcomeAccessRequestOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Outcome("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `OutcomeAccessRequestOutcomeUnspecified` | ACCESS_REQUEST_OUTCOME_UNSPECIFIED       |
| `OutcomeAccessRequestOutcomeApproved`    | ACCESS_REQUEST_OUTCOME_APPROVED          |
| `OutcomeAccessRequestOutcomeDenied`      | ACCESS_REQUEST_OUTCOME_DENIED            |
| `OutcomeAccessRequestOutcomeError`       | ACCESS_REQUEST_OUTCOME_ERROR             |
| `OutcomeAccessRequestOutcomeCancelled`   | ACCESS_REQUEST_OUTCOME_CANCELLED         |