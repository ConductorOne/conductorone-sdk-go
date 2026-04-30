# TaskTypeCertifyOutcome

The outcome of the certification.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeCertifyOutcome("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `TaskTypeCertifyOutcomeCertifyOutcomeUnspecified`  | CERTIFY_OUTCOME_UNSPECIFIED                        |
| `TaskTypeCertifyOutcomeCertifyOutcomeCertified`    | CERTIFY_OUTCOME_CERTIFIED                          |
| `TaskTypeCertifyOutcomeCertifyOutcomeDecertified`  | CERTIFY_OUTCOME_DECERTIFIED                        |
| `TaskTypeCertifyOutcomeCertifyOutcomeError`        | CERTIFY_OUTCOME_ERROR                              |
| `TaskTypeCertifyOutcomeCertifyOutcomeCancelled`    | CERTIFY_OUTCOME_CANCELLED                          |
| `TaskTypeCertifyOutcomeCertifyOutcomeWaitTimedOut` | CERTIFY_OUTCOME_WAIT_TIMED_OUT                     |