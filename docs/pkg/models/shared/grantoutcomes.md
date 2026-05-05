# GrantOutcomes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GrantOutcomesGrantOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GrantOutcomes("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `GrantOutcomesGrantOutcomeUnspecified`  | GRANT_OUTCOME_UNSPECIFIED               |
| `GrantOutcomesGrantOutcomeGranted`      | GRANT_OUTCOME_GRANTED                   |
| `GrantOutcomesGrantOutcomeDenied`       | GRANT_OUTCOME_DENIED                    |
| `GrantOutcomesGrantOutcomeError`        | GRANT_OUTCOME_ERROR                     |
| `GrantOutcomesGrantOutcomeCancelled`    | GRANT_OUTCOME_CANCELLED                 |
| `GrantOutcomesGrantOutcomeWaitTimedOut` | GRANT_OUTCOME_WAIT_TIMED_OUT            |