# SSFReceiverStreamServiceCreateRequestSessionRevokedAction

Per-event-type action configuration.
 Action to take when a session-revoked event is received.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverStreamServiceCreateRequestSessionRevokedActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverStreamServiceCreateRequestSessionRevokedAction("custom_value")
```


## Values

| Name                                                                                      | Value                                                                                     |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `SSFReceiverStreamServiceCreateRequestSessionRevokedActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                                                         |
| `SSFReceiverStreamServiceCreateRequestSessionRevokedActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                                                          |
| `SSFReceiverStreamServiceCreateRequestSessionRevokedActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                                                            |