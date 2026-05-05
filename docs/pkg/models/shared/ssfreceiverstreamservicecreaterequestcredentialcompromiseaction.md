# SSFReceiverStreamServiceCreateRequestCredentialCompromiseAction

Action to take when a credential-compromise event is received.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverStreamServiceCreateRequestCredentialCompromiseActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverStreamServiceCreateRequestCredentialCompromiseAction("custom_value")
```


## Values

| Name                                                                                            | Value                                                                                           |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `SSFReceiverStreamServiceCreateRequestCredentialCompromiseActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                                                               |
| `SSFReceiverStreamServiceCreateRequestCredentialCompromiseActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                                                                |
| `SSFReceiverStreamServiceCreateRequestCredentialCompromiseActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                                                                  |