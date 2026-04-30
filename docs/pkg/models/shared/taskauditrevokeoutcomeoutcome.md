# TaskAuditRevokeOutcomeOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditRevokeOutcomeOutcomeRevokeOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditRevokeOutcomeOutcome("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeUnspecified`  | REVOKE_OUTCOME_UNSPECIFIED                               |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeRevoked`      | REVOKE_OUTCOME_REVOKED                                   |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeDenied`       | REVOKE_OUTCOME_DENIED                                    |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeError`        | REVOKE_OUTCOME_ERROR                                     |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeCancelled`    | REVOKE_OUTCOME_CANCELLED                                 |
| `TaskAuditRevokeOutcomeOutcomeRevokeOutcomeWaitTimedOut` | REVOKE_OUTCOME_WAIT_TIMED_OUT                            |