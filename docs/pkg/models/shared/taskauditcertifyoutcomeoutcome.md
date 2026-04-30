# TaskAuditCertifyOutcomeOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditCertifyOutcomeOutcomeCertifyOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditCertifyOutcomeOutcome("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeUnspecified`  | CERTIFY_OUTCOME_UNSPECIFIED                                |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeCertified`    | CERTIFY_OUTCOME_CERTIFIED                                  |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeDecertified`  | CERTIFY_OUTCOME_DECERTIFIED                                |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeError`        | CERTIFY_OUTCOME_ERROR                                      |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeCancelled`    | CERTIFY_OUTCOME_CANCELLED                                  |
| `TaskAuditCertifyOutcomeOutcomeCertifyOutcomeWaitTimedOut` | CERTIFY_OUTCOME_WAIT_TIMED_OUT                             |