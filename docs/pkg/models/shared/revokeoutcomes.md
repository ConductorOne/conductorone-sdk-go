# RevokeOutcomes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RevokeOutcomesRevokeOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RevokeOutcomes("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `RevokeOutcomesRevokeOutcomeUnspecified`  | REVOKE_OUTCOME_UNSPECIFIED                |
| `RevokeOutcomesRevokeOutcomeRevoked`      | REVOKE_OUTCOME_REVOKED                    |
| `RevokeOutcomesRevokeOutcomeDenied`       | REVOKE_OUTCOME_DENIED                     |
| `RevokeOutcomesRevokeOutcomeError`        | REVOKE_OUTCOME_ERROR                      |
| `RevokeOutcomesRevokeOutcomeCancelled`    | REVOKE_OUTCOME_CANCELLED                  |
| `RevokeOutcomesRevokeOutcomeWaitTimedOut` | REVOKE_OUTCOME_WAIT_TIMED_OUT             |