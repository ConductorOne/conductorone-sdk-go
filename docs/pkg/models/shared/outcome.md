# Outcome

The action ConductorOne took in response to this event.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OutcomeSsfEventOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Outcome("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `OutcomeSsfEventOutcomeUnspecified`         | SSF_EVENT_OUTCOME_UNSPECIFIED               |
| `OutcomeSsfEventOutcomeSessionsRevoked`     | SSF_EVENT_OUTCOME_SESSIONS_REVOKED          |
| `OutcomeSsfEventOutcomeLogged`              | SSF_EVENT_OUTCOME_LOGGED                    |
| `OutcomeSsfEventOutcomePrincipalNotFound`   | SSF_EVENT_OUTCOME_PRINCIPAL_NOT_FOUND       |
| `OutcomeSsfEventOutcomeVerified`            | SSF_EVENT_OUTCOME_VERIFIED                  |
| `OutcomeSsfEventOutcomeStreamStatusUpdated` | SSF_EVENT_OUTCOME_STREAM_STATUS_UPDATED     |
| `OutcomeSsfEventOutcomeUnrecognized`        | SSF_EVENT_OUTCOME_UNRECOGNIZED              |
| `OutcomeSsfEventOutcomeError`               | SSF_EVENT_OUTCOME_ERROR                     |