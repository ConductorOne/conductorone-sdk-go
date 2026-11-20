# SSFReceiverEventOutcome

The action ConductorOne took in response to this event.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverEventOutcomeSsfEventOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverEventOutcome("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `SSFReceiverEventOutcomeSsfEventOutcomeUnspecified`         | SSF_EVENT_OUTCOME_UNSPECIFIED                               |
| `SSFReceiverEventOutcomeSsfEventOutcomeSessionsRevoked`     | SSF_EVENT_OUTCOME_SESSIONS_REVOKED                          |
| `SSFReceiverEventOutcomeSsfEventOutcomeLogged`              | SSF_EVENT_OUTCOME_LOGGED                                    |
| `SSFReceiverEventOutcomeSsfEventOutcomePrincipalNotFound`   | SSF_EVENT_OUTCOME_PRINCIPAL_NOT_FOUND                       |
| `SSFReceiverEventOutcomeSsfEventOutcomeVerified`            | SSF_EVENT_OUTCOME_VERIFIED                                  |
| `SSFReceiverEventOutcomeSsfEventOutcomeStreamStatusUpdated` | SSF_EVENT_OUTCOME_STREAM_STATUS_UPDATED                     |
| `SSFReceiverEventOutcomeSsfEventOutcomeUnrecognized`        | SSF_EVENT_OUTCOME_UNRECOGNIZED                              |
| `SSFReceiverEventOutcomeSsfEventOutcomeError`               | SSF_EVENT_OUTCOME_ERROR                                     |