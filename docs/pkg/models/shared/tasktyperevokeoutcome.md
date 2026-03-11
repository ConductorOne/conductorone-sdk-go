# TaskTypeRevokeOutcome

The outcome of the revoke.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeRevokeOutcomeRevokeOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeRevokeOutcome("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `TaskTypeRevokeOutcomeRevokeOutcomeUnspecified`  | REVOKE_OUTCOME_UNSPECIFIED                       |
| `TaskTypeRevokeOutcomeRevokeOutcomeRevoked`      | REVOKE_OUTCOME_REVOKED                           |
| `TaskTypeRevokeOutcomeRevokeOutcomeDenied`       | REVOKE_OUTCOME_DENIED                            |
| `TaskTypeRevokeOutcomeRevokeOutcomeError`        | REVOKE_OUTCOME_ERROR                             |
| `TaskTypeRevokeOutcomeRevokeOutcomeCancelled`    | REVOKE_OUTCOME_CANCELLED                         |
| `TaskTypeRevokeOutcomeRevokeOutcomeWaitTimedOut` | REVOKE_OUTCOME_WAIT_TIMED_OUT                    |