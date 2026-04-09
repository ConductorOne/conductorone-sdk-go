# SSFReceiverEventSearchServiceSearchRequestOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverEventSearchServiceSearchRequestOutcome("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeUnspecified`         | SSF_EVENT_OUTCOME_UNSPECIFIED                                                         |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeSessionsRevoked`     | SSF_EVENT_OUTCOME_SESSIONS_REVOKED                                                    |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeLogged`              | SSF_EVENT_OUTCOME_LOGGED                                                              |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomePrincipalNotFound`   | SSF_EVENT_OUTCOME_PRINCIPAL_NOT_FOUND                                                 |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeVerified`            | SSF_EVENT_OUTCOME_VERIFIED                                                            |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeStreamStatusUpdated` | SSF_EVENT_OUTCOME_STREAM_STATUS_UPDATED                                               |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeUnrecognized`        | SSF_EVENT_OUTCOME_UNRECOGNIZED                                                        |
| `SSFReceiverEventSearchServiceSearchRequestOutcomeSsfEventOutcomeError`               | SSF_EVENT_OUTCOME_ERROR                                                               |