# SSFReceiverStreamServiceCreateRequestAccountDisabledAction

The accountDisabledAction field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverStreamServiceCreateRequestAccountDisabledActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverStreamServiceCreateRequestAccountDisabledAction("custom_value")
```


## Values

| Name                                                                                       | Value                                                                                      |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `SSFReceiverStreamServiceCreateRequestAccountDisabledActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                                                          |
| `SSFReceiverStreamServiceCreateRequestAccountDisabledActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                                                           |
| `SSFReceiverStreamServiceCreateRequestAccountDisabledActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                                                             |