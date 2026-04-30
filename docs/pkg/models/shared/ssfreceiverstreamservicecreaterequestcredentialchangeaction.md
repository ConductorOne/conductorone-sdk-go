# SSFReceiverStreamServiceCreateRequestCredentialChangeAction

Action to take when a credential-change event is received.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverStreamServiceCreateRequestCredentialChangeActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverStreamServiceCreateRequestCredentialChangeAction("custom_value")
```


## Values

| Name                                                                                        | Value                                                                                       |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `SSFReceiverStreamServiceCreateRequestCredentialChangeActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                                                           |
| `SSFReceiverStreamServiceCreateRequestCredentialChangeActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                                                            |
| `SSFReceiverStreamServiceCreateRequestCredentialChangeActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                                                              |