# SessionRevokedAction

Per-canonical-type action configuration.
 Event types without a config here default to LOG_ONLY.
 Action to take when a session-revoked event is received.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionRevokedActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionRevokedAction("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `SessionRevokedActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                    |
| `SessionRevokedActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                     |
| `SessionRevokedActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                       |